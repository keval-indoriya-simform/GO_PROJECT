package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SwaggerController(context *gin.Context) {

	context.Redirect(http.StatusFound, "/docs/index.html")

}
