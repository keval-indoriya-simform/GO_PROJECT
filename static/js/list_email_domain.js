$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://192.168.49.2:31471/api/v1/email-domains?page=1`)
    datatable=$('#emailDomainTable').DataTable({
        data: retrieveModelData["data"],
        deferRender: true,
        scrollY:        "60vh",
        scrollX:        "80%",
        scrollCollapse: true,
        scroller:       true,
        responsive: true,
        paging: false,
        fixedColumns: {
            leftColumns: 2,
        },
        order: [2, "asc"],
        columns: [
            {
                data: null,
                width: "20px",
                className: "dt-center editor-select",
                defaultContent: '<input class="bulk-select" type="checkbox"/>',
                orderable: false
            },
            {
                data: null,
                width: "20px",
                className: "dt-center editor-edit",
                defaultContent: '<i class="fa fa-pencil"/>',
                orderable: false
            },
            { data: 'customer_location' },
            { data: 'domain_registrar' },
            { data: 'link_for_domain_admin' },
            { data: 'account_number' },
            { data: 'account_number_aliases' },
            { data: 'account_user_name' },
            { data: 'password' },
            { data: 'pin' },
            { data: 'domain' },
            { data: 'a_record1' },
            { data: 'a_record2' },
            { data: 'a_record3' },
            { data: 'a_record4' },
            { data: 'email_hosting' },
            { data: 'link_for_email_admin' },
            { data: 'email_account_number' },
            { data: 'email_account_type' },
            { data: 'email_user_name' },
            { data: 'email_account_password' },
            { data: 'email_pin' },
            { data: 'mx_record1' },
            { data: 'mx_record2' },
            { data: 'website_ip_or_alias' },
            { data: 'web_mail_pop_imap' },
            { data: 'web_mail_exchange' },
            { data: 'notes' }

        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://192.168.49.2:31471/api/v1/email-domains?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#emailDomainTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["email_domain_id"]) === false) {
                list.push(datatable.row(this).data()["email_domain_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["email_domain_id"])
        }
    });

    $('#emailDomainTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["email_domain_id"])
        const myModal = new bootstrap.Modal('#emailDomainModal', {})
        $("#emailDomainModal").removeAttr("data-id")
        $("#emailDomainModal").attr("data-id", datatable.row(this).data()["email_domain_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://192.168.49.2:31471/api/v1/email-domains?email_domain_id="+list.toString(),
            `http://192.168.49.2:31471/network-management-solutions/list/email-domains`)
    })

    $('#emailDomainModal').on('show.bs.modal', async function() {
        let dataId = $("#emailDomainModal").attr("data-id");
        console.log(dataId)
        let typesResponse=await fetch('http://192.168.49.2:31471/api/v1/email-domains/email-account-types?set_limit=false')
        let typesData=await typesResponse.json()
        let emailResponse = await fetch('http://192.168.49.2:31471/api/v1/email-domains?email_domain_id=' + dataId);
        let emailData = await emailResponse.json()
        console.log(emailData)
        let selectedValue = emailData["data"][0]["email_account_type_id"]
        options =""
        for (index = 0; index < typesData["data"].length; index++) {
            if (selectedValue !== emailData["data"][0]["email_account_type_id"]){
                options += `<option data-tokens="`+typesData["data"][index]["email_account_type"]+`" value="`+ typesData["data"][index]["email_account_type_id"] + `">` + typesData["data"][index]["email_account_type"] + `</option>`
            } else {
                options += `<option selected data-tokens="`+typesData["data"][index]["email_account_type"]+`" value="`+ typesData["data"][index]["email_account_type_id"] + `">` + typesData["data"][index]["email_account_type"] + `</option>`
            }
        }
        $("#customer_location").val(emailData["data"][0]["customer_location"])
        $('.selectpicker1').append(options).selectpicker('refresh');
        $("#domain_registrar").val(emailData["data"][0]["domain_registrar"])
        $("#account_number").val(emailData["data"][0]["account_number"])
        $("#link_for_domain_admin").val(emailData["data"][0]["link_for_domain_admin"])
        $("#account_name").val(emailData["data"][0]["account_name"])
        $("#account_number_aliases").val(emailData["data"][0]["account_number_aliases"])
        $("#account_user_name").val(emailData["data"][0]["account_user_name"])
        $("#password").val(emailData["data"][0]["password"])
        $("#pin").val(emailData["data"][0]["pin"])
        $("#domain").val(emailData["data"][0]["domain"])
        $("#a_record_1").val(emailData["data"][0]["a_record1"])
        $("#a_record_2").val(emailData["data"][0]["a_record2"])
        $("#a_record_3").val(emailData["data"][0]["a_record3"])
        $("#a_record_4").val(emailData["data"][0]["a_record4"])
        $("#email_hosting").val(emailData["data"][0]["email_hosting"])
        $("#link_for_email_admin").val(emailData["data"][0]["link_for_email_admin"])
        $("#email_account_number").val(emailData["data"][0]["email_account_number"])
        $("#email_user_name").val(emailData["data"][0]["email_user_name"])
        $("#email_password").val(emailData["data"][0]["email_account_password"])
        $("#email_pin").val(emailData["data"][0]["email_pin"])
        $("#mx_record_1").val(emailData["data"][0]["mx_record1"])
        $("#mx_record_2").val(emailData["data"][0]["mx_record2"])
        $("#website_ip_or_alias").val(emailData["data"][0]["website_ip_or_alias"])
        $("#web_mail_pop_imap").val(emailData["data"][0]["web_mail_pop_imap"])
        $("#web_mail_exchange").val(emailData["data"][0]["web_mail_exchange"])
        $("#notes").val(emailData["data"][0]["notes"])
        $("#phone_setting_note").val(emailData["data"][0]["phone_setting_note"])


        let update_button = $("#update-btn")
        update_button.click(async function(){
            await patchModel(
                JSON.stringify({
                    domain_registrar: $("#domain_registrar").val(),
                    account : {
                        account_number: $("#account_number").val(),
                        link_for_domain_admin: $("#link_for_domain_admin").val(),
                        name:$("#account_name").val(),
                        account_number_aliases: $("#account_number_aliases").val(),
                        user_name:$("#account_user_name").val()
                    },
                    password:$("#password").val(),
                    pin:$("#pin").val(),
                    domain: $("#domain").val(),
                    a_record_1:$("#a_record_1").val(),
                    a_record_2:$("#a_record_2").val(),
                    a_record_3:$("#a_record_3").val(),
                    a_record_4:$("#a_record_4").val(),
                    email_account:{
                        email_account_type_id:  parseInt($('.selectpicker1').val()),
                        email_hosting:$("#email_hosting").val(),
                        link_for_email_admin:$("#link_for_email_admin").val(),
                        account_number:$("#email_account_number").val(),
                        user_name:$("#email_user_name").val(),
                        password:$("#email_password").val(),
                        pin: $("#email_pin").val(),
            },
                    mx_record_1:$("#mx_record_1").val(),
                    mx_record_2:$("#mx_record_2").val(),
                    website_ip_or_alias:$("#website_ip_or_alias").val(),
                    web_mail_pop_imap:$("#web_mail_pop_imap").val(),
                    web_mail_exchange:$("#web_mail_exchange").val(),
                    notes:$("#notes").val(),
                    phone_setting_note:$("#phone_setting_note").val(),
                    updated_by_user_id: parseInt($("#user_id").val()),
                }),
                "http://192.168.49.2:31471/api/v1/email-domains?email_domain_id="+dataId,
                `http://192.168.49.2:31471/network-management-solutions/list/email-domains`
            )
        });
    });

    $('#emailDomainModal').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});