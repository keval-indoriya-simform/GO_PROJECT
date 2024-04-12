$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers?page=1`)
    datatable =$('#serverTable').DataTable({
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
            { data: 'host_name' },
            { data: 'customers' },
            { data: 'hardware_as_a_service' },
            { data: 'service_tag' },
            { data: 'expression_service_code' },
            { data: 'location' },
            { data: 'power_connect_type' },
            { data: 'warranty' },
            { data: 'type' },
            { data: 'purchase_date' },
            { data: 'expiration_date' },
            { data: 'days_left' },
            { data: 'ownership' },
            { data: 'order_number' },
            { data: 'os_platform' },
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#serverTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["server_id"]) === false) {
                list.push(datatable.row(this).data()["server_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["server_id"])
        }
    });

    $("#patch-data-btn").click(function () {
        const myModal = new bootstrap.Modal('#serverModal', {})
        console.log(list[0])
        $("#serverModal").removeAttr("data-id")
        $("#serverModal").attr("data-id", list[0])
        myModal.show()
    })

    $('#serverTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["server_id"])
        const myModal = new bootstrap.Modal('#serverModal', {})
        $("#serverModal").removeAttr("data-id")
        $("#serverModal").attr("data-id", datatable.row(this).data()["server_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers?server_id="+list.toString(),
            `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/network-management-solutions/list/servers`)
    })

    $('#serverModal').on('show.bs.modal', async function() {
        let dataId = $("#serverModal").attr("data-id");
        console.log(dataId)
        let serverResponse = await fetch('http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers?server_id=' + dataId);
        let serverData = await serverResponse.json()
        console.log(serverData)
        $("#validationHostName").val(serverData["data"][0]["host_name"])
        $("#validationCustomers").val(serverData["data"][0]["customers"])

        let option_hardware_as_a_service = ""
        if (serverData["data"][0]["hardware_as_a_service"] === false) {
            option_hardware_as_a_service = `<option selected value="false">No</option><option value="true">Yes</option>`
        } else {
            option_hardware_as_a_service = `<option value="false">No</option><option selected value="true">Yes</option>`
        }
        $("#validationHardwareAsService").html(option_hardware_as_a_service).selectpicker('refresh');
        $("#validationOsPlatform").val(serverData["data"][0]["os_platform"])
        $("#validationServiceTag").val(serverData["data"][0]["service_tag"])
        $("#validationLocation").val(serverData["data"][0]["location"])
        $("#validationWarranty").val(serverData["data"][0]["warranty"])
        $("#validationType").val(serverData["data"][0]["type"])
        $("#validationTypeRequired").val(serverData["data"][0]["power_connect_type"])
        $("#validationPurchaseDate").val(new Date(serverData["data"][0]["purchase_date"]).toISOString().split('T')[0])
        $("#validationExpirationDate").val(new Date(serverData["data"][0]["expiration_date"]).toISOString().split('T')[0])
        $("#validationDaysLeft").val(serverData["data"][0]["days_left"])
        $("#validationOwnership").val(serverData["data"][0]["ownership"])
        $("#validationOrderNumber").val(serverData["data"][0]["order_number"])
        $("#validationDescription").val(serverData["data"][0]["description"])
        $("#validationIdrac").val(serverData["data"][0]["idrac"])

        let update_button = $("#update-btn")
        update_button.click(async function(){
            if(checkExpirationDate(document.getElementById("validationPurchaseDate").value,document.getElementById("validationExpirationDate").value)) {
                if (document.getElementById("validationHardwareAsService").value === "false") {
                    flag = JSON.parse('false')
                } else {
                    flag = JSON.parse('true')
                }
                await patchModel(
                    JSON.stringify({
                        host_name: $("#validationHostName").val(),
                        hardware_as_a_service: flag,
                        os_platform: $("#validationOsPlatform").val(),
                        service_tag: $("#validationServiceTag").val(),
                        location: $("#validationLocation").val(),
                        warranty: $("#validationWarranty").val(),
                        power_connect_type: $("#validationType").val(),
                        type: $("#validationTypeRequired").val(),
                        purchase_date: new Date($("#validationPurchaseDate").val()),
                        expiration_date: new Date($("#validationExpirationDate").val()),
                        days_left: $("#validationDaysLeft").val(),
                        ownership: $("#validationOwnership").val(),
                        order_number: $("#validationOrderNumber").val(),
                        description: $("#validationDescription").val(),
                        idrac: $("#validationIdrac").val(),
                        updated_by_user_id: parseInt($("#user_id").val())
                    }),
                    "http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers?server_id=" + dataId,
                    `http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/api/v1/servers?page=`
                )
            } else{
            document.getElementById("invalid-date").innerHTML="Expiration date should be greater than purchase date"
        }
        });
        function checkExpirationDate(startDate,endDate){
            return new Date(endDate) > new Date(startDate);
        }
    });

    $('#serverModal').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});