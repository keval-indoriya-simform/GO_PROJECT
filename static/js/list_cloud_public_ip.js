$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://192.168.49.2:31471/api/v1/cloud-public-ips?page=1`)
    datatable = $('#cloudPublicIpTable').DataTable({
        data: retrieveModelData["data"],
        deferRender: true,
        scrollY:        "60vh",
        scrollX:        "80%",
        scrollCollapse: true,
        scroller:       true,
        responsive: true,
        paging: false,
        fixedColumns: {
            leftColumns: 2,
        },
        order: [2, "asc"],
        columns: [
            {
                data: null,
                width: "20px",
                className: "dt-center editor-select",
                defaultContent: '<input class="bulk-select" type="checkbox"/>',
                orderable: false
            },
            {
                data: null,
                width: "20px",
                className: "dt-center editor-edit",
                defaultContent: '<i class="fa fa-pencil"/>',
                orderable: false
            },
            { data: 'ip_address' },
            { data: 'customer_location' },
            { data: 'post_forward_ip' },
            { data: 'cloud_vm_name' },
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://192.168.49.2:31471/api/v1/cloud-public-ips?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#cloudPublicIpTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["cloud_public_ip_id"]) === false) {
                list.push(datatable.row(this).data()["cloud_public_ip_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["cloud_public_ip_id"])
        }
    });

    $('#cloudPublicIpTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["cloud_public_ip_id"])
        const myModal = new bootstrap.Modal('#cloudPublicIpModal', {})
        $("#cloudPublicIpModal").removeAttr("data-id")
        $("#cloudPublicIpModal").attr("data-id", datatable.row(this).data()["cloud_public_ip_id"])
        myModal.show()
    });

    $("#delete-data-btn").click(function () {
        delete_data("http://192.168.49.2:31471/api/v1/cloud-public-ips?cloud_public_ip_id="+list.toString(),
            `http://192.168.49.2:31471/network-management-solutions/list/cloud-public-ips` )
    })

    $('#cloudPublicIpModal').on('show.bs.modal', async function() {
        let dataId = $("#cloudPublicIpModal").attr("data-id");
        console.log(dataId)
        let cloudPublicIpResponse = await fetch('http://192.168.49.2:31471/api/v1/cloud-public-ips?cloud_public_ip_id=' + dataId);
        let cloudPublicIpData = await cloudPublicIpResponse.json()
        console.log(cloudPublicIpData)
        let postForwardResponse = await fetch('http://192.168.49.2:31471/api/v1/cloud-private-ips?set_limit=false');
        let postForwardData = await postForwardResponse.json()
        console.log(postForwardData)
        let selectedValue = cloudPublicIpData["data"][0]["post_forward_ip"]
        let options =""
        for (let index = 0; index < postForwardData["data"].length; index++) {
            console.log(selectedValue !== postForwardData["data"][index]["ipv4_assignment"])
            console.log(selectedValue)
            console.log(postForwardData["data"][0]["ipv4_assignment"])
            if (selectedValue !== postForwardData["data"][index]["ipv4_assignment"]){
                options += `<option data-tokens="`+postForwardData["data"][index]["ipv4_assignment"]+`" value="`+ postForwardData["data"][index]["cloud_private_ip_id"] + `">` + postForwardData["data"][index]["ipv4_assignment"] + `</option>`
            } else {
                options += `<option selected data-tokens="`+postForwardData["data"][index]["ipv4_assignment"]+`" value="`+ postForwardData["data"][index]["cloud_private_ip_id"] + `">` + postForwardData["data"][index]["ipv4_assignment"] + `</option>`
            }
        }
        $('#validationPostForwardIp').append(options).selectpicker('refresh');

        $("#validationCloudPublicIps").val(cloudPublicIpData["data"][0]["ip_address"])
        $("#validationCustomerLocation").val(cloudPublicIpData["data"][0]["customer_location"])
        $("#validationCloudVmName").val(cloudPublicIpData["data"][0]["cloud_vm_name"])

        let update_button = $("#update-btn")
        update_button.click(async function(){
            await patchModel(
                JSON.stringify({
                    ip_address:  $("#validationCloudPublicIps").val(),
                    post_forward_ip: $("#validationPostForwardIp option:selected").text(),
                    cloud_vm_name : $("#validationCloudVmName").val(),
                    updated_by_user_id: parseInt($("#user_id").val())
                }),
                "http://192.168.49.2:31471/api/v1/cloud-public-ips?cloud_public_ip_id="+dataId,
                `http://192.168.49.2:31471/api/v1/cloud-public-ips?page=`
            )
        });
    });

    $('#cloudPublicIpModal').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});