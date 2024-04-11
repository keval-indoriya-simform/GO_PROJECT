package services

import (
	"Application/models"
	"gorm.io/gorm"
	"log"
)

func RetrieveServices(queries map[string]string) (status int, resp models.Response) {
	pageNum, getPageErr := GetPageNumber(queries["page"])
	if resp = CheckErrorService(getPageErr, models.FailParsingMsg, models.FailParsingStatus); resp == (models.Response{}) {
		setLimit, exists := queries["set_limit"]
		totalCount := resp.GetCount(queries["model_name"])
		log.Println(totalCount)
		if exists && setLimit == "false" {
			resp.SetFullLimit(int(totalCount))
		}
		resp.Page = pageNum
		resp.Sort = GetSortBy(queries["order_by"], queries["model_id"])
		retrievedModelData, retrieveCount, retrievedError := retrieveModelData(queries, models.Paginate(&resp))
		var errorExists bool
		resp, errorExists = CheckDBErrorService(retrievedError, "retrieve", resp)
		if !errorExists {
			if retrieveCount == 20 {
				models.SetPagination(&resp, resp.GetLimit(), resp.GetSort(), totalCount)
			} else {
				models.SetPagination(&resp, resp.GetLimit(), resp.GetSort(), retrieveCount)
			}
			resp = models.DataResponse(models.DataFoundMsg, models.DataFoundStatus, retrievedModelData, resp)

		}
	}
	return resp.Status, resp
}

func retrieveModelData(queries map[string]string, paginate func(db *gorm.DB) *gorm.DB) (modelRetrieved any, count int64,
	modelRetrievedError error) {
	retrieveRecords := make([]map[string]interface{}, 0)
	switch queries["model_name"] {
	case "CloudOrOnsite":
		retrieveRecords, modelRetrievedError = models.RetrieveCloudOrOnsiteRecords(paginate)
	case "CloudPrivateIp":
		retrieveRecords, modelRetrievedError = models.RetrieveCloudPrivateIpRecords(queries, paginate)
	case "CloudPublicIp":
		retrieveRecords, modelRetrievedError = models.RetrieveCloudPublicIpRecords(queries, paginate)
	case "CustomerLan":
		retrieveRecords, modelRetrievedError = models.RetrieveCustomerLanRecords(queries, paginate)
	case "CustomerLocation":
		retrieveRecords, modelRetrievedError = models.RetrieveCustomerLocationRecords(queries, paginate)
	case "Customer":
		retrieveRecords, modelRetrievedError = models.RetrieveCustomerRecords(queries, paginate)
	case "EmailAccountType":
		retrieveRecords, modelRetrievedError = models.RetrieveEmailAccountTypeRecords(paginate)
	case "EmailDomain":
		retrieveRecords, modelRetrievedError = models.RetrieveEmailDomainRecords(queries, paginate)
	case "InstalledFirewall":
		retrieveRecords, modelRetrievedError = models.RetrieveInstalledFirewallRecords(queries, paginate)
	case "Note":
		retrieveRecords, modelRetrievedError = models.RetrieveNoteRecords(queries, paginate)
	case "Server":
		retrieveRecords, modelRetrievedError = models.RetrieveServerRecords(queries, paginate)
	case "Software":
		retrieveRecords, modelRetrievedError = models.RetrieveSoftwareRecords(queries, paginate)
	case "User":
		retrieveRecords, modelRetrievedError = models.RetrieveUsers(paginate)
	case "DashboardCount":
		retrieveRecords, modelRetrievedError = models.RetrieveDashboardCount(paginate)
	default:
		retrieveRecords = nil
		modelRetrievedError = models.NoModelFoundError
	}
	count = int64(len(retrieveRecords))
	modelRetrieved = retrieveRecords

	return
}
