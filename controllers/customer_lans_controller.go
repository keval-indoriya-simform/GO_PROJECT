package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateCustomerLanController
// Description To create customer lan
// Create CustomerLan godoc
//
//	@Summary		Create Customer Lan
//	@Description	Add a Customer Lan
//	@Tags			Customer Lan
//	@Accept			json
//	@Produce		json
//	@Param			CustomerLan	body		models.CreateCustomerLan{internet_provider1=models.CreateInternetProvider{wan_config_ipv4=models.CreateWanConfig,wan_config_ipv6=models.CreateWanConfig},internet_provider2=models.CreateInternetProvider{wan_config_ipv4=models.CreateWanConfig,wan_config_ipv6=models.CreateWanConfig},switches=models.CreateSwitches,wireless=models.CreateWireless}	true	"CustomerLan Data"
//	@Success		200			{object}	models.CustomerLan
//	@Router			/customer-lans [post]
func CreateCustomerLanController(context *gin.Context) {
	var customerLanModel models.CustomerLan
	bindJSONError := context.ShouldBindJSON(&customerLanModel)
	context.JSON(

		services.CreateServices(customerLanModel, "CustomerLan", bindJSONError),
	)
}

// UpdateCustomerLanController
// Description To update customer lan
//
//	@Summary		Update Customer Lan
//	@Description	Update Customer Lan
//	@Tags			Customer Lan
//	@Accept			json
//	@Produce		json
//	@Param			CustomerLan		body		models.UpdateCustomerLan{internet_provider1=models.UpdateInternetProvider{wan_config_ipv4=models.UpdateWanConfig,wan_config_ipv6=models.UpdateWanConfig},internet_provider2=models.UpdateInternetProvider{wan_config_ipv4=models.UpdateWanConfig,wan_config_ipv6=models.UpdateWanConfig},switches=models.UpdateSwitches,wireless=models.UpdateWireless}	true	"CustomerLan Data"
//	@Param			customer_lan_id	query		string				false	"Customer Lan ID"
//	@Success		200				{object}	models.CustomerLan
//	@Router			/customer-lans [patch]
func UpdateCustomerLanController(context *gin.Context) {
	customerLanID := context.Query("customer_lan_id")
	var customerLanModel models.CustomerLan
	bindJSONError := context.ShouldBindJSON(&customerLanModel)
	context.JSON(

		services.UpdateServices(customerLanID, customerLanModel, "CustomerLan", bindJSONError),
	)
}

// DeleteCustomerLanController
// Description To delete customer lan
// Delete Customer Lan godoc
//
//	@Summary		Delete Customer Lan
//	@Description	Delete Customer Lan
//	@Tags			Customer Lan
//	@Produce		json
//	@Param			customer_lan_id	query		string	false	"Cloud Public IP ID"
//	@Param			User_id			header		string	false	"USER ID"
//	@Success		200				{object}	models.CustomerLan
//	@Router			/customer-lans [delete]
func DeleteCustomerLanController(context *gin.Context) {
	customerLanID := context.Query("customer_lan_id")
	userID := context.Request.Header.Get("User_id")
	context.JSON(

		services.DeleteServices(customerLanID, userID, "CustomerLan"),
	)
}

// RetrieveCustomerLanController
// Description To filter customer lan
//
//	@Summary		Retrieve Customer Lan
//	@Description	Retrieve Customer Lan
//	@Tags			Customer Lan
//	@Produce		json
//	@Param			customer_location				query		string	false	"Customer Location"
//	@Param			customer_lan_network_on_site	query		string	false	"Customer Lan Network On Site"
//	@Param			assign_to						query		string	false	"Assign To"
//	@Param			page							query		int		false	"Page"
//	@Param			order_by						query		string	false	"Order By"
//
//	@Success		200								{object}	models.CustomerLan
//	@Router			/customer-lans [get]
func RetrieveCustomerLanController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "CustomerLan"
	queries["model_id"] = "customer_lan_id"
	queries["customer_lan_id"] = context.Query("customer_lan_id")
	queries["customer_location"] = context.Query("customer_location")
	queries["customer_lan_network_on_site"] = context.Query("customer_lan_network_on_site")
	queries["assign_to"] = context.Query("assign_to")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	queries["select_column"] = context.Query("select_column")
	queries["append_select"] = context.Query("append_select")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
