package controllers

import (
	"Application/models"
	"Application/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateServerController
// Description To create server
// Create server godoc
//
//	@Summary		Create Server
//	@Description	Add    Server
//	@Tags			Servers
//	@Accept			json
//	@Produce		json
//	@Param			Server	body		models.CreateServer	true	"Server Data"
//	@Success		201		{object}	models.Server
//	@Router			/servers [post]
func CreateServerController(context *gin.Context) {
	var serverModel models.Server
	bindJsonErr := context.ShouldBindJSON(&serverModel)
	fmt.Println(serverModel)
	context.JSON(

		services.CreateServices(serverModel, "Server", bindJsonErr),
	)
}

// UpdateServerController
// Description To Update server
//
//	@Summary		Update Server
//	@Description	Update Server
//	@Tags			Servers
//	@Accept			json
//	@Produce		json
//	@Param			Server		body		models.UpdateServer	true	"Server"
//	@Param			server_id	query		string			false	"Server ID"
//	@Success		201			{object}	models.Server
//	@Router			/servers [patch]
func UpdateServerController(context *gin.Context) {
	var serverModel models.Server
	serverID := context.Query("server_id")
	bindJsonErr := context.ShouldBindJSON(&serverModel)
	context.JSON(

		services.UpdateServices(serverID, serverModel, "Server", bindJsonErr),
	)
}

// DeleteServerController
// Description To Delete server
//
//	@Summary		Delete Server
//	@Description	Delete Server
//	@Tags			Servers
//	@Produce		json
//	@Param			server_id	query		string	false	"Server ID"
//	@Param			USER_ID		header		string	false	"USER ID"
//	@Success		201			{object}	models.Server
//	@Router			/servers [delete]
func DeleteServerController(context *gin.Context) {
	serverID := context.Query("server_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(serverID, userID, "Server"))
}

// RetrieveServerController
// Description to get all server list using applied filters
//
//	@Summary		Retrieve  Server
//	@Description	Retrieve  Server
//	@Tags			Servers
//	@Produce		json
//	@Param			server_id		query		string	false	"Server ID"
//	@Param			service_tag		query		string	false	"Service Tag"
//	@Param			equipment_type	query		string	false	"equipment type"
//	@Param			type			query		string	false	"type"
//	@Param			expiration_date	query		string	false	"expiration date"
//	@Param			page			query		int		false	"Page"
//	@Param			order_by		query		string	false	"Order By"
//
//	@Success		201				{object}	models.Note
//	@Router			/servers [get]
func RetrieveServerController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "Server"
	queries["model_id"] = "server_id"
	queries["server_id"] = context.Query("server_id")
	queries["service_tag"] = context.Query("service_tag")
	queries["equipment_type"] = context.Query("equipment_type")
	queries["type"] = context.Query("type")
	queries["expiration_date"] = context.Query("expiration_date")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
