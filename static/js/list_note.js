$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://192.168.49.2:31471/api/v1/notes?page=1`)
    datatable = $('#noteTable').DataTable({
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
            { data: 'subject', title: 'Subject' },
            { data: 'customer_location', title: 'Related To' },
            { data: 'note', title: 'Notes' },
            { data: 'attachment', title: 'Attachment' },
            { data: 'created_by', title: 'Created By' },
            { data: 'created_at', title: 'Date Created' }
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://192.168.49.2:31471/api/v1/notes?page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#noteTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["note_id"]) === false) {
                list.push(datatable.row(this).data()["note_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["note_id"])
        }
    });

    $('#noteTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["note_id"])
        const myModal = new bootstrap.Modal('#modelNote', {})
        $("#modelNote").removeAttr("data-id")
        $("#modelNote").attr("data-id", datatable.row(this).data()["note_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://192.168.49.2:31471/api/v1/notes?note_id="+list.toString(),
            `http://192.168.49.2:31471/network-management-solutions/list/notes`)
    })

    $('#modelNote').on('show.bs.modal', async function() {
        let dataId = $("#modelNote").attr("data-id");
        let noteResponse = await fetch('http://192.168.49.2:31471/api/v1/notes?note_id=' + dataId);
        let noteData = await noteResponse.json()
        console.log(noteData)
        select_users = $("#assign_to")
        let userResponse = await fetch('http://192.168.49.2:31471/api/v1/users?set_limit=false');
        let UserData = await userResponse.json();
        console.log(UserData)
        let options =""
        $("#subject").val(noteData["data"][0]["subject"])
        $("#customer_location_name").val(noteData["data"][0]["customer_location"])
        $("#fileName").text(noteData["data"][0]["attachment"])
        let selectedValue = noteData["data"][0]["assigned_to_id"]
        for (index = 0; index < UserData["data"].length; index++) {
            if (selectedValue !== UserData["data"][index]["user_id"]){
                options += `<option data-tokens="` + UserData["data"][index]["name"] + `" value="` + UserData["data"][index]["user_id"] + `">` + UserData["data"][index]["name"] + `</option>`;
            } else {
                options += `<option selected data-tokens="` + UserData["data"][index]["name"] + `" value="` + UserData["data"][index]["user_id"] + `">` + UserData["data"][index]["name"] + `</option>`;
            }
        }
        select_users.html(options).selectpicker('refresh');

        // $("#attachment").val(noteData["data"][0]["attachment"])
        $("#attachment").on('change', function () {
            $("#fileName").text("")
        })
        $("#note").val(noteData["data"][0]["note"])

        let update_button = $("#update-btn")
        update_button.click(async function(){
            let note_model={}
            note_model["subject"] = $("#subject").val()
            note_model["updated_by_user_id"] = parseInt($("#user_id").val())
            note_model["customer_location_id"] = parseInt($("#customer_location_name").val())
            note_model["note"] = $("#note").val()
            if (select_users.val() === null) {
                note_model["assigned_to_id"] = parseInt($("#user_id").val())
            } else {
                note_model["assigned_to_id"] = parseInt(select_users.val())
            }
            if ( $("#attachment").val()  === "" ) {
                note_model["attachment"] = $("#fileName").text()
            } else {
                note_model["attachment"] = $("#attachment").val()
            }
            await patchModel(
                JSON.stringify(note_model),
                "http://192.168.49.2:31471/api/v1/notes?note_id="+dataId,
                `http://192.168.49.2:31471/api/v1/notes?page=`
            )
        });
    });

    $('#modelNote').on('hide.bs.modal', function() {
        $("#update-btn").unbind()
    });
});