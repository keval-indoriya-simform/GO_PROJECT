$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://localhost:8080/api/v1/softwares?page=1`)
    datatable = $('#softwareTable').DataTable({
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
            { data: 'customer_name' },
            { data: 'name' },
            { data: 'version' },
            { data: 'license_key' },
            { data: 'other_license_info' },
            { data: 'install_date' },
            { data: 'expiry_date' },
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://localhost:8080/api/v1/softwares?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#softwareTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["software_id"]) === false) {
                list.push(datatable.row(this).data()["software_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["software_id"])
        }
    });

    $('#softwareTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["software_id"])
        const myModal = new bootstrap.Modal('#modalUpdateSoftware', {})
        $("#modalUpdateSoftware").removeAttr("data-id")
        $("#modalUpdateSoftware").attr("data-id", datatable.row(this).data()["software_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://localhost:8080/api/v1/softwares?software_id="+list.toString(),
            `http://localhost:8080/network-management-solutions/list/softwares`)
    })

    $('#modalUpdateSoftware').on('show.bs.modal', async function() {
        let dataId = $("#modalUpdateSoftware").attr("data-id");
        let softwareResponse = await fetch('http://localhost:8080/api/v1/softwares?software_id=' + dataId);
        let softwareData = await softwareResponse.json()
        console.log(softwareData)

        $("#customer_location_name").val(softwareData["data"][0]["customer_locations"])
        $("#name").val(softwareData["data"][0]["name"])
        $("#version").val(softwareData["data"][0]["version"])
        $("#license_key").val(softwareData["data"][0]["license_key"])
        $("#other_license_info").val(softwareData["data"][0]["other_license_info"])
        $("#server_or_vm").val(softwareData["data"][0]["server_or_vm"])
        $("#install_date").val(new Date(softwareData["data"][0]["install_date"]).toISOString().split('T')[0])
        $("#expiry_date").val(new Date(softwareData["data"][0]["expiry_date"]).toISOString().split('T')[0])
        $("#notes").val(softwareData["data"][0]["notes"])

        let software_model = {}

        let update_button = $("#update-btn")
        update_button.click(async function(){
            await patchModel(
                JSON.stringify({
                    name: $("#name").val(),
                    version: $("#version").val(),
                    license_key: $("#license_key").val(),
                    other_license_info: $("#other_license_info").val(),
                    server_or_vm: $("#server_or_vm").val(),
                    install_date: new Date($("#install_date").val()),
                    expiry_date: new Date($("#expiry_date").val()),
                    notes: $("#notes").val(),
                    updated_by_user_id: parseInt($("#user_id").val()),
                }),
                "http://localhost:8080/api/v1/softwares?software_id="+dataId,
                `http://localhost:8080/api/v1/softwares?page=`
            )
        });
    });

    $('#modalUpdateSoftware').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});