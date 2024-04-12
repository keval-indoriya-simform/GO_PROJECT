function passwordVisibility() {
    var passwordField = document.querySelector(".password-field");
    var passwordToggle = document.querySelector(".password-toggle");

    if (passwordField.type === "password") {
        passwordField.type = "text";
        passwordToggle.innerHTML = '<i class="fas fa-eye-slash"></i>';
    } else {
        passwordField.type = "password";
        passwordToggle.innerHTML = '<i class="fas fa-eye"></i>';
    }
}
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
const loginValidation=()=>{
    'use strict'
    const flag=validation()

    if (flag===true){
        createCustomer(customer_model).then(data=>{
            $('.needs-validation').trigger("reset");
            $('.needs-validation').removeClass('was-validated')
        } )
    }
    return flag
}

$(document).ready(async function() {
    $("#login").click(async function(){
                let response = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/login', {

                    // Adding method type
                    method: "POST",

                    // Adding body or contents to send
                    body: JSON.stringify({
                        username : $("#username").val(),
                        password : $("#password").val()
                    }),

                    // Adding headers to the request
                    headers: {
                        "Content-type": "text/plain; charset=UTF-8"
                    }
                });
                let data_response = await response.json()
                if (data_response["status"] === 200) {
                    $.toast({
                        type:"success",
                        autoDismiss: true,
                        message: 'login Successful!'
                    });
                    window.location = "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/network-management-solutions/list/customer-locations"
                } else {
                    $.toast({
                        type: "info",
                        autoDismiss: true,
                        message: data_response.message
                    });
                    $("#spanIncorrectPassword")[0].classList.remove("d-none")
                    $("#spanIncorrectPassword")[0].classList.add("d-block")
                }
    });
});