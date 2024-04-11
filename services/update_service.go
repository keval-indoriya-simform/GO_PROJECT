package services

import (
	"Application/models"
	"strconv"
)

func UpdateServices(modelID string, model any, modelName string, bindJSONError error) (status int, resp models.Response) {

	if modelID != "" {
		convertID, convertIDError := strconv.Atoi(modelID)
		if resp = CheckErrorService(convertIDError, models.FailParsingMsg+", didn't got model's primary key", models.InvalidIDStatus); resp == (models.Response{}) {
			if resp = CheckErrorService(bindJSONError, models.FailParsingMsg, models.FailParsingStatus); resp == (models.Response{}) {
				updatedModel, updateError := updateModelData(model, int32(convertID), modelName)
				var errorExists bool
				resp, errorExists = CheckDBErrorService(updateError, "patch")
				if !errorExists {
					resp = models.DataResponse(models.DataUpdateMsg, models.DataWriteStatus, updatedModel)
				}
			}
		}
	} else {
		resp = models.DataResponse(models.FailParsingMsg, models.InvalidIDStatus, make([]interface{}, 0))
	}
	return resp.Status, resp
}

func updateModelData(model any, modelID int32, modelName string) (modelUpdated any, updateError error) {
	switch modelName {
	case "CloudPrivateIp":
		cloudPrivateIpModel := model.(models.CloudPrivateIp)
		cloudPrivateIpModel.CloudPrivateIPId = modelID
		updateError = cloudPrivateIpModel.UpdateCloudPrivateIp()
		modelUpdated = cloudPrivateIpModel
	case "CloudPublicIp":
		cloudPublicIpModel := model.(models.CloudPublicIp)
		cloudPublicIpModel.CloudPublicIpID = modelID
		updateError = cloudPublicIpModel.UpdateCloudPublicIp()
		modelUpdated = cloudPublicIpModel
	case "CustomerLan":
		customerLanModel := model.(models.CustomerLan)
		customerLanModel.CustomerLanID = modelID
		updateError = customerLanModel.UpdateCustomerLan()
		modelUpdated = customerLanModel
	case "CustomerLocation":
		customerLocationModel := model.(models.CustomerLocation)
		customerLocationModel.CustomerLocationID = modelID
		updateError = customerLocationModel.UpdateCustomerLocation()
		modelUpdated = customerLocationModel
	case "Customer":
		customerModel := model.(models.Customer)
		customerModel.CustomerID = modelID
		updateError = customerModel.UpdateCustomer()
		modelUpdated = customerModel
	case "EmailDomain":
		emailDomainModel := model.(models.EmailDomain)
		emailDomainModel.EmailDomainID = modelID
		updateError = emailDomainModel.UpdateEmailDomain()
		modelUpdated = emailDomainModel
	case "InstalledFirewall":
		installedFirewallModel := model.(models.InstalledFirewall)
		installedFirewallModel.InstalledFirewallID = modelID
		updateError = installedFirewallModel.UpdateInstalledFirewall()
		modelUpdated = installedFirewallModel
	case "Note":
		noteModel := model.(models.Note)
		noteModel.NoteID = modelID
		updateError = noteModel.UpdateNote()
		modelUpdated = noteModel
	case "Server":
		serverModel := model.(models.Server)
		serverModel.ServerID = modelID
		updateError = serverModel.UpdateServer()
		modelUpdated = serverModel
	case "Software":
		softwareModel := model.(models.Software)
		softwareModel.SoftwareId = modelID
		updateError = softwareModel.UpdateSoftware()
		modelUpdated = softwareModel
	case "User":
		userModel := model.(models.User)
		userModel.UserID = modelID
		if userModel.Password != "" {
			userModel.Password, updateError = HashPasswordGenerateService(userModel.Password)
			if updateError != nil {
				return
			}
		}
		updateError = userModel.UpdateUser()
		modelUpdated = userModel
	case "UserRole":
		userRoleModel := model.(models.UserRole)
		userRoleModel.UserRoleID = modelID
		updateError = userRoleModel.UpdateUserRole()
		modelUpdated = userRoleModel
	default:
		updateError = models.NoModelFoundError
		modelUpdated = nil
	}
	return
}
