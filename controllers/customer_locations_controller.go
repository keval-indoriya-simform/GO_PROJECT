package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateCustomerLocationController
// Description To create customer location
// Create CustomerLocation godoc
//
//	@Summary		Create Customer Location
//	@Description	Add    Customer Location
//	@Tags			Customer Location
//	@Accept			json
//	@Produce		json
//	@Param			CustomerLocation	body		models.CreateCustomerLocation	true	"CustomerLocation Data"
//	@Success		200					{object}	models.CustomerLocation
//	@Router			/customer-locations [post]
func CreateCustomerLocationController(context *gin.Context) {
	var customerLocationModel models.CustomerLocation
	bindJSONError := context.ShouldBindJSON(&customerLocationModel)
	context.JSON(

		services.CreateServices(customerLocationModel, "CustomerLocation", bindJSONError),
	)
}

// UpdateCustomerLocationController
// Description To update customer location
//
//	@Summary		Update Customer Location
//	@Description	Update Customer Location
//	@Tags			Customer Location
//	@Accept			json
//	@Produce		json
//	@Param			CustomerLocation		body		models.UpdateCustomerLocation	true	"CustomerLan Data"
//	@Param			customer_location_id	query		string									false	"Customer Location ID"
//	@Success		200						{object}	models.CustomerLocation
//	@Router			/customer-locations [patch]
func UpdateCustomerLocationController(context *gin.Context) {
	customerLocationID := context.Query("customer_location_id")
	var customerLocationModel models.CustomerLocation
	bindJSONError := context.ShouldBindJSON(&customerLocationModel)
	context.JSON(

		services.UpdateServices(customerLocationID, customerLocationModel, "CustomerLocation", bindJSONError),
	)
}

// DeleteCustomerLocationController
// Description To delete customer location
// Delete customer location godoc
//
//	@Summary		Delete Customer Location
//	@Description	Delete Customer Location
//	@Tags			Customer Location
//	@Produce		json
//	@Param			customer_location_id	query		string	false	"Customer Location ID"
//	@Param			USER_ID					header		string	false	"USER ID"
//	@Success		200						{object}	models.CustomerLocation
//	@Router			/customer-locations [delete]
func DeleteCustomerLocationController(context *gin.Context) {
	customerLocationID := context.Query("customer_location_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(customerLocationID, userID, "CustomerLocation"),
	)
}

// RetrieveCustomerLocationController
// Description To get customer location
//
//	@Summary		Retrieve Customer Location
//	@Description	Retrieve Customer Location
//	@Tags			Customer Location
//	@Produce		json
//	@Param			customer_location_id	query		string	false	"Customer Location ID"
//	@Param			customer_location_name	query		string	false	"Customer Location Name"
//	@Param			customer_name			query		string	false	"Customer Name"
//	@Param			assigned_to				query		string	false	"Assign To"
//	@Param			page					query		int		false	"Page"
//	@Param			order_by				query		string	false	"Order By"
//	@Param			select_column			query		string	false	"Select Column"
//	@Param			set_limit				query		string	false	"Set Limit"
//	@Param			append_select			query		string	false	"Append Select"
//
//	@Success		200						{object}	models.CustomerLocation
//	@Router			/customer-locations [get]
func RetrieveCustomerLocationController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "CustomerLocation"
	queries["model_id"] = "customer_location_id"
	queries["customer_location_id"] = context.Query("customer_location_id")
	queries["customer_location_name"] = context.Query("customer_location_name")
	queries["customer_name"] = context.Query("customer_name")
	queries["assigned_to"] = context.Query("assigned_to")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	queries["select_column"] = context.Query("select_column")
	queries["set_limit"] = context.Query("set_limit")
	queries["append_select"] = context.Query("append_select")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
