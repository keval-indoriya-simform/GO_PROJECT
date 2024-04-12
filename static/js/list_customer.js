$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers?page=1`)
    datatable=$('#customerTable').DataTable({
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
            { data: 'name' },
            { data: 'voip' },
            { data: 'internet' },
            { data: 'firewall' },
            { data: 'backup_software' },
            { data: 'is_active' },
            { data: 'cloud_or_onsite' },
            { data: 'hardware_as_a_service' }
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#customerTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["customer_id"]) === false) {
                list.push(datatable.row(this).data()["customer_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["customer_id"])
        }
    });

    $('#customerTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["customer_id"])
        const myModal = new bootstrap.Modal('#customerModal', {})
        $("#customerModal").removeAttr("data-id")
        $("#customerModal").attr("data-id", datatable.row(this).data()["customer_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers?customer_id="+list.toString(),
            `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/network-management-solutions/list/customers`)
    })

    $('#customerModal').on('show.bs.modal', async function() {
        let dataId = $("#customerModal").attr("data-id");
        console.log(dataId)
        let cloudResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers/cloud-or-onsites?set_limit=false');
        let cloudData = await cloudResponse.json();
        options =""
        let customerResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers?customer_id=' + dataId);
        let customerData = await customerResponse.json()
        console.log(customerData)
        $("#name").val(customerData["data"][0]["name"])
        let selectedValue = customerData["data"][0]["cloud_or_onsite_id"]
        for (index = 0; index < cloudData["data"].length; index++) {
            if (selectedValue !== cloudData["data"][index]["cloud_or_onsite_id"]){
                options += `<option data-tokens="` + cloudData["data"][index]["name"] + `" value="` + cloudData["data"][index]["cloud_or_onsite_id"] + `">` + cloudData    ["data"][index]["name"] + `</option>`;
            } else {
                options += `<option selected data-tokens="` + cloudData["data"][index]["name"] + `" value="` + cloudData["data"][index]["cloud_or_onsite_id"] + `">` + cloudData    ["data"][index]["name"] + `</option>`;
            }
        }
        $('.selectpicker').append(options).selectpicker('refresh');
        $("#backup_software").val(customerData["data"][0]["backup_software"])
        $("#active_check").prop('checked', customerData["data"][0]["is_active"])
        $("#voip_check").prop('checked',customerData["data"][0]["voip"])
        $("#internet_check").prop('checked',customerData["data"][0]["internet"])
        $("#firewall_check").prop('checked',customerData["data"][0]["firewall"])
        $("#hardware_as_a_service_check").prop('checked',customerData["data"][0]["hardware_as_a_service"])
        $("#description").val(customerData["data"][0]["description"])

        let update_button = $("#update-btn")
        update_button.click(async function(){
            console.log($("#active_check").is(':checked'))
            await patchModel(
                JSON.stringify({
                    name:  $("#name").val(),
                    cloud_or_onsite_id: parseInt($("#cloud_or_onsite").val()),
                    backup_software : $("#backup_software").val(),
                    is_active : $("#active_check").is(':checked'),
                    voip :  $("#voip_check").is(':checked'),
                    internet : $("#internet_check").is(':checked'),
                    firewall : $("#firewall_check").is(':checked'),
                    hardware_as_a_service : $("#hardware_as_a_service_check").is(':checked'),
                    description : $("#description").val(),
                    updated_by_user_id: parseInt($("#user_id").val())
                }),
                "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers?customer_id="+dataId,
                `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customers?page=`
            )
        });
    });

    $('#customerModal').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});0