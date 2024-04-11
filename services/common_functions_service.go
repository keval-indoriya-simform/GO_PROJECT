package services

import (
	"Application/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"strings"
)

func CheckErrorService(err error, msg string, status int) (resp models.Response) {
	if err != nil {
		resp = models.DataResponse(msg, status, make([]interface{}, 0))
	}
	return
}

func CheckDBErrorService(err error, operation string, resp ...models.Response) (res models.Response, errorExsist bool) {
	if len(resp) == 1 {
		res = resp[0]
	}
	if err != nil {
		errorExsist = true
		switch operation {
		case "create":
			checkCreateErrorService(err, &res)
		case "patch":
			checkPatchErrorService(err, &res)
		case "delete":
			checkDeleteErrorService(err, &res)
		case "retrieve":
			checkRetrieveErrorService(err, &res)
		default:
			return
		}
	}
	return
}

func checkCreateErrorService(err error, resp *models.Response) {
	switch checkError := err.Error(); {
	case strings.Contains(checkError, models.InvalidUserPermissionError.Error()):
		*resp = models.DataResponse(models.PermissionDeniedMsg, models.PermissionDeniedStatus, make([]interface{}, 0), *resp)
	case strings.Contains(checkError, "duplicate"):
		*resp = models.DataResponse(models.DataWriteFailMsg+" Data Entered is Not Unique", models.DataWriteFailStatus, make([]interface{}, 0), *resp)
	default:
		*resp = models.DataResponse(models.DataWriteFailMsg, models.DataWriteFailStatus, make([]interface{}, 0), *resp)
	}

}

func checkDeleteErrorService(err error, resp *models.Response) {
	switch checkError := err.Error(); {
	case strings.Contains(checkError, models.IsPrimaryKeyError.Error()):
		*resp = models.DataResponse(models.DataIsPrimaryMsg, models.DataWriteFailStatus, make([]interface{}, 0), *resp)
	case strings.Contains(checkError, models.NotFoundError.Error()):
		*resp = models.DataResponse(models.DataNotFoundMsg, models.DataNotFoundStatus, make([]interface{}, 0), *resp)
	default:
		*resp = models.DataResponse(models.DataUpdateFailMsg, models.DataWriteFailStatus, make([]interface{}, 0), *resp)
	}

}

func checkPatchErrorService(err error, resp *models.Response) {
	switch checkError := err.Error(); {
	case strings.Contains(checkError, models.InvalidUserPermissionError.Error()):
		*resp = models.DataResponse(models.PermissionDeniedMsg, models.PermissionDeniedStatus, make([]interface{}, 0), *resp)
	case strings.Contains(checkError, models.NotFoundError.Error()):
		*resp = models.DataResponse(models.DataNotFoundMsg, models.DataNotFoundStatus, make([]interface{}, 0), *resp)
	case strings.Contains(checkError, "duplicate"):
		*resp = models.DataResponse(models.DataUpdateFailMsg+" Data Entered is Not Unique", models.DataWriteFailStatus, make([]interface{}, 0), *resp)
	default:
		*resp = models.DataResponse(models.DataUpdateFailMsg, models.DataWriteFailStatus, make([]interface{}, 0), *resp)
	}
	return
}

func checkRetrieveErrorService(err error, resp *models.Response) {
	switch checkError := err.Error(); {
	case strings.Contains(checkError, models.NotFoundError.Error()):
		*resp = models.DataResponse(models.DataNotFoundMsg, models.DataNotFoundStatus, make([]interface{}, 0), *resp)
	default:
		*resp = models.DataResponse(models.DBErrorMsg, models.DBErrorStatus, make([]interface{}, 0), *resp)
	}
	return
}

func CheckCustomErrorService(err error, customError, customMsg, msg string, status int) (resp models.Response) {
	if err != nil && strings.Contains(err.Error(), customError) {
		resp = models.DataResponse(msg, status, make([]interface{}, 0))
		resp.Message += " " + customMsg
	}
	return
}

func LogErrorService(fileName, functionName, errorMsg, reason string) {
	log.Printf("\nFilename: %s --> FunctionName: %s --> ErrorMessage: %s --> Reason: %s", fileName, functionName, errorMsg, reason)
}

func HashPasswordGenerateService(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetPageNumber(page string) (int, error) {
	if page != "" {
		convertPage, convertPageError := strconv.Atoi(page)
		if convertPageError != nil {
			log.Println(convertPageError)
			return 0, convertPageError
		}
		return convertPage, nil
	}
	return 0, nil
}

func GetSortBy(orderBy, modelID string) string {
	if orderBy != "" {
		return orderBy
	} else {
		return modelID
	}
}

func GetConditions(context *gin.Context) string {
	filters := map[string]string{
		"customer_name": "customers.name",
		"software_name": "softwares.name",
		"install_date":  "install_date",
		"expiry_date":   "expiry_date",
	}
	signs := map[string]string{
		"customer_name": "ILIKE",
		"software_name": "ILIKE",
		"install_date":  "=",
		"expiry_date":   "=",
	}
	queryString := "softwares.deleted_at IS NULL "

	for key, val := range filters {
		query := context.Query(key)

		if signs[key] == "ILIKE" {
			if query != "" {
				queryString += "AND " + val + " " + signs[key] + " '%" + query + "%' "

			}
		} else {
			if query != "" {
				queryString += "AND " + " " + val + signs[key] + "'" + query + "' "
			}
		}

	}
	return queryString
}

func checkIDExists(modelID, userID string) (resp models.Response) {
	if modelID == "" && userID == "" {
		resp = models.DataResponse(models.FailParsingMsg, models.InvalidUSerIDModelIDStatus, make([]interface{}, 0))
	} else if modelID == "" && userID != "" {
		resp = models.DataResponse(models.FailParsingMsg, models.InvalidIDStatus, make([]interface{}, 0))
	} else if modelID != "" && userID == "" {
		resp = models.DataResponse(models.FailParsingMsg, models.InvalidUserIDStatus, make([]interface{}, 0))
	}
	return
}
