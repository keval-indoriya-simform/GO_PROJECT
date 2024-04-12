$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?page=1`)
    datatable = $('#cloudPrivateIPTable').DataTable({
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
        order: [2 , "asc"],
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
            { data: 'cloud_vm_name' },
            { data: 'name' },
            { data: 'ipv4_assignment' },
            { data: 'ipv6_assignment' },

        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#cloudPrivateIPTable').on('click','.editor-select', async function(){
        check = this.firstChild
        console.log($(check).is(':checked'))
        if ($(check).is(':checked') === true) {
            console.log(datatable.row(this).data()["cloud_private_ip_id"])
            if (list.includes(datatable.row(this).data()["cloud_private_ip_id"]) === false) {
                list.push(datatable.row(this).data()["cloud_private_ip_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["cloud_private_ip_id"])
        }
    });

    $("#delete-btn").click(function () {
        delete_data("http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?cloud_private_ip_id="+list.toString(),
            'http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/network-management-solutions/list/cloud-private-ips');
    })

    $('#cloudPrivateIPTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["cloud_private_ip_id"])
        const myModal = new bootstrap.Modal('#modalCloudPrivateIP', {})
        $("#modalCloudPrivateIP").removeAttr("data-id")
        $("#modalCloudPrivateIP").attr("data-id", datatable.row(this).data()["cloud_private_ip_id"])
        myModal.show()
    });

    $('#modalCloudPrivateIP').on('show.bs.modal', async function() {
        let dataId = $("#modalCloudPrivateIP").attr("data-id");
        console.log(dataId)
        let cloudPrivateIpResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?cloud_private_ip_id=' + dataId);
        let cloudPrivateIpData = await cloudPrivateIpResponse.json()
        console.log(cloudPrivateIpData)
        $("#ipv4_assignment").val(cloudPrivateIpData["data"][0]["ipv4_assignment"])
        $("#ipv6_assignment").val(cloudPrivateIpData["data"][0]["ipv6_assignment"])
        $("#description").val(cloudPrivateIpData["data"][0]["description"])

        let userResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/users?set_limit=false');
        let UserData = await userResponse.json();
        let selectedValue = cloudPrivateIpData["data"][0]["assigned_to_id"]
        let options = ""
        for (index = 0; index < UserData["data"].length; index++) {
            if (selectedValue != UserData["data"][index]["user_id"]){
                options += `<option data-tokens="` + UserData["data"][index]["name"] + `" value="` + UserData["data"][index]["user_id"] + `">` + UserData["data"][index]["name"] + `</option>`;
            } else {
                options += `<option selected data-tokens="` + UserData["data"][index]["name"] + `" value="` + UserData["data"][index]["user_id"] + `">` + UserData["data"][index]["name"] + `</option>`;
            }
        }
        $("#assign_to").html(options).selectpicker('refresh');

        let update_button = $("#update-btn")
        update_button.click(async function(){
            if ($("#assign_to").val() === null) {
                assignID= parseInt($("#user_id").val())
            } else {
                assignID = parseInt($("#assign_to").val())
            }
            await patchModel(
                JSON.stringify({
                    ipv4_assignment:  $("#ipv4_assignment").val(),
                    ipv6_assignment: $("#ipv6_assignment").val(),
                    description: $("#description").val(),
                    assigned_to_id: assignID,
                    updated_by_user_id: parseInt($("#user_id").val()),
                }),
                "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?cloud_private_ip_id="+dataId,
                'http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/cloud-private-ips?page='
            )
        });
    });

    $('#modalCloudPrivateIP').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});