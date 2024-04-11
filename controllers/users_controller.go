package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// Create Users godoc
//
//	 	@title 			Users
//		@Summary		Create User
//		@Description	Add    User
//		@Tags			User
//		@Accept			json
//		@Produce		json
//		@Param			User	body		    models.CreateUser	true	"User Data"
//		@Success		201			{object}	models.User
//		@Router			/users/register [post]
func RegisterUserController(context *gin.Context) {
	var userModel models.User
	bindJSONError := context.ShouldBindJSON(&userModel)
	context.JSON(

		services.CreateServices(userModel, "User", bindJSONError),
	)
}

func AddUserRoleController(context *gin.Context) {
	var userRoleModel models.UserRole
	bindJSONError := context.ShouldBindJSON(&userRoleModel)
	context.JSON(

		services.CreateServices(userRoleModel, "UserRole", bindJSONError),
	)
}

func UpdateUserRoleController(context *gin.Context) {
	userRoleID := context.Query("user_role_id")
	var userRoleModel models.UserRole
	bindJSONError := context.ShouldBindJSON(&userRoleModel)
	context.JSON(

		services.UpdateServices(userRoleID, userRoleModel, "UserRole", bindJSONError),
	)
}

func UpdateUserProfileController(context *gin.Context) {
	userID := context.Query("user_id")
	var userModel models.User
	bindJSONError := context.ShouldBindJSON(&userModel)
	context.JSON(

		services.UpdateServices(userID, userModel, "User", bindJSONError),
	)
}

// Get Users godoc
// @title 			Users
// @Summary		Retrieve  Users
// @Description	Retrieve Users
// @Tags			User
// @Produce		json
// @Success		200					{object}	models.User
// @Router			/users [get]
func RetrieveUsersController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "User"
	queries["model_id"] = "user_id"
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
