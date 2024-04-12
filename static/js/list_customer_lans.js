$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-lans?page=1`)
    datatable = $('#customerLanTable').DataTable({
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
            { data: 'network_on_site', title: "Network On Site" },
            {
                data: "customer_location",
                title: "Related To",
                className: "popup",
                render: function (data) {
                    return '<span class="popup-element">' + data + '</span>';
                }
            },
            { data: 'backup_date', title:"Backup Date" },
            { data: 'equipment_installed', title: "Equipment Installed" },
            { data: 'gateway_type', title: "Gateway Type" },
            { data: 'wireless_unit', title: "Wireless Unit" },
            { data: 'extra', title: "Extra" },
            { data: 'firewall_ipv4_lan_address', title: "Firewall LAN Address (IPV4)" },
            { data: 'firewall_ipv6_lan_address', title: "Firewall LAN Address (IPV6)" },
            { data: 'internet_providers1_name', title: "Internet Provider 1 Name" },
            { data: 'internet_providers1_account_or_pin', title: "Internet Provider 1 Account/pin" },
            { data: 'internet_providers1_primary_dns', title: "Internet Provider 1 Primary DNS" },
            { data: 'internet_providers1_secondary_dns', title: "Internet Provider 1 Secondary DNS" },
            { data: 'internet_providers1_speed', title: "Internet Provider 1 Speed" },
            { data: 'firewall_wan1_ipv4', title: "Firewall WAN-1 (IPV4)" },
            { data: 'gateway_wan1_ipv4', title: "Gateway WAN-1 (IPV4)" },
            { data: 'sub_net_mask_wan1_ipv4', title: "Sub Net Mask WAN-1 (IPV4)" },
            { data: 'firewall_wan1_ipv6', title: "Firewall WAN-1 (IPV6)" },
            { data: 'gateway_wan1_ipv6', title: "Gateway WAN-1 (IPV6)" },
            { data: 'sub_net_mask_wan1_ipv6', title: "Sub Net Mask WAN-1 (IPV6)" },
            { data: 'internet_providers2_account_or_pin', title: "Internet Provider 2 Account/pin" },
            { data: 'internet_providers2_name', title: "Internet Provider 2 Name" },
            { data: 'internet_providers2_primary_dns', title: "Internet Provider 2 Primary DNS" },
            { data: 'internet_providers2_secondary_dns', title: "Internet Provider 2 Secondary DNS" },
            { data: 'internet_providers2_speed', title: "Internet Provider 2 Speed" },
            { data: 'firewall_wan2_ipv4', title: "Firewall WAN-2 (IPV4)" },
            { data: 'gateway_wan2_ipv4', title: "Gateway WAN-2 (IPV4)" },
            { data: 'sub_net_mask_wan2_ipv4', title: "Sub Net Mask WAN-2 (IPV4)" },
            { data: 'firewall_wan2_ipv6', title: "Firewall WAN-2 (IPV6)" },
            { data: 'gateway_wan2_ipv6', title: "Gateway WAN-2 (IPV6)" },
            { data: 'sub_net_mask_wan2_ipv6', title: "Sub Net Mask WAN-2 (IPV6)" },
            { data: 'ipv4_assignment', title: "Private IP Assignment" },
            { data: 'management_server', title: "Management Server" },
            { data: 'number_of_internet_connection', title: "Number Of Internet Connection" },
            { data: 'print_or_scanner_ip_address', title: "Print Or Scanner IP Address" },
            { data: 'print_or_scanner_password', title: "Print Or Scanner Password" },
            { data: 'print_or_scanner_user_name', title: "Print Or Scanner Username" },
            { data: 'scan_to_email_email_or_password', title: "Scan To Email: Email Or Password" },
            { data: 'scan_to_email_smtp_server_port', title: "Scan To Email: SMTP Server Port" },
            { data: 'scan_to_folder_location', title: "Scan To Folder: Location" },
            { data: 'scan_to_folder_username_or_password', title: "Scan To Folder: Username Or Password" },
            { data: 'version_of_last_backup', title: "Version Of Last Backup" },
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-lans?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }

    list = []
    $('#customerLanTable').on('click','.editor-select', async function(){
        let check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["customer_lan_id"]) === false) {
                list.push(datatable.row(this).data()["customer_lan_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["customer_lan_id"])
        }
    });

    $('#customerLanTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["customer_lan_id"])
        const myModal = new bootstrap.Modal('#modalCustomerLan', {})
        $("#modalCustomerLan").removeAttr("data-id")
        $("#modalCustomerLan").attr("data-id", datatable.row(this).data()["customer_lan_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
            delete_data("http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-lans?customer_lan_id="+list.toString(),
                `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/network-management-solutions/list/customer-lans`)
    })

    $('#modalCustomerLan').on('show.bs.modal', async function() {
        var dataId = $(this).attr("data-id");
        let customerLanResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-lans?' +
            'select_column=customer_lans.*,internet_providers1.other as internet_providers1_other,' +
            'wan_configs1_ipv4.ip_range as ip_range_wan1_ipv4, wan_configs1_ipv4.comment as comment_wan1_ipv4,' +
            'wan_configs1_ipv6.ip_range as ip_range_wan1_ipv6, wan_configs1_ipv6.comment as comment_wan1_ipv6,' +
            'wan_configs2_ipv4.ip_range as ip_range_wan2_ipv4, wan_configs2_ipv4.comment as comment_wan2_ipv4,' +
            'wan_configs2_ipv6.ip_range as ip_range_wan2_ipv6, wan_configs2_ipv6.comment as comment_wan2_ipv6,' +
            'internet_providers2.other as internet_providers2_other' +
            '&append_select=true&customer_lan_id=' + dataId);
        let customerLanData = await customerLanResponse.json()
        console.log("hii")
        console.log(customerLanData)
        $("#customer_location_name").val(customerLanData["data"][0]["customer_location"])
        $("#customer_lan_network_onsite").val(customerLanData["data"][0]["network_on_site"])
        $("#number_of_internet_connection").val(customerLanData["data"][0]["number_of_internet_connection"])
        $("#internet_provider1_name").val(customerLanData["data"][0]["internet_providers1_name"])
        $("#internet_provider1_account_pin").val(customerLanData["data"][0]["internet_providers1_account_or_pin"])
        $("#internet_provider1_other").val(customerLanData["data"][0]["internet_providers1_other"])
        $("#internet_provider1_speed").val(customerLanData["data"][0]["internet_providers1_speed"])
        $("#firewall_wan1_ipv4").val(customerLanData["data"][0]["firewall_wan1_ipv4"])
        $("#subnet_mask_wan1_ipv4").val(customerLanData["data"][0]["sub_net_mask_wan1_ipv4"])
        $("#gateway_wan1_ipv4").val(customerLanData["data"][0]["gateway_wan1_ipv4"])
        $("#ip_range_wan1_ipv4").val(customerLanData["data"][0]["ip_range_wan1_ipv4"])
        $("#comment_wan1_ipv4").val(customerLanData["data"][0]["comment_wan1_ipv4"])
        $("#firewall_wan1_ipv6").val(customerLanData["data"][0]["firewall_wan1_ipv6"])
        $("#subnet_mask_wan1_ipv6").val(customerLanData["data"][0]["sub_net_mask_wan1_ipv6"])
        $("#gateway_wan1_ipv6").val(customerLanData["data"][0]["gateway_wan1_ipv6"])
        $("#ip_range_wan1_ipv6").val(customerLanData["data"][0]["ip_range_wan1_ipv6"])
        $("#comment_wan1_ipv6").val(customerLanData["data"][0]["comment_wan1_ipv6"])
        $("#internet_provider1_primary_dns").val(customerLanData["data"][0]["internet_providers1_primary_dns"])
        $("#internet_provider1_secondary_dns").val(customerLanData["data"][0]["internet_providers1_secondary_dns"])
        $("#internet_provider2_name").val(customerLanData["data"][0]["internet_providers2_name"])
        $("#internet_provider2_account_pin").val(customerLanData["data"][0]["internet_providers2_account_or_pin"])
        $("#internet_provider2_other").val(customerLanData["data"][0]["internet_providers2_other"])
        $("#internet_provider2_speed").val(customerLanData["data"][0]["internet_providers2_speed"])
        $("#firewall_wan2_ipv4").val(customerLanData["data"][0]["firewall_wan2_ipv4"])
        $("#subnet_mask_wan2_ipv4").val(customerLanData["data"][0]["sub_net_mask_wan2_ipv4"])
        $("#gateway_wan2_ipv4").val(customerLanData["data"][0]["gateway_wan2_ipv4"])
        $("#ip_range_wan2_ipv4").val(customerLanData["data"][0]["ip_range_wan2_ipv4"])
        $("#comment_wan2_ipv4").val(customerLanData["data"][0]["comment_wan2_ipv4"])
        $("#firewall_wan2_ipv6").val(customerLanData["data"][0]["firewall_wan2_ipv6"])
        $("#subnet_mask_wan2_ipv6").val(customerLanData["data"][0]["sub_net_mask_wan2_ipv6"])
        $("#gateway_wan2_ipv6").val(customerLanData["data"][0]["gateway_wan2_ipv6"])
        $("#ip_range_wan2_ipv6").val(customerLanData["data"][0]["ip_range_wan2_ipv6"])
        $("#comment_wan2_ipv6").val(customerLanData["data"][0]["comment_wan2_ipv6"])
        $("#internet_provider2_primary_dns").val(customerLanData["data"][0]["internet_providers2_primary_dns"])
        $("#internet_provider2_secondary_dns").val(customerLanData["data"][0]["internet_providers2_secondary_dns"])
        $("#cloudPrivateIp").val(customerLanData["data"][0]["ipv4_assignment"])
        $("#firewall_ipv4_lan_address").val(customerLanData["data"][0]["firewall_ipv4_lan_address"])
        $("#firewall_ipv6_lan_address").val(customerLanData["data"][0]["firewall_ipv6_lan_address"])
        $("#gateway_type").val(customerLanData["data"][0]["gateway_type"])
        $("#equipment_installed").val(customerLanData["data"][0]["equipment_installed"])
        $("#backup_date").val(new Date(customerLanData["data"][0]["backup_date"]).toISOString().split('T')[0])
        $("#version_of_last_backup").val(customerLanData["data"][0]["version_of_last_backup"])
        $("#extra").val(customerLanData["data"][0]["extra"])
        $("#print_or_scanner_type").val(customerLanData["data"][0]["print_or_scanner_type"])
        $("#print_or_scanner_ip_address").val(customerLanData["data"][0]["print_or_scanner_ip_address"])
        $("#print_or_scanner_username").val(customerLanData["data"][0]["print_or_scanner_user_name"])
        $("#print_or_scanner_password").val(customerLanData["data"][0]["print_or_scanner_password"])
        $("#scan_to_folder_location").val(customerLanData["data"][0]["scan_to_folder_location"])
        $("#scan_to_folder_username_or_password").val(customerLanData["data"][0]["scan_to_folder_username_or_password"])
        $("#scan_to_email_smtp_server_port").val(customerLanData["data"][0]["scan_to_email_smtp_server_port"])
        $("#scan_to_email_email_or_password").val(customerLanData["data"][0]["scan_to_email_email_or_password"])
        $("#management_server").val(customerLanData["data"][0]["management_server"])
        $("#management_ip_address").val(customerLanData["data"][0]["management_ip_address"])
        $("#management_notes").val(customerLanData["data"][0]["management_notes"])
        $("#onsite_backup_server_or_nas_type").val(customerLanData["data"][0]["on_site_backup_server_or_nas_type"])
        $("#onsite_backup_server_or_nas_ip_address").val(customerLanData["data"][0]["on_site_backup_server_or_nas_ip_address"])
        $("#wireless_unit").val(customerLanData["data"][0]["wireless_unit"])
        $("#wireless_ip_address").val(customerLanData["data"][0]["wireless_ip_address"])
        $("#wireless_admin_username").val(customerLanData["data"][0]["wireless_admin_username"])
        $("#wireless_admin_password").val(customerLanData["data"][0]["wireless_admin_password"])
        $("#wireless_ssid").val(customerLanData["data"][0]["wireless_ssid"])
        $("#wireless_password").val(customerLanData["data"][0]["wireless_password"])
        $("#wireless_connection_type").val(customerLanData["data"][0]["wireless_connection_type"])
        $("#wireless_notes").val(customerLanData["data"][0]["wireless_notes"])
        $("#switch_brand_or_model").val(customerLanData["data"][0]["switch_brand_or_model"])
        $("#switch_credentials").val(customerLanData["data"][0]["switch_credentials"])
        $("#switch_manage").val(customerLanData["data"][0]["switch_manage"])
        $("#switch_ip_address").val(customerLanData["data"][0]["switch_ip_address"])
        $("#switch_install_date").val(new Date(customerLanData["data"][0]["switch_install_date"]).toISOString().split('T')[0])
        $("#switch_notes").val(customerLanData["data"][0]["switch_notes"])
        $("#lan_notes").val(customerLanData["data"][0]["lan_notes"])
            // $("#image_file1").val(customerLanData["data"][0]["customer_location"])
            // $("#image_file2").val(customerLanData["data"][0]["customer_location"])
            // $("#image_file3").val(customerLanData["data"][0]["customer_location"])

        const cloudPrivateIpResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?' +
            'select_column=cloud_private_ips.cloud_private_ip_id,cloud_private_ips.ipv4_assignment&set_limit=false');
        const cloudPrivateIpData = await cloudPrivateIpResponse.json();
        let options =""
        let selectedValue = customerLanData["data"][0]["private_ip_assignment_id"]
        for (let index = 0; index < cloudPrivateIpData["data"].length; index++) {
            if (selectedValue !== cloudPrivateIpData["data"][index]["cloud_private_ip_id"]){
                options += `<option data-tokens="` + cloudPrivateIpData["data"][index]["ipv4_assignment"] + `" value="` + cloudPrivateIpData["data"][index]["cloud_private_ip_id"] + `">` + cloudPrivateIpData["data"][index]["ipv4_assignment"] + `</option>`;
            } else {
                options += `<option selected data-tokens="` + cloudPrivateIpData["data"][index]["ipv4_assignment"] + `" value="` + cloudPrivateIpData["data"][index]["cloud_private_ip_id"] + `">` + cloudPrivateIpData["data"][index]["ipv4_assignment"] + `</option>`;
            }
        }
        let select_cloud_private_ip = $("#cloudPrivateIp")
        select_cloud_private_ip.append(options).selectpicker('refresh');
        let update_button = $("#update-btn")
        update_button.click(async function(){
            let customer_lan = jsonParsing()
            customer_lan["updated_by_user_id"] = parseInt($("#user_id").val())
            console.log(customer_lan)
            await patchModel(
                JSON.stringify(customer_lan),
                "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-lans?customer_lan_id=" + dataId,
                `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-lans?page=`
            )
        });
    });

    $('#modalCustomerLan').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
    $('#customerLanTable').on('click','.popup', function(event){
        if (!$(event.target).hasClass("popup-element")) {
            $(".popup-content").css("display", "flex");
        }
        var clickedData = $(this).text();
        console.log(clickedData)
        fetch('/api/v1/customer-locations?set_limit=false')
            .then(response => response.json())
            .then(response => {
                [response].forEach(item => {
                    console.log(item)
                    for(var i= 0; i< item.data.length; i++){
                        if(clickedData===item.data[i].name){
                            $("#customer_location").text(item.data[i].name)
                            $("#customer_name").text(item.data[i].customers)
                        }
                    }
                });
            });
    });

    $(".close").click(function (){
        $(".popup-content").css("display", "none");
    });

});