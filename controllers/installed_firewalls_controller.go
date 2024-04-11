package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateInstalledFirewallController
// Description To create installed firewall
// Create installed firewall godoc
//
//	@Summary		Create Installed Firewall
//	@Description	Add    Installed Firewall
//	@Tags			Installed Firewall
//	@Accept			json
//	@Produce		json
//	@Param			InstalledFirewall	body		models.CreateInstalledFirewall{internet_provider1=models.CreateInternetProvider{wan_config_ipv4=models.CreateWanConfig,wan_config_ipv6=models.CreateWanConfig},internet_provider2=models.CreateInternetProvider{wan_config_ipv4=models.CreateWanConfig,wan_config_ipv6=models.CreateWanConfig}}	true	"InstalledFirewall Data"
//	@Success		200					{object}	models.InstalledFirewall
//	@Router			/installed-firewalls [post]
func CreateInstalledFirewallController(context *gin.Context) {
	var installedFirewallModel models.InstalledFirewall
	bindJSONError := context.ShouldBindJSON(&installedFirewallModel)
	context.JSON(

		services.CreateServices(installedFirewallModel, "InstalledFirewall", bindJSONError),
	)
}

// UpdateInstalledFirewallController
// Description To update installed firewall
//
//	@Summary		Update Installed Firewall
//	@Description	Update Installed Firewall
//	@Tags			Installed Firewall
//	@Accept			json
//	@Produce		json
//	@Param			InstalledFirewall		body		models.UpdateInstalledFirewall{internet_provider1=models.UpdateInternetProvider{wan_config_ipv4=models.UpdateWanConfig,wan_config_ipv6=models.UpdateWanConfig},internet_provider2=models.UpdateInternetProvider{wan_config_ipv4=models.UpdateWanConfig,wan_config_ipv6=models.UpdateWanConfig}}	true	"Installed Firewall"
//	@Param			installed_firewall_id	query		string						false	"Installed Firewall ID"
//	@Success		200						{object}	models.InstalledFirewall
//	@Router			/installed-firewalls [patch]
func UpdateInstalledFirewallController(context *gin.Context) {
	installedFirewallID := context.Query("installed_firewall_id")
	var installedFirewallModel models.InstalledFirewall
	bindJSONError := context.ShouldBindJSON(&installedFirewallModel)
	context.JSON(

		services.UpdateServices(installedFirewallID, installedFirewallModel, "InstalledFirewall", bindJSONError),
	)
}

// DeleteInstalledFirewallController
// Description To delete installed firewall
//
//	@Summary		Delete Installed Firewall
//	@Description	Delete Installed Firewall
//	@Tags			Installed Firewall
//	@Produce		json
//	@Param			installed_firewall_id	query		string	false	"Installed Firewall ID"
//	@Param			USER_ID					header		string	false	"USER ID"
//	@Success		200						{object}	models.InstalledFirewall
//	@Router			/installed-firewalls [delete]
func DeleteInstalledFirewallController(context *gin.Context) {
	installedFirewallID := context.Query("installed_firewall_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(installedFirewallID, userID, "InstalledFirewall"),
	)
}

// RetrieveInstalledFirewallController
// Description To get installed firewall
//
//	@Summary		Retrieve  Installed Firewall
//	@Description	Retrieve  Installed Firewall
//	@Tags			Installed Firewall
//	@Produce		json
//	@Param			version_backup		query		string	false	"Domain"
//	@Param			firewall_wan1_ipv4	query		string	false	"Email Domain Id"
//	@Param			backup_date			query		string	false	"Customer Location"
//	@Param			page				query		int		false	"Page"
//	@Param			order_by			query		string	false	"Order By"
//	@Param			select_column		query		string	false	"Select Column"
//	@Param			append_select		query		string	false	"Append Select"
//
//	@Success		200					{object}	models.InstalledFirewall
//	@Router			/installed-firewalls [get]
func RetrieveInstalledFirewallController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "InstalledFirewall"
	queries["model_id"] = "installed_firewall_id"
	queries["installed_firewall_id"] = context.Query("installed_firewall_id")
	queries["version_backup"] = context.Query("version_backup")
	queries["firewall_wan1_ipv4"] = context.Query("firewall_wan1_ipv4")
	queries["backup_date"] = context.Query("backup_date")
	queries["customer_location_id"] = context.Query("customer_location_id")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	queries["select_column"] = context.Query("select_column")
	queries["append_select"] = context.Query("append_select")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
