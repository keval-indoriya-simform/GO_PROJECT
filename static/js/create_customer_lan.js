$(document).ready(async function() {
    let submit_button = $("#submit-btn"),
    select_customer_location = $("#customer_location_name")
    const customerLocationResponse = await fetch('http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name&set_limit=false');
    const customerLocationData = await customerLocationResponse.json();
    options =""
    for (let index = 0; index < customerLocationData["data"].length; index++) {
        options += `<option data-tokens="`+customerLocationData["data"][index]["name"]+`" value="`+ customerLocationData["data"][index]["customer_location_id"] + `">` + customerLocationData["data"][index]["name"] + `</option>`;
    }
    select_customer_location.append(options).selectpicker('refresh');

    const cloudPrivateIpResponse = await fetch('http://localhost:8080/api/v1/cloud-private-ips?select_column=cloud_private_ips.cloud_private_ip_id,cloud_private_ips.ipv4_assignment&set_limit=false');
    const cloudPrivateIpData = await cloudPrivateIpResponse.json();
    let options_private_ip =""
    for (let index = 0; index < cloudPrivateIpData["data"].length; index++) {
        options_private_ip += `<option data-tokens="`+cloudPrivateIpData["data"][index]["ipv4_assignment"]+`" value="`+ cloudPrivateIpData["data"][index]["cloud_private_ip_id"] + `">` + cloudPrivateIpData["data"][index]["ipv4_assignment"] + `</option>`;
    }
    let select_cloud_private_ip = $("#cloudPrivateIp")
    select_cloud_private_ip.append(options_private_ip).selectpicker('refresh');

    submit_button.click(async function(){
        let customer_lan = jsonParsing()
        customer_lan["updated_by_user_id"] = parseInt($("#user_id").val())
        await postModel(
            JSON.stringify(customer_lan),
            "http://localhost:8080/api/v1/customer-lans"
        )
    });


});

function checkWanConfigEmpty(object){
    return (object["firewall"] !== "" || object["sub_net_mask"] !== "" || object["gateway"] !== "" || object["ip_range"] !== "" || object["comment"] !== "");
}

function checkInternetProviderEmpty(object){
    return (object["name"] === "" && object["account_or_pin"] === "" && object["other"] === "" && object["speed"] === "" && object["primary_dns"] === "" && object["secondary_dns"] === "");
}

