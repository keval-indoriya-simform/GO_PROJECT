const getCloudPublicIpForm= document.getElementById('cloud-public-ip-form')
getCloudPublicIpForm.addEventListener('submit',async (e) =>{
    'use strict'
    e.preventDefault()
    let response;
    let cloud_public_ip_model= {};
    try {
        cloud_public_ip_model=getCloudPublicIpFormData(cloud_public_ip_model)
        cloud_public_ip_model["created_by_user_id"] = parseInt($("#user_id").val())
        console.log(cloud_public_ip_model)
        response = await postModel(
            JSON.stringify(cloud_public_ip_model),
            "http://localhost:8080/api/v1/cloud-public-ips"
        )
        return response.json()
    } catch (error) {
        console.error(error)
    }

});

function getCloudPublicIpFormData(cloud_public_ip_model){
    cloud_public_ip_model["ip_address"]=document.getElementById("validationCloudPublicIps").value
    cloud_public_ip_model["customer_location_id"]=parseInt($("#customer_location_name").val())
    cloud_public_ip_model["post_forward_ip"]=$("#validationPostForwardIp option:selected").text()
    cloud_public_ip_model["cloud_vm_name"]=document.getElementById("validationCloudVmName").value
    return cloud_public_ip_model
}

$(document).ready(async function(){
    select_customer_location = $("#customer_location_name")
    const customerLocationResponse = await fetch('http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name&set_limit=false');
    const customerLocationData = await customerLocationResponse.json();
    options =""
    for (let index = 0; index < customerLocationData["data"].length; index++) {
        options += `<option data-tokens="`+customerLocationData["data"][index]["name"]+`" value="`+ customerLocationData["data"][index]["customer_location_id"] + `">` + customerLocationData["data"][index]["name"] + `</option>`;
    }
    select_customer_location.append(options).selectpicker('refresh');

    select_post_forward_ip = $("#validationPostForwardIp")
    const postForwardResponse = await fetch('http://localhost:8080/api/v1/cloud-private-ips');
    const postForwardData = await postForwardResponse.json();
    console.log(postForwardData)
    let options_post_forward =""
    for (let index = 0; index < postForwardData["data"].length; index++) {
        options_post_forward += `<option data-tokens="`+postForwardData["data"][index]["ipv4_assignment"]+`" value="`+ postForwardData["data"][index]["cloud_private_ip_id"] + `">` + postForwardData["data"][index]["ipv4_assignment"] + `</option>`;
    }
    select_post_forward_ip.append(options_post_forward).selectpicker('refresh');
});