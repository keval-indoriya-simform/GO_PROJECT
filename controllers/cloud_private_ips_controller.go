package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateCloudPrivateIpController
// Creating a cloud private ip [post]
// Create CloudPrivateIP godoc
//
//	@Summary		Create Cloud Private IP
//	@Description	Add a Cloud Private IP
//	@Tags			Cloud Private IP
//	@Accept			json
//	@Produce		json
//	@Param			CloudPrivateIP	body				 models.CreateCloudPrivateIp		true	"CloudPrivateIp Data"
//	@Success		 200						{object} models.CloudPrivateIp
//	@Router			/cloud-private-ips [post]
func CreateCloudPrivateIpController(context *gin.Context) {
	var cloudPrivateIpModel models.CloudPrivateIp
	bindJsonError := context.ShouldBindJSON(&cloudPrivateIpModel)
	context.JSON(

		services.CreateServices(cloudPrivateIpModel, "CloudPrivateIp", bindJsonError),
	)
}

// UpdateCloudPrivateIpController
// Updating cloud private ip
// Update CloudPrivateIP godoc
//
//	@Summary		Update Cloud Private IP
//	@Description	Update a Cloud Private IP
//	@Tags			Cloud Private IP
//	@Accept			json
//	@Produce		json
//	@Param			CloudPrivateIP		body		models.UpdateCloudPrivateIp	true	"CloudPrivateIp Data"
//	@Param			cloud_private_ip_id	query		string								false	"Cloud Private IP ID"
//	@Success		200					{object}	models.CloudPrivateIp
//	@Router			/cloud-private-ips [patch]
func UpdateCloudPrivateIpController(context *gin.Context) {
	var cloudPrivateIpModel models.CloudPrivateIp
	cloudPrivateIpId := context.Query("cloud_private_ip_id")
	bindJsonError := context.ShouldBindJSON(&cloudPrivateIpModel)
	context.JSON(

		services.UpdateServices(cloudPrivateIpId, cloudPrivateIpModel, "CloudPrivateIp", bindJsonError),
	)

}

// Delete CloudPrivateIP godoc
//
//	@Summary		Delete Cloud Private IP
//	@Description	Delete a Cloud Private IP
//	@Tags			Cloud Private IP
//	@Produce		json
//	@Param			cloud_private_ip_id	query		string	false	"Cloud Private IP ID"
//	@Param			USER_ID				header		string	false	"USER ID"
//	@Success		200					{object}	models.CloudPrivateIp
//	@Router			/cloud-private-ips [delete]
func DeleteCloudPrivateIpController(context *gin.Context) {
	cloudPrivateIpID := context.Query("cloud_private_ip_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(cloudPrivateIpID, userID, "CloudPrivateIp"),
	)
}

// RetrieveCloudPrivateIpController
// Get cloud Private IP
// Get CloudPrivateIP godoc
//
//	@Summary		Retrieve  Cloud Private IP
//	@Description	Retrieve  Cloud Private IP
//	@Tags			Cloud Private IP
//	@Produce		json
//	@Param			page				query		string	false	"Page"
//	@Param			cloud_private_ip_id	query		string	false	"Cloud Private IP ID"
//	@Param			select_column		query		string	false	"Select Column"
//	@Param			append_select		query		string	false	"Append Select"
//	@Param			set_limit			query		string	false	"Set Limit"
//	@Success		200					{object}	models.CloudPrivateIp
//
//	@Router			/cloud-private-ips [get]
func RetrieveCloudPrivateIpController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "CloudPrivateIp"
	queries["model_id"] = "cloud_private_ip_id"
	queries["page"] = context.Query("page")
	queries["cloud_private_ip_id"] = context.Query("cloud_private_ip_id")
	queries["select_column"] = context.Query("select_column")
	queries["append_select"] = context.Query("append_select")
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
