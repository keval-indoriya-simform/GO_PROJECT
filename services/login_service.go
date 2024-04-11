package services

import (
	"Application/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginUserService(context *gin.Context, userModel models.User, bindJSONError error) (status int, resp models.Response) {
	if resp = CheckErrorService(bindJSONError, models.FailParsingMsg, models.FailParsingStatus); resp == (models.Response{}) {
		if userCredentials := userModel.CheckCredentials(); len(userCredentials) != 0 {
			token := GenerateToken(userCredentials["user_id"].(int32), userCredentials["name"].(string), userCredentials["email"].(string), userCredentials["role_name"].(string))
			if len(token) != 0 {
				session := sessions.Default(context)
				session.Set("token", token)
				saveSessionError := session.Save()
				if saveSessionError != nil {
					log.Println(saveSessionError)
					resp = models.DataResponse(models.SessionNotCreatedMsg, models.FailParsingStatus, make([]interface{}, 0))
				} else {
					resp = models.DataResponse(models.DataFoundMsg, models.DataFoundStatus, make([]interface{}, 0))
				}
			} else {
				resp = models.DataResponse(models.TokenNotGeneratedMsg, models.FailParsingStatus, make([]interface{}, 0))
			}
		} else {
			resp = models.DataResponse(models.PermissionDeniedMsg, models.PermissionDeniedStatus, make([]interface{}, 0))
		}
	}
	return resp.Status, resp
}
