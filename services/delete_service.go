package services

import (
	"Application/models"
	"strconv"
	"strings"
	"time"
)

func DeleteServices(modelID, userID, modelName string) (status int, resp models.Response) {
	if resp = checkIDExists(modelID, userID); resp == (models.Response{}) {
		id := strings.Split(strings.TrimSpace(modelID), ",")
		var (
			convertModelID      int
			convertModelIDError error
		)
		convertedID := make([]int32, 0)
		for index, _ := range id {
			convertModelID, convertModelIDError = strconv.Atoi(id[index])
			if convertModelIDError != nil {
				break
			}
			convertedID = append(convertedID, int32(convertModelID))
		}
		if resp = CheckErrorService(convertModelIDError, models.FailParsingMsg+", didn't got model's primary key", models.InvalidIDStatus); resp == (models.Response{}) {
			convertUserID, convertUserIDError := strconv.Atoi(userID)
			if resp = CheckErrorService(convertUserIDError, models.FailParsingMsg+", didn't got user_id", models.InvalidUserIDStatus); resp == (models.Response{}) {
				deletedModel, deleteError := deleteModelData(int32(convertUserID), modelName, convertedID)
				var errorExists bool
				resp, errorExists = CheckDBErrorService(deleteError, "delete")
				if !errorExists {
					resp = models.DataResponse(models.DataDeleteMsg, models.DataWriteStatus, deletedModel)
				}
			}
		}
	}
	return resp.Status, resp
}

func deleteModelData(userID int32, modelName string, modelID []int32) (modelUpdated any, deleteError error) {
	currentTime := time.Now()
	switch modelName {
	case "CloudPrivateIp":
		cloudPrivateIpModel := models.CloudPrivateIp{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = cloudPrivateIpModel.DeleteCloudPrivateIp(modelID)
		modelUpdated = cloudPrivateIpModel
	case "CloudPublicIp":
		cloudPublicIpModel := models.CloudPublicIp{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = cloudPublicIpModel.DeleteCloudPublicIp(modelID)
		modelUpdated = cloudPublicIpModel
	case "CustomerLan":
		customerLanModel := models.CustomerLan{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = customerLanModel.DeleteCustomerLan(modelID)
		modelUpdated = customerLanModel
	case "CustomerLocation":
		customerLocationModel := models.CustomerLocation{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = customerLocationModel.DeleteCustomerLocation(modelID)
		modelUpdated = customerLocationModel
	case "Customer":
		customerModel := models.Customer{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = customerModel.DeleteCustomer(modelID)
		modelUpdated = customerModel
	case "EmailDomain":
		emailDomainModel := models.EmailDomain{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = emailDomainModel.DeleteEmailDomain(modelID)
		modelUpdated = emailDomainModel
	case "InstalledFirewall":
		installedFirewallModel := models.InstalledFirewall{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = installedFirewallModel.DeleteInstalledFirewall(modelID)
		modelUpdated = installedFirewallModel
	case "Note":
		noteModel := models.Note{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = noteModel.DeleteNote(modelID)
		modelUpdated = noteModel
	case "Server":
		serverModel := models.Server{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = serverModel.DeleteServer(modelID)
		modelUpdated = serverModel
	case "Software":
		softwareModel := models.Software{
			DeletedAt:       &currentTime,
			DeletedByUserID: userID,
		}
		deleteError = softwareModel.DeleteSoftware(modelID)
		modelUpdated = softwareModel
	default:
		deleteError = models.NoModelFoundError
		modelUpdated = nil
	}
	return
}
