const getServerForm = document.getElementById("server-form")
getServerForm.addEventListener("submit",async (e)=>{
    'use strict'
    e.preventDefault();
    if(checkExpirationDate(document.getElementById("validationPurchaseDate").value,document.getElementById("validationExpirationDate").value)){
        let response, server_model={};
        try{
            server_model=getServerFormData(server_model)
            response = await postModel(
                JSON.stringify(server_model),
                "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers"
            )
            return response.json()
        }catch(error){
            console.error(error)
        }
    }else{
        document.getElementById("invalid-date").innerHTML="Expiration date should be greater than purchase date"
    }
});

function checkExpirationDate(startDate,endDate){
    return new Date(endDate) > new Date(startDate);
}

$(document).ready(async function(){
    select_customer_location = $("#customer_location_name")
    const customerLocationResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name&set_limit=false');
    const customerLocationData = await customerLocationResponse.json();
    options =""
    for (let index = 0; index < customerLocationData["data"].length; index++) {
        options += `<option data-tokens="`+customerLocationData["data"][index]["name"]+`" value="`+ customerLocationData["data"][index]["customer_location_id"] + `">` + customerLocationData["data"][index]["name"] + `</option>`;
    }
    select_customer_location.append(options).selectpicker('refresh');
});

function getServerFormData(server_model){
    server_model["host_name"]=document.getElementById("validationHostName").value
    server_model["customer_location_id"]=parseInt(document.getElementById("customer_location_name").value)
    if(document.getElementById("validationHardwareAsService").value==="false"){
        server_model["hardware_as_a_service"]=JSON.parse('false')
    }else{
        server_model["hardware_as_a_service"]=JSON.parse('true')
    }
    server_model["os_platform"]=document.getElementById("validationOsPlatform").value
    server_model["service_tag"]=document.getElementById("validationServiceTag").value
    server_model["location"]=document.getElementById("validationLocation").value
    server_model["warranty"]=document.getElementById("validationWarranty").value
    server_model["power_connect_type"]=document.getElementById("validationType").value
    server_model["type"]=document.getElementById("validationTypeRequired").value
    server_model["purchase_date"]=new Date(document.getElementById("validationPurchaseDate").value)
    server_model["expiration_date"]=new Date(document.getElementById("validationExpirationDate").value)
    server_model["days_left"]=document.getElementById("validationDaysLeft").value
    server_model["ownership"]=document.getElementById("validationOwnership").value
    server_model["order_number"]=document.getElementById("validationOrderNumber").value
    server_model["description"]=document.getElementById("validationDescription").value
    server_model["idrac"]=document.getElementById("validationIdrac").value
    server_model["created_by_user_id"] = parseInt($("#user_id").val())
    return server_model
}