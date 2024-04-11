$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id&append_select=true&page=1`)
    datatable = $('#customerLocationTable').DataTable({
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
            {
                data: 'customers',
                title: 'Customers',
                className: "popup text-center",
                render: function (data) {
                    return '<span class="popup-element">' + data + '</span>';
                }
            },
            {
                data: 'name',
                title: 'Name',
                className: "text-center"
            }
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id&append_select=true&page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#customerLocationTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["customer_location_id"]) === false) {
                list.push(datatable.row(this).data()["customer_location_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["customer_location_id"])
        }
    });

    $('#customerLocationTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["customer_location_id"])
        const myModal = new bootstrap.Modal('#modalCustomerLocation', {})
        $("#modalCustomerLocation").removeAttr("data-id")
        $("#modalCustomerLocation").attr("data-id", datatable.row(this).data()["customer_location_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://localhost:8080/api/v1/customer-locations?customer_location_id="+list.toString(),
            `http://localhost:8080/network-management-solutions/list/customer-locations`)
    })

    $('#modalCustomerLocation').on('show.bs.modal', async function() {
        let dataId = $("#modalCustomerLocation").attr("data-id");
        console.log(dataId)
        let customerLocationResponse = await fetch('http://localhost:8080/api/v1/customer-locations' +
            '?customer_location_id=' + dataId+'&select_column=customer_locations.is_primary,' +
            'customer_locations.description&append_select=true');
        let customerLocationData = await customerLocationResponse.json()
        console.log(customerLocationData)
        $("#validationCustom05").val(customerLocationData["data"][0]["customers"])
        $("#validationCustom02").val(customerLocationData["data"][0]["name"])
        $("#validationTextarea").val(customerLocationData["data"][0]["description"])

        let update_button = $("#update-btn")
        update_button.click(async function(){
            await patchModel(
                JSON.stringify({
                    name:  $("#validationCustom02").val(),
                    description: $("#validationTextarea").val(),
                    // is_primary : customer_is_primary.is(':checked'),
                    updated_by_user_id: parseInt($("#user_id").val())
                }),
                "http://localhost:8080/api/v1/customer-locations?customer_location_id="+dataId,
                `http://localhost:8080/api/v1/customer-locations?select_column=customer_locations.customer_location_id&append_select=true&page=`
            )
        });
    });

    $('#modalCustomerLocation').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });

    $('#customerLocationTable').on('click','.popup', function(event) {
        if (!$(event.target).hasClass("popup-element")) {
            $(".popup-content").css("display", "flex");
        }
        var clickedData = $(this).text();
        fetch('/api/v1/customers')
            .then(response => response.json())
            .then(response => {
                console.log(response);
                [response].forEach(item => {
                    for (var i = 0; i < item.data.length; i++) {
                        console.log(item.data[i])
                        if (clickedData === item.data[i]["name"]) {
                            $("#name").text(item.data[i].name)
                            $("#is_active").text(item.data[i].is_active)
                            $("#cloud_or_onsite").text(item.data[i].cloud_or_onsite)
                        }
                    }
                });
            });
        });
        $(".close").click(function (){
            $(".popup-content").css("display", "none");
        });
    });
