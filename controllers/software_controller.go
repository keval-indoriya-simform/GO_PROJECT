package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// Create Software godoc
//
//	@Summary		Create Software
//	@Description	Add    Software
//	@Tags			Software
//	@Accept			json
//	@Produce		json
//	@Param			Software	body		models.CreateSoftware	true	"Software Data"
//	@Success		201			{object}	models.Software
//	@Router			/softwares [post]
func CreateSoftwareController(context *gin.Context) {
	var softwareModel models.Software
	bindJsonError := context.ShouldBindJSON(&softwareModel)
	context.JSON(

		services.CreateServices(softwareModel, "Software", bindJsonError),
	)

}

// Update Software godoc
//
//	@Summary		Update Software
//	@Description	Update Software
//	@Tags			Software
//	@Accept			json
//	@Produce		json
//	@Param			Software	body		models.UpdateSoftware	true	"Software"
//	@Param			software_id	query		string			false	"Software ID"
//	@Success		201			{object}	models.Software
//	@Router			/softwares [patch]
func UpdateSoftwareController(context *gin.Context) {
	var softwareModel models.Software
	softwareId := context.Query("software_id")
	bindJsonError := context.ShouldBindJSON(&softwareModel)
	context.JSON(

		services.UpdateServices(softwareId, softwareModel, "Software", bindJsonError),
	)
}

// @Summary		Delete Software
// @Description	Delete Software
// @Tags			Software
// @Produce		json
// @Param			software_id	query		string	false	"Software ID"
// @Param			USER_ID		header		string	false	"USER ID"
// @Success		201			{object}	models.Software
// @Router			/softwares [delete]
func DeleteSoftwareController(context *gin.Context) {
	userId := context.Request.Header.Get("USER_ID")
	softwareId := context.Query("software_id")
	context.JSON(

		services.DeleteServices(softwareId, userId, "Software"),
	)
}

// @Summary		Retrieve  Software
// @Description	Retrieve Software
// @Tags			Software
// @Produce		json
// @Param			software_id		query	string	false	"Software ID"
// @Param			select_column	query	string	false	"Select Column"
// @Param			append_select	query	string	false	"Append Select"
// @Param			page			query	int		false	"Page"
// @Router			/softwares [get]
func RetrieveSoftwareController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "Software"
	queries["model_id"] = "software_id"
	queries["page"] = context.Query("page")
	queries["software_id"] = context.Query("software_id")
	queries["select_column"] = context.Query("select_column")
	queries["append_select"] = context.Query("append_select")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
