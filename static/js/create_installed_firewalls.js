let create_button = $("#submit-btn"),
    select_customer_location = $("#validationCustom05"),
    customer_lan_on_site = $("#validationCustom02"),
    brand = $("#validationCustom03"),
    equipment = $("#validationCustom04"),
    installed_date = $("#datepicker-1"),
    firewall_wan1_ipv4 = $("#validationCustom06"),
    ip_range = $("#validationCustom07"),
    subnet_mask_wan1_ipv4 = $("#validationCustom08"),
    gateway_wan1_ipv4 = $("#validationCustom09"),
    firewall_wan1_ipv6 = $("#validationCustom10"),
    subnet_mask_wan1_ipv6 = $("#validationCustom11"),
    gateway_wan1_ipv6 = $("#validationCustom12"),
    firewall_ipv4_lan_address = $("#validationCustom13"),
    firewall_ipv6_lan_address = $("#validationCustom14"),
    current_version = $("#validationCustom15"),
    version_backup = $("#validationCustom16"),
    backup_date = $("#datepicker-2"),
    extra = $("#validationCustom17"),
    firewall_wan2_ipv4 = $("#validationCustom18"),
    subnet_mask_wan2_ipv4 = $("#validationCustom19"),
    gateway_wan2_ipv4 = $("#validationCustom20"),
    firewall_wan2_ipv6 = $("#validationCustom21"),
    subnet_mask_wan2_ipv6 = $("#validationCustom22"),
    gateway_wan2_ipv6 = $("#validationCustom23")

$(document).ready(async function() {
    const customerResponse = await fetch('http://192.168.49.2:31471/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name&set_limit=false');
    const customerData = await customerResponse.json();
    options =""
    for (index = 0; index < customerData["data"].length; index++) {
        options += `<option data-tokens="`+customerData["data"][index]["name"]+`" value="`+ customerData["data"][index]["customer_location_id"] + `">` + customerData["data"][index]["name"] + `</option>`;
    }
    $('.selectpicker').append(options).selectpicker('refresh');

    create_button.click(async function(){
        let installed_firewalls = {
            customer_location_id : parseInt(select_customer_location.val()),
            customer_lan_on_site: customer_lan_on_site.val(),
            brand: brand.val(),
            equipment : equipment.val(),
            installed_date : new Date(installed_date.val()),
            firewall_ipv4_lan_address: firewall_ipv4_lan_address.val(),
            firewall_ipv6_lan_address:firewall_ipv6_lan_address.val(),
            current_version:current_version.val(),
            version_backup:version_backup.val(),
            backup_date: new Date(backup_date.val()),
            extra:extra.val(),
            internet_provider1: {
                wan_config_ipv4:{
                    firewall: firewall_wan1_ipv4.val(),
                    sub_net_mask: subnet_mask_wan1_ipv4.val(),
                    gateway: gateway_wan1_ipv4.val(),
                    ip_version:"4",
                    ip_range:ip_range.val()
                }
            },
            internet_provider2:{},
            created_by_user_id:parseInt($("#user_id").val())
        }
        let wan_config2_ipv4 = {
                firewall: firewall_wan2_ipv4.val(),
                sub_net_mask: subnet_mask_wan2_ipv4.val(),
                gateway: gateway_wan2_ipv4.val(),
                ip_version:"4",
            },
            wan_config2_ipv6 = {
                firewall: firewall_wan2_ipv6.val(),
                sub_net_mask: subnet_mask_wan2_ipv6.val(),
                gateway: gateway_wan2_ipv6.val(),
                ip_version:"6",
            },wan_config1_ipv6={
                firewall: firewall_wan1_ipv6.val(),
                sub_net_mask: subnet_mask_wan1_ipv6.val(),
                gateway: gateway_wan1_ipv6.val(),
                ip_version:"6",
            }
        if (checkWanConfigEmpty(wan_config1_ipv6) === true){
            installed_firewalls["internet_provider1"]["wan_config_ipv6"] = wan_config1_ipv6
        }

        if (checkWanConfigEmpty(wan_config2_ipv4) === true ){
            installed_firewalls["internet_provider2"]["wan_config_ipv4"] = wan_config2_ipv4
        }
        if (checkWanConfigEmpty(wan_config2_ipv6) === true){
            installed_firewalls["internet_provider2"]["wan_config_ipv6"] = wan_config2_ipv6
        }


        if (Object.keys(installed_firewalls["internet_provider2"]).length === 0) {
            delete installed_firewalls["internet_provider2"]
        }


        await postModel(
            JSON.stringify(installed_firewalls),
            "http://192.168.49.2:31471/api/v1/installed-firewalls"
        )
    });
});

function checkWanConfigEmpty(object){
    return object["firewall"] !== "" && object["sub_net_mask"] !== "" && object["gateway"] !== "";
}