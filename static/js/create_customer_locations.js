let create_button = $("#submit-btn"),
    select_customer = $("#validationCustom05"),
    customer_location_name = $("#validationCustom02"),
    customer_location_description = $("#validationTextarea"),
    customer_is_primary = $("#flexSwitchCheckChecked")

$(document).ready(async function() {
    const customerResponse = await fetch('http://192.168.49.2:31471/api/v1/customers?select_column=customers.customer_id,customers.name&set_limit=false');
    const customerData = await customerResponse.json();
    options =""
    for (index = 0; index < customerData["data"].length; index++) {
        options += `<option data-tokens="`+customerData["data"][index]["name"]+`" value="`+ customerData["data"][index]["customer_id"] + `">` + customerData["data"][index]["name"] + `</option>`;
    }
    $('.selectpicker').append(options).selectpicker('refresh');
    create_button.click(async function(){
        console.log(customer_is_primary.is(':checked'))
        await postModel(
            JSON.stringify({
                customer_id : parseInt(select_customer.val()),
                name: customer_location_name.val(),
                description: customer_location_description.val(),
                is_primary : customer_is_primary.is(':checked'),
                created_by_user_id: parseInt($("#user_id").val()) 
            }),
            "http://192.168.49.2:31471/api/v1/customer-locations"
        )
    });
});