    let cloud_private_ip_model={}
    const validation = () => {
        'use strict'
        const forms = document.getElementsByClassName('needs-validation')
        let flag = false
        // Loop over them and prevent submission
        Array.from(forms).forEach(form => {
            if (!form.checkValidity()) {

            } else {
                flag = true
            }
            form.classList.add('was-validated')
        })

        return flag
    }


    const createCloudPrivateIpModel = () => {
        'use strict'
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        const flag = validation()
        console.log($("#user_id").val())
        cloud_private_ip_model["created_by_user_id"] = parseInt($("#user_id").val())
        if (flag === true) {
            createCloudPrivateIp(cloud_private_ip_model).then(data => {

                if (data.status=="201"){
                    $('.needs-validation').trigger("reset");
                    document.getElementsByClassName('needs-validation')[0].classList.remove("was-validated")
                }
            })
        }
    }


    const OnHandleChange = (key, val) => {
        cloud_private_ip_model[key] = val
    }


    const createCloudPrivateIp = async (payload) => {
        console.log(payload)
        if (select_users.val() == null) {
            payload["assigned_to_id"] = parseInt($("#user_id").val())
        } else {
            payload["assigned_to_id"] = parseInt(select_users.val())
        }
        return await postModel(JSON.stringify(payload),"http://192.168.49.2:31471/api/v1/cloud-private-ips")

    }

    $(document).ready(async function() {
        select_users = $("#assign_to")
        const userResponse = await fetch('http://192.168.49.2:31471/api/v1/users');
        const UserData = await userResponse.json();
        options =""
        for (index = 0; index < UserData["data"].length; index++) {
            options += `<option data-tokens="`+UserData["data"][index]["name"]+`" value="`+ UserData["data"][index]["user_id"] + `">` + UserData["data"][index]["name"] + `</option>`;
        }
        select_users.append(options).selectpicker('refresh');
    });