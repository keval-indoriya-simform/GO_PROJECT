package controllers

import (
	"Application/models"
	"Application/services"
	"github.com/gin-gonic/gin"
)

// CreateNoteController
// Description To create note
// Create installed firewall godoc
//
//	@Summary		Create Notes
//	@Description	Add    Notes
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			Note	body		models.CreateNote	true	"Note Data"
//	@Success		201		{object}	models.Note
//	@Router			/notes [post]
func CreateNoteController(context *gin.Context) {
	var notesModel models.Note
	bindJSONError := context.ShouldBindJSON(&notesModel)
	context.JSON(

		services.CreateServices(notesModel, "Note", bindJSONError),
	)
}

// UpdateNoteController
// Description To update note
//
//	@Summary		Update Notes
//	@Description	Update Notes
//	@Tags			Notes
//	@Accept			json
//	@Produce		json
//	@Param			Note	body		models.UpdateNote	true	"Note"
//	@Param			note_id	query		string		false	"Note ID"
//	@Success		201		{object}	models.Note
//	@Router			/notes [patch]
func UpdateNoteController(context *gin.Context) {
	noteID := context.Query("note_id")
	var notesModel models.Note
	bindJSONError := context.ShouldBindJSON(&notesModel)
	context.JSON(

		services.UpdateServices(noteID, notesModel, "Note", bindJSONError),
	)
}

// DeleteNoteController
// Description To delete note
//
//	@Summary		Delete Notes
//	@Description	Delete Notes
//	@Tags			Notes
//	@Produce		json
//	@Param			note_id	query		string	false	"Note ID"
//	@Param			USER_ID	header		string	false	"USER ID"
//	@Success		201		{object}	models.Note
//	@Router			/notes [delete]
func DeleteNoteController(context *gin.Context) {
	noteID := context.Query("note_id")
	userID := context.Request.Header.Get("USER_ID")
	context.JSON(

		services.DeleteServices(noteID, userID, "Note"),
	)
}

// RetrieveNoteController
// Description to view all notes
//
//	@Summary		Retrieve  Notes
//	@Description	Retrieve  Notes
//	@Tags			Notes
//	@Produce		json
//	@Param			note_id		query		string	false	"Note ID"
//	@Param			subject		query		string	false	"Subject"
//	@Param			page		query		int		false	"Page"
//	@Param			order_by	query		string	false	"Order By"
//
//	@Success		201			{object}	models.Note
//	@Router			/notes [get]
func RetrieveNoteController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "Note"
	queries["model_id"] = "note_id"
	queries["note_id"] = context.Query("note_id")
	queries["subject"] = context.Query("subject")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")
	context.JSON(
		services.RetrieveServices(queries),
	)
}
