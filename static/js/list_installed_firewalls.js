$(document).ready(async function(){
    let retrieveModelData = await retrieveModel(`http://localhost:8080/api/v1/installed-firewalls?select_column=installed_firewalls.installed_firewall_id&append_select=true&page=1`)
    datatable = $('#installedFirewallTable').DataTable({

        data: retrieveModelData["data"],
        deferRender: true,
        scrollY:        "60vh",
        scrollX:        "80%",
        scrollCollapse: true,
        scroller:       true,
        responsive: true,
        paging: false,
        order: [2, "asc"],
        fixedColumns: {
            leftColumns: 2
        },
        columns: [
            {
                data: null,
                width: "40px",
                className: "dt-center editor-select",
                defaultContent: '<input class="bulk-select" type="checkbox"/>',
                orderable: false
            },
            {
                data: null,
                width: "40px",
                className: "dt-center editor-edit",
                defaultContent: '<i class="fa fa-pencil"/>',
                orderable: false
            },
            {
                data: 'customer_locations',
                title: 'Customer Locations',
                className: "text-center"
            },
            {
                data: 'customer_lan_on_site',
                title: 'Customer LAN (on-site)',
                className: "text-center"
            },
            {
                data: 'brand',
                title: 'Brand',
                className: "text-center"
            },
            {
                data: 'current_version',
                title: 'Current Version',
                className: "text-center"
            },
            {
                data: 'version_backup',
                title: 'Version Backup',
                className: "text-center"
            },
            {
                data: 'backup_date',
                title: 'BackUp Date',
                className: "text-center"
            },
            {
                data: 'equipment',
                title: 'Equipment',
                className: "text-center"
            },
            {
                data: 'installed_date',
                title: 'Install Date',
                className: "text-center"
            },
            {
                data: 'firewall_wan1_ipv4',
                title: 'FireWall WAN-1(ipv4)',
                className: "text-center"
            },
            {
                data: 'sub_net_mask_wan1_ipv4',
                title: 'Subnet Mask WAN-1 IPv4',
                className: "text-center"
            },
            {
                data: 'gateway_wan1_ipv4',
                title: 'GATEWAY WAN-1(ipv4)',
                className: "text-center"
            },
            {
                data: 'firewall_wan1_ipv6',
                title: 'FireWall WAN-1(ipv6)',
                className: "text-center"
            },
            {
                data: 'sub_net_mask_wan1_ipv6',
                title: 'Subnet Mask WAN-1 IPv6',
                className: "text-center"
            },
            {
                data: 'gateway_wan1_ipv6',
                title: 'GATEWAY WAN- 1(ipv6)',
                className: "text-center"
            },
            {
                data: 'firewall_wan2_ipv4',
                title: 'FireWall WAN-2(ipv4)',
                className: "text-center"
            },
            {
                data: 'sub_net_mask_wan2_ipv4',
                title: 'Subnet Mask WAN-2 IPv4',
                className: "text-center"
            },
            {
                data: 'gateway_wan2_ipv4',
                title: 'GATEWAY WAN-2(ipv4)',
                className: "text-center"
            },
            {
                data: 'firewall_wan2_ipv6',
                title: 'FireWall WAN-2(ipv6)',
                className: "text-center"
            },
            {
                data: 'sub_net_mask_wan2_ipv6',
                title: 'Subnet Mask WAN-2 IPv6',
                className: "text-center"
            },
            {
                data: 'gateway_wan2_ipv6',
                title: 'GATEWAY WAN-2(ipv6)',
                className: "text-center"
            },
            {
                data: 'firewall_ipv4_lan_address',
                title: 'Firewall IPv4 LAN address',
                className: "text-center"
            },
            {
                data: 'firewall_ipv6_lan_address',
                title: 'Firewall IPv6 LAN address',
                className: "text-center"
            },
            {
                data: 'extra',
                title: 'Extra',
                className: "text-center"
            }
        ],
    });

    paginationLoad(retrieveModelData["total_pages"], `http://localhost:8080/api/v1/installed-firewalls?select_column=installed_firewalls.installed_firewall_id&append_select=true&page=`)

    function RemoveElementFromArray(number) {
        this.list.forEach((value,index)=>{
            if(value==number) this.list.splice(index,1);
        });
    }
    list = []

    $('#installedFirewallTable').on('click','.editor-select', async function(){
        check = this.firstChild
        if ($(check).is(':checked') == true) {
            if (list.includes(datatable.row(this).data()["installed_firewall_id"]) === false) {
                list.push(datatable.row(this).data()["installed_firewall_id"])
            }
        } else {
            RemoveElementFromArray(datatable.row(this).data()["installed_firewall_id"])
        }
    });

    $('#installedFirewallTable').on('click','.editor-edit', function(){
        console.log(datatable.row(this).data()["installed_firewall_id"])
        const myModal = new bootstrap.Modal('#modalInstalledFirewall', {})
        $("#modalInstalledFirewall").removeAttr("data-id")
        $("#modalInstalledFirewall").attr("data-id", datatable.row(this).data()["installed_firewall_id"])
        myModal.show()
    });

    $("#delete-btn").click(function () {
        delete_data("http://localhost:8080/api/v1/installed-firewalls?installed_firewall_id="+list.toString(),
            `http://localhost:8080/network-management-solutions/list/installed-firewalls`)
    })

    $('#modalInstalledFirewall').on('show.bs.modal', async function() {
        let dataId = $("#modalInstalledFirewall").attr("data-id");
        console.log(dataId)
        let installedFirewallResponse = await fetch('http://localhost:8080/api/v1/installed-firewalls/?select_column=installed_firewalls.installed_firewall_id,wan_configs1_ipv4.ip_range&append_select=true&installed_firewall_id=' + dataId);
        let installedFirewallData = await installedFirewallResponse.json()
        console.log(installedFirewallData)

        let select_customer_location = $("#validationCustom05"),
            customer_lan_on_site = $("#validationCustom02"),
            brand = $("#validationCustom03"),
            equipment = $("#validationCustom04"),
            installed_date = $("#datepicker-1"),
            firewall_wan1_ipv4 = $("#validationCustom06"),
            ip_range = $("#validationCustom07"),
            subnet_mask_wan1_ipv4 = $("#validationCustom08"),
            gateway_wan1_ipv4 = $("#validationCustom09"),
            firewall_wan1_ipv6 = $("#validationCustom10"),
            subnet_mask_wan1_ipv6 = $("#validationCustom11"),
            gateway_wan1_ipv6 = $("#validationCustom12"),
            firewall_ipv4_lan_address = $("#validationCustom13"),
            firewall_ipv6_lan_address = $("#validationCustom14"),
            current_version = $("#validationCustom15"),
            version_backup = $("#validationCustom16"),
            backup_date = $("#datepicker-2"),
            extra = $("#validationCustom17"),
            firewall_wan2_ipv4 = $("#validationCustom18"),
            subnet_mask_wan2_ipv4 = $("#validationCustom19"),
            gateway_wan2_ipv4 = $("#validationCustom20"),
            firewall_wan2_ipv6 = $("#validationCustom21"),
            subnet_mask_wan2_ipv6 = $("#validationCustom22"),
            gateway_wan2_ipv6 = $("#validationCustom23")

        select_customer_location.val(installedFirewallData["data"][0]["customer_locations"])
        customer_lan_on_site.val(installedFirewallData["data"][0]["customer_lan_on_site"])
        brand.val(installedFirewallData["data"][0]["brand"])
        equipment.val(installedFirewallData["data"][0]["equipment"])
        installed_date.val(new Date(installedFirewallData["data"][0]["installed_date"]).toISOString().split('T')[0])
        firewall_wan1_ipv4.val(installedFirewallData["data"][0]["firewall_wan1_ipv4"])
        ip_range.val(installedFirewallData["data"][0]["ip_range"])
        subnet_mask_wan1_ipv4.val(installedFirewallData["data"][0]["sub_net_mask_wan1_ipv4"])
        gateway_wan1_ipv4.val(installedFirewallData["data"][0]["gateway_wan1_ipv4"])
        firewall_wan1_ipv6.val(installedFirewallData["data"][0]["firewall_wan1_ipv6"])
        subnet_mask_wan1_ipv6.val(installedFirewallData["data"][0]["sub_net_mask_wan1_ipv6"])
        gateway_wan1_ipv6.val(installedFirewallData["data"][0]["gateway_wan1_ipv6"])
        firewall_ipv4_lan_address.val(installedFirewallData["data"][0]["firewall_ipv4_lan_address"])
        firewall_ipv6_lan_address.val(installedFirewallData["data"][0]["firewall_ipv6_lan_address"])
        current_version.val(installedFirewallData["data"][0]["current_version"])
        version_backup.val(installedFirewallData["data"][0]["version_backup"])
        backup_date.val(new Date(installedFirewallData["data"][0]["backup_date"]).toISOString().split('T')[0])
        extra.val(installedFirewallData["data"][0]["extra"])
        firewall_wan2_ipv4.val(installedFirewallData["data"][0]["firewall_wan2_ipv4"])
        subnet_mask_wan2_ipv4.val(installedFirewallData["data"][0]["sub_net_mask_wan2_ipv4"])
        gateway_wan2_ipv4.val(installedFirewallData["data"][0]["gateway_wan2_ipv4"])
        firewall_wan2_ipv6.val(installedFirewallData["data"][0]["firewall_wan2_ipv6"])
        subnet_mask_wan2_ipv6.val(installedFirewallData["data"][0]["sub_net_mask_wan2_ipv6"])
        gateway_wan2_ipv6.val(installedFirewallData["data"][0]["gateway_wan2_ipv6"])

        let update_button = $("#update-btn")
        update_button.click(function(){
            let installed_firewalls = {
                customer_lan_on_site: customer_lan_on_site.val(),
                brand: brand.val(),
                equipment : equipment.val(),
                installed_date : new Date(installed_date.val()),
                firewall_ipv4_lan_address: firewall_ipv4_lan_address.val(),
                firewall_ipv6_lan_address:firewall_ipv6_lan_address.val(),
                current_version:current_version.val(),
                version_backup:version_backup.val(),
                backup_date: new Date(backup_date.val()),
                extra:extra.val(),
                internet_provider1: {
                    wan_config_ipv4:{
                        firewall: firewall_wan1_ipv4.val(),
                        sub_net_mask: subnet_mask_wan1_ipv4.val(),
                        gateway: gateway_wan1_ipv4.val(),
                        ip_version:"4",
                        ip_range:ip_range.val()
                    }
                },
                internet_provider2: {},
                updated_by_user_id: parseInt($("#user_id").val())
            }
            let wan_config2_ipv4 = {
                    firewall: firewall_wan2_ipv4.val(),
                    sub_net_mask: subnet_mask_wan2_ipv4.val(),
                    gateway: gateway_wan2_ipv4.val(),
                    ip_version:"4",
                },
                wan_config2_ipv6 = {
                    firewall: firewall_wan2_ipv6.val(),
                    sub_net_mask: subnet_mask_wan2_ipv6.val(),
                    gateway: gateway_wan2_ipv6.val(),
                    ip_version:"6",
                },wan_config1_ipv6={
                    firewall: firewall_wan1_ipv6.val(),
                    sub_net_mask: subnet_mask_wan1_ipv6.val(),
                    gateway: gateway_wan1_ipv6.val(),
                    ip_version:"6",
                }
            if (checkWanConfigEmpty(wan_config1_ipv6) === true){
                installed_firewalls["internet_provider1"]["wan_config_ipv6"] = wan_config1_ipv6
            }
            if (checkWanConfigEmpty(wan_config2_ipv4) === true ){
                installed_firewalls["internet_provider2"]["wan_config_ipv4"] = wan_config2_ipv4
            }
            if (checkWanConfigEmpty(wan_config2_ipv6) === true){
                installed_firewalls["internet_provider2"]["wan_config_ipv6"] = wan_config2_ipv6
            }
            if (Object.keys(installed_firewalls["internet_provider2"]).length === 0) {
                delete installed_firewalls["internet_provider2"]
            }
            patchModel(
                JSON.stringify(installed_firewalls),
                "http://localhost:8080/api/v1/installed-firewalls?installed_firewall_id="+dataId,
                `http://localhost:8080/api/v1/installed-firewalls?select_column=installed_firewalls.installed_firewall_id&append_select=true&page=`
            )
        });

        $('#modalInstalledFirewall').on('hide.bs.modal', function() {
            $("#update-btn").unbind()
        });
    });

    function checkWanConfigEmpty(object){
        return object["firewall"] !== "" && object["sub_net_mask"] !== "" && object["gateway"] !== "";
    }
});

