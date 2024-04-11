package controllers

import (
	"Application/services"
	"github.com/gin-gonic/gin"
)

func RetrieveDashboardCount(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "DashboardCount"
	queries["model_id"] = "field_name"
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(

		services.RetrieveServices(queries),
	)

}
