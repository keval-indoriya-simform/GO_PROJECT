const validation = () => {
    'use strict'
    const forms = document.getElementsByClassName('needs-validation')
    let flag=false
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
        if (!form.checkValidity()) {

        }else{
            flag=true
        }
        form.classList.add('was-validated')
    })

    return flag
}

$(document).ready(async function() {
    $("#register").click(async function(){
        if (validation()) {
            let response = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/users/register', {

                // Adding method type
                method: "POST",

                // Adding body or contents to send
                body: JSON.stringify({
                    name: $("#fullname").val(),
                    email: $("#email").val(),
                    contact: parseInt($("#contact").val()),
                    department: $("#department").val(),
                    username: $("#username").val(),
                    is_active: false,
                }),

                // Adding headers to the request
                headers: {
                    "Content-type": "application/json; charset=UTF-8"
                }
            });
            let data_response = await response.json()
            if (data_response["status"] === 201) {
                $.toast({
                    type: "success",
                    autoDismiss: true,
                    message: 'Register Successful!'
                });
                // window.location = "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/network-management-solutions/list/customer-locations"
            } else {
                $.toast({
                    type: "info",
                    autoDismiss: true,
                    message: "Register Failed"
                });
            }
        } else {
            $.toast({
                type: "info",
                autoDismiss: true,
                message: "data validation failed"
            });
        }
    });
});