package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//	   User Login godoc
//
//		@Summary		User Login
//		@Description	User Login
//		@Tags			User Login
//		@Accept			json
//		@Produce		json
//		@Param			Login	body		models.Login	true	"Login Data"
//		@Success		201		{object}	models.Login
//		@Router			/login [post]
func UserLoginController(context *gin.Context) {
	var userModel models.User
	bindJSONError := context.ShouldBindJSON(&userModel)
	context.JSON(

		services.LoginUserService(context, userModel, bindJSONError),
	)
}

func LogoutController(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
	if sessionSaveError := session.Save(); sessionSaveError != nil {
		log.Fatal(sessionSaveError)
	}
	context.Redirect(http.StatusFound, "/network-management-solutions")
}
