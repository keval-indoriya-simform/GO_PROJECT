let customer_model={}

$(document).ready(async function() {
    submit_button = $("#submit-btn"),
    cloudOrOnsite = $("#cloud_or_onsite")
    active_check = $("#active_check")
    voip_check = $("#voip_check")
    internet_check = $("#internet_check")
    firewall_check = $("#firewall_check")
    hardware_as_a_service_check = $("#hardware_as_a_service_check")
    const cloudOrOnsiteResponse = await fetch('http://localhost:8080/api/v1/customers/cloud-or-onsites');
    const cloudOrOnsiteData = await cloudOrOnsiteResponse.json();
    options =""
    for (index = 0; index < cloudOrOnsiteData["data"].length; index++) {
        options += `<option data-tokens="`+cloudOrOnsiteData["data"][index]["name"]+`" value="`+ cloudOrOnsiteData["data"][index]["cloud_or_onsite_id"] + `">` + cloudOrOnsiteData    ["data"][index]["name"] + `</option>`;
    }
    $('.selectpicker').append(options).selectpicker('refresh');
    submit_button.click(async function(){
        customer_model["name"] = $("#name").val()
        customer_model["description"] = $("#description").val()
        customer_model["backup_software"] = $("#backup_software").val()
        customer_model["cloud_or_onsite_id"] = parseInt(cloudOrOnsite.val())
        customer_model["is_active"] = active_check.is(':checked')
        customer_model["voip"] = voip_check.is(':checked')
        customer_model["internet"] = internet_check.is(':checked')
        customer_model["firewall"] = firewall_check.is(':checked')
        customer_model["hardware_as_a_service"] = hardware_as_a_service_check.is(':checked')
        customer_model["created_by_user_id"] = parseInt($("#user_id").val())
        await postModel(
            JSON.stringify(customer_model),
            "http://localhost:8080/api/v1/customers"
        )
    });

});