let note_model={}

$(document).ready(async function() {
    submit_button = $("#submit-btn")
    select_customer_location = $("#customer_location_name")
    const customerLocationResponse = await fetch('http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id,customer_locations.name&set_limit=false');
    const customerLocationData = await customerLocationResponse.json();
    options =""
    for (index = 0; index < customerLocationData["data"].length; index++) {
        options += `<option data-tokens="`+customerLocationData["data"][index]["name"]+`" value="`+ customerLocationData["data"][index]["customer_location_id"] + `">` + customerLocationData["data"][index]["name"] + `</option>`;
    }
    select_customer_location.append(options).selectpicker('refresh');

    select_users = $("#assign_to")
    const userResponse = await fetch('http://localhost:8080/api/v1/users');
    const UserData = await userResponse.json();
    options =""
    for (index = 0; index < UserData["data"].length; index++) {
        options += `<option data-tokens="`+UserData["data"][index]["name"]+`" value="`+ UserData["data"][index]["user_id"] + `">` + UserData["data"][index]["name"] + `</option>`;
    }
    select_users.append(options).selectpicker('refresh');

    submit_button.click(async function(){
        note_model["subject"] = $("#subject").val()
        note_model["created_by_user_id"] = parseInt($("#user_id").val())
        note_model["customer_location_id"] = parseInt(select_customer_location.val())
        note_model["attachment"] = $("#attachment").val()
        note_model["note"] = $("#note").val()
        if (select_users.val() == null) {
            note_model["assigned_to_id"] = parseInt($("#user_id").val())
        } else {
            note_model["assigned_to_id"] = parseInt(select_users.val())
        }

        await postModel(
            JSON.stringify(note_model),
            "http://localhost:8080/api/v1/notes"
        )
    });
});