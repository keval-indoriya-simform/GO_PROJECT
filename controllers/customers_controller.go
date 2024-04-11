package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateCustomerController
// Description To create customer
// Create Customer godoc
//
//	@Summary		Create Customer
//	@Description	Add    Customer
//	@Tags			Customer
//	@Accept			json
//	@Produce		json
//	@Param			Customer	body		models.CreateCustomer	true	"Customer Data"
//	@Success		200			{object}	models.Customer
//	@Router			/customers [post]
func CreateCustomerController(context *gin.Context) {
	var customerModel models.Customer
	bindJSONError := context.ShouldBindJSON(&customerModel)
	context.JSON(

		services.CreateServices(customerModel, "Customer", bindJSONError),
	)
}

// UpdateCustomerController
// Description To update customer
//
//	@Summary		Update Customer
//	@Description	Update Customer
//	@Tags			Customer
//	@Accept			json
//	@Produce		json
//	@Param			Customer	body		models.UpdateCustomer	true	"Customer Data"
//	@Param			customer_id	query		string							false	"Customer ID"
//	@Success		200			{object}	models.Customer
//	@Router			/customers [patch]
func UpdateCustomerController(context *gin.Context) {
	customerID := context.Query("customer_id")
	var customerModel models.Customer
	bindJSONError := context.ShouldBindJSON(&customerModel)
	context.JSON(
		services.UpdateServices(customerID, customerModel, "Customer", bindJSONError),
	)
}

// DeleteCustomerController
// Description To delete customer
//
//	@Summary		Delete Customer
//	@Description	Delete Customer
//	@Tags			Customer
//	@Produce		json
//	@Param			customer_id	query		string	false	"Customer ID"
//	@Param			USER_ID		header		string	false	"USER ID"
//	@Success		200			{object}	models.Customer
//	@Router			/customers [delete]
func DeleteCustomerController(context *gin.Context) {
	customerID := context.Query("customer_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(customerID, userID, "Customer"),
	)
}

// RetrieveCustomerController
// Description to view all customers
//
//	@Summary		Retrieve Customer
//	@Description	Retrieve Customer
//	@Tags			Customer
//	@Produce		json
//	@Param			customer_name	query		string	false	"Customer Name"
//	@Param			customer_id		query		string	false	"Customer ID"
//	@Param			created_by		query		string	false	"Created By"
//	@Param			page			query		int		false	"Page"
//	@Param			order_by		query		string	false	"Order By"
//	@Param			select_column	query		string	false	"Select Column"
//	@Param			set_limit		query		string	false	"Set Limit"
//
//	@Success		200				{object}	models.Customer
//	@Router			/customers [get]
func RetrieveCustomerController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "Customer"
	queries["model_id"] = "customer_id"
	queries["customer_name"] = context.Query("customer_name")
	queries["customer_id"] = context.Query("customer_id")
	queries["created_by"] = context.Query("created_by")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	queries["select_column"] = context.Query("select_column")
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(
		services.RetrieveServices(queries),
	)
}

// RetrieveCloudOrOnsiteController
// Description to view all cloud or onsite types
func RetrieveCloudOrOnsiteController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "CloudOrOnsite"
	queries["model_id"] = "cloud_or_onsite_id"
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
