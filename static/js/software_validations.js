let software_model={}

const validation = () => {
    'use strict'
    const forms = document.getElementsByClassName('needs-validation')
    let flag=false
    Array.from(forms).forEach(form => {
        if (!form.checkValidity()) {

        }else{
            flag=true
        }
        form.classList.add('was-validated')
    })

    return flag
}


const softwareValidation=()=>{
    'use strict'
    const flag=validation()

    if (flag===true){
        createSoftware(software_model).then(data=>{

            if (data.status=="201"){
                $('.needs-validation').trigger("reset");
                document.getElementsByClassName('needs-validation')[0].classList.remove("was-validated")
            }

        })
    }
}


const OnHandleChange=(key,val)=>{
    software_model[key]=val
}


$(document).ready(async function () {
    select_customer_location = $("#customer_location_name")
    const customerLocationResponse = await fetch('http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name&set_limit=false');
    const customerLocationData = await customerLocationResponse.json();
    options =""
    for (let index = 0; index < customerLocationData["data"].length; index++) {
        options += `<option data-tokens="`+customerLocationData["data"][index]["name"]+`" value="`+ customerLocationData["data"][index]["customer_location_id"] + `">` + customerLocationData["data"][index]["name"] + `</option>`;
    }
    select_customer_location.append(options).selectpicker('refresh');
})

const createSoftware=async(payload)=>{
    console.log(payload)
    payload["customer_location_id"] = parseInt($("#customer_location_name").val())
    payload["created_by_user_id"] = parseInt($("#user_id").val())
   return  await postModel(JSON.stringify(payload),"http://localhost:8080/api/v1/softwares")
}