const jsonParsing = () => {
    'use strict'
    let customer_lan_model = {}
    let internet_provider1_model = {}
    let internet_provider2_model = {}
    let ip1_wan_ipv4 = {}
    let ip1_wan_ipv6 = {}
    let ip2_wan_ipv4 = {}
    let ip2_wan_ipv6 = {}

    ip1_wan_ipv4["firewall"] = $("#firewall_wan1_ipv4").val()
    ip1_wan_ipv4["sub_net_mask"] = $("#subnet_mask_wan1_ipv4").val()
    ip1_wan_ipv4["gateway"] = $("#gateway_wan1_ipv4").val()
    ip1_wan_ipv4["ip_version"] = "4"
    ip1_wan_ipv4["ip_range"] = $("#ip_range_wan1_ipv4").val()
    ip1_wan_ipv4["comment"] = $("#comment_wan1_ipv4").val()

    ip1_wan_ipv6["firewall"] = $("#firewall_wan1_ipv6").val()
    ip1_wan_ipv6["sub_net_mask"] = $("#subnet_mask_wan1_ipv6").val()
    ip1_wan_ipv6["gateway"] = $("#gateway_wan1_ipv6").val()
    ip1_wan_ipv6["ip_version"] = "6"
    ip1_wan_ipv6["ip_range"] = $("#ip_range_wan1_ipv6").val()
    ip1_wan_ipv6["comment"] = $("#comment_wan1_ipv6").val()

    ip2_wan_ipv4["firewall"] = $("#firewall_wan2_ipv4").val()
    ip2_wan_ipv4["sub_net_mask"] = $("#subnet_mask_wan2_ipv4").val()
    ip2_wan_ipv4["gateway"] = $("#gateway_wan2_ipv4").val()
    ip2_wan_ipv4["ip_version"] = "4"
    ip2_wan_ipv4["ip_range"] = $("#ip_range_wan2_ipv4").val()
    ip2_wan_ipv4["comment"] = $("#comment_wan2_ipv4").val()

    ip2_wan_ipv6["firewall"] = $("#firewall_wan2_ipv6").val()
    ip2_wan_ipv6["sub_net_mask"] = $("#subnet_mask_wan2_ipv6").val()
    ip2_wan_ipv6["gateway"] = $("#gateway_wan2_ipv6").val()
    ip2_wan_ipv6["ip_version"] = "6"
    ip2_wan_ipv6["ip_range"] = $("#ip_range_wan2_ipv6").val()
    ip2_wan_ipv6["comment"] = $("#comment_wan2_ipv6").val()

    internet_provider1_model["name"] = $("#internet_provider1_name").val()
    internet_provider1_model["account_or_pin"] = $("#internet_provider1_account_pin").val()
    internet_provider1_model["other"] = $("#internet_provider1_other").val()
    internet_provider1_model["speed"] = $("#internet_provider1_speed").val()
    internet_provider1_model["primary_dns"] = $("#internet_provider1_primary_dns").val()
    internet_provider1_model["secondary_dns"] = $("#internet_provider1_secondary_dns").val()
    internet_provider1_model["wan_config_ipv4"] = ip1_wan_ipv4

    internet_provider2_model["name"] = $("#internet_provider2_name").val()
    internet_provider2_model["account_or_pin"] = $("#internet_provider2_account_pin").val()
    internet_provider2_model["other"] = $("#internet_provider2_other").val()
    internet_provider2_model["speed"] = $("#internet_provider2_speed").val()
    internet_provider2_model["primary_dns"] = $("#internet_provider2_primary_dns").val()
    internet_provider2_model["secondary_dns"] = $("#internet_provider2_secondary_dns").val()

    customer_lan_model["created_by_user_id"] = parseInt($("#user_id").val())
    customer_lan_model["customer_location_id"] = parseInt($("#customer_location_name").val())
    customer_lan_model["network_on_site"] = $("#customer_lan_network_onsite").val()
    customer_lan_model["number_of_internet_connection"] = parseInt($("#number_of_internet_connection").val())
    customer_lan_model["private_ip_assignment_id"] = parseInt($("#cloudPrivateIp").val())
    customer_lan_model["firewall_ipv4_lan_address"] = $("#firewall_ipv4_lan_address").val()
    customer_lan_model["firewall_ipv6_lan_address"] = $("#firewall_ipv6_lan_address").val()
    customer_lan_model["gateway_type"] = $("#gateway_type").val()
    customer_lan_model["equipment_installed"] = $("#equipment_installed").val()
    customer_lan_model["backup_date"] = new Date($("#backup_date").val())
    customer_lan_model["version_of_last_backup"] = $("#version_of_last_backup").val()
    customer_lan_model["extra"] = $("#extra").val()
    customer_lan_model["lan_notes"] = $("#lan_notes").val()
    customer_lan_model["scan_to_folder_location"] = $("#scan_to_folder_location").val()
    customer_lan_model["scan_to_folder_username_or_password"] = $("#scan_to_folder_username_or_password").val()
    customer_lan_model["scan_to_email_smtp_server_port"] = $("#scan_to_email_smtp_server_port").val()
    customer_lan_model["scan_to_email_email_or_password"] = $("#scan_to_email_email_or_password").val()
    customer_lan_model["on_site_backup_server_or_nas_type"] = $("#onsite_backup_server_or_nas_type").val()
    customer_lan_model["on_site_backup_server_or_nas_ip_address"] = $("#onsite_backup_server_or_nas_ip_address").val()
    customer_lan_model["management_server"] = $("#management_server").val()
    customer_lan_model["management_notes"] = $("#management_notes").val()
    customer_lan_model["management_ip_address"] = $("#management_ip_address").val()
    customer_lan_model["wireless_unit"] = $("#wireless_unit").val()
    customer_lan_model["wireless_ip_address"] = $("#wireless_ip_address").val()
    customer_lan_model["wireless_admin_username"] = $("#wireless_admin_username").val()
    customer_lan_model["wireless_admin_password"] = $("#wireless_admin_password").val()
    customer_lan_model["wireless_ssid"] = $("#wireless_ssid").val()
    customer_lan_model["wireless_password"] = $("#wireless_password").val()
    customer_lan_model["wireless_connection_type"] = $("#wireless_connection_type").val()
    customer_lan_model["wireless_notes"] = $("#wireless_notes").val()
    customer_lan_model["switch_brand_or_model"] = $("#switch_brand_or_model").val()
    customer_lan_model["switch_credentials"] = $("#switch_credentials").val()
    customer_lan_model["switch_manage"] = $("#switch_manage").val()
    customer_lan_model["switch_ip_address"] = $("#switch_ip_address").val()
    customer_lan_model["switch_install_date"] = new Date($("#switch_install_date").val())
    let list_of_links = []
    if ($("#image_file1").val().replace(/C:\\fakepath\\/i, '') !== "") {
        list_of_links.push($("#image_file1").val().replace(/C:\\fakepath\\/i, ''))
    }
    if ($("#image_file2").val().replace(/C:\\fakepath\\/i, '') !== "") {
        list_of_links.push($("#image_file2").val().replace(/C:\\fakepath\\/i, ''))
    }
    if ($("#image_file2").val().replace(/C:\\fakepath\\/i, '') !== "") {
        list_of_links.push($("#image_file3").val().replace(/C:\\fakepath\\/i, ''))
    }

    customer_lan_model["switch_image_links"] = list_of_links.toString()

    customer_lan_model["switch_notes"] = $("#switch_notes").val()
    customer_lan_model["print_or_scanner_type"] = $("#print_or_scanner_type").val()
    customer_lan_model["print_or_scanner_user_name"] = $("#print_or_scanner_username").val()
    customer_lan_model["print_or_scanner_password"] = $("#print_or_scanner_password").val()
    customer_lan_model["print_or_scanner_ip_address"] = $("#print_or_scanner_ip_address").val()
    customer_lan_model["internet_provider1"] = internet_provider1_model
    customer_lan_model["internet_provider2"] = internet_provider2_model

    if (checkWanConfigEmpty(ip1_wan_ipv6) === true){
        customer_lan_model["internet_provider1"]["wan_config_ipv6"] = ip1_wan_ipv6
    }

    if (checkWanConfigEmpty(ip2_wan_ipv4) === true ){
        customer_lan_model["internet_provider2"]["wan_config_ipv4"] = ip2_wan_ipv4
    }

    if (checkWanConfigEmpty(ip2_wan_ipv6) === true){
        customer_lan_model["internet_provider2"]["wan_config_ipv6"] = ip2_wan_ipv6
    }

    if (checkInternetProviderEmpty(customer_lan_model["internet_provider2"]) === true && checkWanConfigEmpty(ip2_wan_ipv4) === false && checkWanConfigEmpty(ip2_wan_ipv6) === false) {
        delete customer_lan_model["internet_provider2"]
    }

    return customer_lan_model
}


