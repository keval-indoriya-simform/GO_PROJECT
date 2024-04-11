package controllers

import (
	"Application/services"
	"github.com/gin-gonic/gin"
)

func HealthCheckController(context *gin.Context) {
	context.JSON(

		services.PingDatabase(),
	)
}
