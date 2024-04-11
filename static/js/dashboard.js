$(document).ready(function (){
    $(".widget").on('click', function () {
        var clickedData = $(this)[0].firstElementChild.firstElementChild.firstElementChild.innerText;
        const myModal = new bootstrap.Modal('#dashboardModal', {})
        $("#dashboardModal").removeAttr("data-id")
        $("#dashboardModal").attr("data-id", clickedData)
        myModal.show();
    });
    $('#dashboardModal').on('show.bs.modal', async function() {
        let dataId = $("#dashboardModal").attr("data-id");
        $("#heading").text(dataId+" Count")
    });
});


