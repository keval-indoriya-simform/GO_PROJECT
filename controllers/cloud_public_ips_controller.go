package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateCloudPublicIPController
// Description To create cloud public ip
// CreateCloudPublicIPController
// Creating a cloud public ip [post]
// Create CloudPublicIP godoc
//
//	@Summary		Create Cloud Public IP
//	@Description	Add a Cloud Public IP
//	@Tags			Cloud Public IP
//	@Accept			json
//	@Produce		json
//	@Param			CloudPublicIp		body		models.CreateCloudPublicIp	true	"CloudPublicIp Data"
//	@Success		200					{object}	models.CloudPublicIp
//	@Router			/cloud-public-ips [post]
func CreateCloudPublicIPController(context *gin.Context) {
	var cloudPublicIpModel models.CloudPublicIp
	bindJSONError := context.ShouldBindJSON(&cloudPublicIpModel)
	context.JSON(

		services.CreateServices(cloudPublicIpModel, "CloudPublicIp", bindJSONError),
	)
}

// UpdateCloudPublicIPController
// Description To Update cloud public ip
// Update CloudPublicIP godoc
//
//	@Summary		Update Cloud Public IP
//	@Description	Update a Cloud Public IP
//	@Tags			Cloud Public IP
//	@Accept			json
//	@Produce		json
//	@Param			CloudPublicIp	body			models.UpdateCloudPublicIp	true	"CloudPublicIp Data"
//	@Param			cloud_public_ip_id	query		string								false	"Cloud Public IP ID"
//	@Success		200					{object}	models.CloudPublicIp
//	@Router			/cloud-public-ips [patch]
func UpdateCloudPublicIPController(context *gin.Context) {
	var cloudPublicIpModel models.CloudPublicIp
	cloudPublicIpID := context.Query("cloud_public_ip_id")
	bindJsonErr := context.ShouldBindJSON(&cloudPublicIpModel)
	context.JSON(

		services.UpdateServices(cloudPublicIpID, cloudPublicIpModel, "CloudPublicIp", bindJsonErr),
	)
}

// DeleteCloudPublicIPController
// Description To Delete cloud public ip
// Delete CloudPublicIP godoc
//
//	@Summary		Delete Cloud Public IP
//	@Description	Delete a Cloud Public IP
//	@Tags			Cloud Public IP
//	@Produce		json
//	@Param			cloud_public_ip_id	query		string	false	"Cloud Public IP ID"
//	@Param			USER_ID				header		string	false	"USER ID"
//	@Success		200					{object}	models.CloudPublicIp
//	@Router			/cloud-public-ips [delete]
func DeleteCloudPublicIPController(context *gin.Context) {
	cloudPublicIpID := context.Query("cloud_public_ip_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(cloudPublicIpID, userID, "CloudPublicIp"))
}

// RetrieveCloudPublicIPController
// Description To get all cloud public ip
// Get CloudPublicIP godoc
//
//	@Summary		Retrieve Cloud Public IP
//	@Description	Retrieve a Cloud Public IP
//	@Tags			Cloud Public IP
//	@Produce		json
//	@Param			cloud_public_ip_id		query		string	false	"Cloud Public IP ID"
//	@Param			ip_address				query		string	false	"Ip Address"
//	@Param			location				query		string	false	"Location"
//	@Param			cloud_vm_name			query		string	false	"Cloud VM Name"
//	@Param			page					query		int		false	"Page"
//	@Param			order_by				query		string	false	"Order By"
//	@Param			post_forward_ip			query       string  false   "Post Forward Ip"
//	@Success		200						{object}	models.CloudPublicIp
//	@Router			/cloud-public-ips [get]
func RetrieveCloudPublicIPController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "CloudPublicIp"
	queries["model_id"] = "cloud_public_ip_id"
	queries["cloud_public_ip_id"] = context.Query("cloud_public_ip_id")
	queries["ip_address"] = context.Query("ip_address")
	queries["post_forward_ip"] = context.Query("post_forward_ip")
	queries["location"] = context.Query("location")
	queries["cloud_vm_name"] = context.Query("cloud_vm_name")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(
		services.RetrieveServices(queries))
}
