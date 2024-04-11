package services

import (
	"Application/models"
)

func CreateServices(model any, modelName string, bindJSONError error) (status int, resp models.Response) {
	var errorExists bool
	if resp = CheckErrorService(bindJSONError, models.FailParsingMsg, models.FailParsingStatus); resp == (models.Response{}) {
		createdModel, createError := createModelData(model, modelName)
		resp, errorExists = CheckDBErrorService(createError, "create")
		if !errorExists {
			resp = models.DataResponse(models.DataWriteMsg, models.DataWriteStatus, createdModel)
		}
	}
	return resp.Status, resp
}

func createModelData(model any, modelName string) (modelCreated any, createError error) {
	switch modelName {
	case "CloudPrivateIp":
		cloudPrivateIpModel := model.(models.CloudPrivateIp)
		createError = cloudPrivateIpModel.CreateCloudPrivateIp()
		modelCreated = cloudPrivateIpModel
	case "CloudPublicIp":
		cloudPublicIpModel := model.(models.CloudPublicIp)
		createError = cloudPublicIpModel.CreateCloudPublicIp()
		modelCreated = cloudPublicIpModel
	case "CustomerLan":
		customerLanModel := model.(models.CustomerLan)
		createError = customerLanModel.CreateCustomerLan()
		modelCreated = customerLanModel
	case "CustomerLocation":
		customerLocationModel := model.(models.CustomerLocation)
		createError = customerLocationModel.CreateCustomerLocation()
		modelCreated = customerLocationModel
	case "Customer":
		customerModel := model.(models.Customer)
		createError = customerModel.CreateCustomer()
		modelCreated = customerModel
	case "EmailDomain":
		emailDomainModel := model.(models.EmailDomain)
		createError = emailDomainModel.CreateEmailDomain()
		modelCreated = emailDomainModel
	case "InstalledFirewall":
		installedFirewallModel := model.(models.InstalledFirewall)
		createError = installedFirewallModel.CreateInstalledFirewall()
		modelCreated = installedFirewallModel
	case "Note":
		noteModel := model.(models.Note)
		createError = noteModel.CreateNote()
		modelCreated = noteModel
	case "Server":
		serverModel := model.(models.Server)
		createError = serverModel.CreateServer()
		modelCreated = serverModel
	case "Software":
		softwareModel := model.(models.Software)
		createError = softwareModel.CreateSoftware()
		modelCreated = softwareModel
	case "User":
		userModel := model.(models.User)
		userModel.Password, createError = HashPasswordGenerateService(userModel.Password)
		if createError != nil {
			return
		}
		createError = userModel.CreateUser()
		modelCreated = userModel
	case "UserRole":
		userRoleModel := model.(models.UserRole)
		createError = userRoleModel.CreateUserRole()
		modelCreated = userRoleModel
	default:
		createError = models.NoModelFoundError
		modelCreated = nil
	}
	return
}
