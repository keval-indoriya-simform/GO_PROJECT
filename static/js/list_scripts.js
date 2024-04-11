async function delete_data(delete_url, reload_url){
    let deleteConfirm = confirm("Are you sure?");
    if (deleteConfirm === true) {
        try {
            const response = await fetch(delete_url , {
                method: "DELETE",
                mode: "cors",
                cache: "no-cache",
                headers: {
                    "Content-Type": "application/json",
                    "USER_ID": parseInt($("#user_id").val()),
                },
            });
            if (response.status != 204) {
            const resp = await response.json();
            if (resp["status"] === 201) {
                window.location = reload_url
            } else {
                console.log("error while deleting")
            }
            }
        } catch (error) {
            console.log("Error:", error);
        }
    }
}

function paginationLoad(pageNum, url) {
    pageStr = `<li id="prevPage" class="prev disable"><a id="linkPrev"><i class="fa-solid fa-angle-left"></i></a></li>`;
    pageStr += `<li class="active"><a onclick="pageReload('`+ url + `', `+1 + `, ` + pageNum +`)">`+ 1 +`</a></li>`;
    for (let i = 2; i <= pageNum; i++) {
        pageStr += `<li><a onclick="pageReload('`+ url + `', `+i + `, ` + pageNum +`)">`+ i +`</a></li>`;
    }
    pageStr += `<li id="nextPage" class="next"><a id="linkNext" onclick="nextPage('` + url + `', ` + pageNum +`)"><i class="fa-solid fa-angle-right"></i></a></li>`;
    $(".pagination").html(pageStr)
    if (pageNum == 1) {
        $("#nextPage")[0].classList.add("disable")
        $("#linkNext")[0].removeAttribute("onclick")
        $("#linkPrev")[0].removeAttribute("onclick")
    }
}

async function pageReload(url, pageNum, maxPageNum) {
    $("#nextPage")[0].classList.remove("disable")
    $("#prevPage")[0].classList.remove("disable")
    $("#linkNext")[0].removeAttribute("onclick")
    $("#linkPrev")[0].removeAttribute("onclick")
    if (pageNum != maxPageNum) {
        $("#linkNext").attr("onclick",`nextPage('` + url + `', ` + maxPageNum+ `)`);
    }
    if (pageNum != 1) {
        $("#linkPrev").attr("onclick", `previousPage('` + url + `', ` + maxPageNum+ `)`);
    }
    if (pageNum == maxPageNum) {
        $("#nextPage")[0].classList.add("disable")
    }
    if (pageNum == 1) {
        $("#prevPage")[0].classList.add("disable")
    }
    $(".pagination li").each( function () {
        this.classList.remove("active")
        if (this.firstChild.text == pageNum) {
            this.classList.add("active")
        }
    })
    let dataResponse = await retrieveModel(url+pageNum)
    datatable.clear().rows.add(dataResponse["data"]).draw();
}

async function previousPage(url, maxPageNum) {
    let page = parseInt($(".pagination .active").text())-1

    await pageReload(url, page, maxPageNum)
}

async function nextPage(url, maxPageNum) {
    let page = parseInt($(".pagination .active").text())+1
    await pageReload(url, page, maxPageNum)
}