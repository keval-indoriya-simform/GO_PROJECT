
const checkValidation = () => {
    'use strict'
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.getElementsByClassName('needs-validation')
    let flag = false
    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {

        if (!form.checkValidity()) {
            event.preventDefault()
            event.stopPropagation()
        } else {
            flag = true
        }

        form.classList.add('was-validated')

    })
    return flag
}

const cancelValidation = () => {
    'use strict'
    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')

    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
            event.preventDefault()
            event.stopPropagation()
            form.reset();
            form.classList.remove('was-validated')
    })
}


async function postModel(jsonObj, url){
    if (checkValidation() === true){
        const response = await fetch(url, {

            // Adding method type
            method: "POST",

            // Adding body or contents to send
            body: jsonObj,

            // Adding headers to the request
            headers: {
                "Content-type": "text/plain; charset=UTF-8"
            }
        });
        let data_response = await response.json()
        if (data_response["status"] === 201) {
            $.toast({
                type:"success",
                autoDismiss: true,
                message: 'Created Successful!'
            });
            $('.needs-validation').trigger("reset");
            document.getElementsByClassName('needs-validation')[0].classList.remove("was-validated")
            $(".selectpicker").selectpicker("refresh");
        } else {
            $.toast({
                type:"info",
                autoDismiss: true,
                message: data_response.message
            });
        }
        return data_response;
    } else {
        $.toast({
            type:"info",
            autoDismiss: true,
            message: "Data Validation Failed..!"
        });
    }
}

async function patchModel(jsonObj, patch_url, reload_url){
    if (checkValidation() === true) {
        const response = await fetch(patch_url, {

            // Adding method type
            method: "PATCH",

            // Adding body or contents to send
            body: jsonObj,

            // Adding headers to the request
            headers: {
                "Content-type": "text/plain; charset=UTF-8"
            }
        });
        if (response.status != 204) {
            const data1 = await response.json();
            if (data1["status"] === 201) {
                $.toast({
                    type: "success",
                    autoDismiss: true,
                    message: 'Updated Successful!'
                });
                $('.needs-validation').trigger("reset");
                document.getElementsByClassName('needs-validation')[0].classList.remove("was-validated")
                $(".selectpicker").selectpicker("refresh");
                $('.modal').modal('hide');
                let page = parseInt($(".pagination .active").text())
                let dataResponse = await retrieveModel(reload_url + page)
                datatable.clear().rows.add(dataResponse["data"]).draw();

            } else {
                $.toast({
                    type: "info",
                    autoDismiss: true,
                    message: data1.message
                });
            }
        } else {
            $.toast({
                type: "info",
                autoDismiss: true,
                message: "Data Validation Failed..!"
            });
        }
    } else {
        $.toast({
            type: "danger",
            autoDismiss: true,
            message: "No Data Available"
        });
    }
}

async function retrieveModel(url) {
    let data = await fetch(url, {

        // Adding method type
        method: "GET",

        // Adding headers to the request
        headers: {
            "Content-type": "text/plain; charset=UTF-8"
        }
    });
    if (data.status != 204) {
        return data.json()
    }
    else {
        return {"data" : {}, "total_pages" : 1}
    }

}
