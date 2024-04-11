package models

import (
	"Application/initializers"
	"errors"
	"gorm.io/gorm"
	"math"
)

var (
	NotFoundError              = errors.New("no data found")
	InvalidUserPermissionError = errors.New("invalid permission")
	NoModelFoundError          = errors.New("model not found")
	IsPrimaryKeyError          = errors.New("it is primary location")
)

const (
	HealthCheckPassStatus      = 200
	HealthCheckFailStatus      = 500
	DataWriteStatus            = 201
	DataWriteFailStatus        = 400
	FailParsingStatus          = 406
	DataNotFoundStatus         = 204
	PermissionDeniedStatus     = 401
	DataFoundStatus            = 200
	InvalidIDStatus            = 428
	InvalidUserIDStatus        = 411
	InvalidUSerIDModelIDStatus = 421
	DBErrorStatus              = 500
)

const (
	HealthCheckPassMsg   = "Health Check Successful"
	HealthCheckFailMsg   = "Health Check Fail"
	DataWriteMsg         = "Record has been successfully inserted!"
	DataWriteFailMsg     = "Record insertion failed!"
	FailParsingMsg       = "Something Went Wrong Parsing Data"
	DataNotFoundMsg      = "Records Not Found"
	DataIsPrimaryMsg     = "Records Found is Primary"
	DataUpdateFailMsg    = "Record update failed!"
	DataUpdateMsg        = "Record has been successfully updated!"
	DataDeleteFailMsg    = "Record delete failed!"
	DataDeleteMsg        = "Record has been successfully deleted!"
	PermissionDeniedMsg  = "permission Denied"
	TokenNotGeneratedMsg = "Unable to generate token"
	SessionNotCreatedMsg = "Unable to create session"
	TokenGeneratedMsg    = "Token Generated Successfully"
	DataFoundMsg         = "Records have been successfully retrieved!"
	DBErrorMsg           = "Something Went Wrong Forming DB Query"
)

var (
	PageLimit int
	OrderBy   string
)

type Response struct {
	Limit      int         `json:"limit,omitempty"`
	Page       int         `json:"page,omitempty"`
	Sort       string      `json:"sort,omitempty"`
	TotalRows  int64       `json:"total_rows,omitempty"`
	TotalPages int         `json:"total_pages,omitempty"`
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func DataResponse(msg string, status int, data interface{}, resp ...Response) (res Response) {
	if len(resp) == 1 {
		res = resp[0]
	}
	res.Message = msg
	res.Status = status
	res.Data = data
	return
}

func (dataResponseModel *Response) GetOffset() int {
	return (dataResponseModel.GetPage() - 1) * dataResponseModel.GetLimit()
}

func (dataResponseModel *Response) SetTotalPages() {
	dataResponseModel.TotalPages = int(math.Ceil(float64(dataResponseModel.TotalRows) / float64(dataResponseModel.GetLimit())))
}

func (dataResponseModel *Response) GetCount(modelName string) int64 {
	var (
		value       any
		count       int64
		whereClause = "deleted_at IS NULL"
	)
	switch modelName {
	case "CloudOrOnsite":
		whereClause = ""
		value = CloudOrOnsite{}
	case "CloudPrivateIp":
		value = CloudPrivateIp{}
	case "CloudPublicIp":
		value = CloudPublicIp{}
	case "CustomerLan":
		value = CustomerLan{}
	case "CustomerLocation":
		value = CustomerLocation{}
	case "Customer":
		value = Customer{}
	case "EmailAccountType":
		whereClause = ""
		value = EmailAccountType{}
	case "EmailDomain":
		value = EmailDomain{}
	case "InstalledFirewall":
		value = InstalledFirewall{}
	case "Note":
		value = Note{}
	case "Server":
		value = Server{}
	case "Software":
		value = Software{}
	case "User":
		whereClause = ""
		value = User{}
	case "UserRole":
		value = UserRole{}
	case "DashboardCount":
		initializers.DB().Raw("SELECT count(*) FROM public.get_counts()").Count(&count)
		return count

	default:
		value = nil
	}
	initializers.DB().Model(value).Where(whereClause).Count(&count)
	return count
}

func (dataResponseModel *Response) SetFullLimit(limit int) {
	dataResponseModel.Limit = limit
}

func (dataResponseModel *Response) GetLimit() int {
	if dataResponseModel.Limit == 0 {
		dataResponseModel.Limit = PageLimit
	}
	return dataResponseModel.Limit
}

func (dataResponseModel *Response) GetPage() int {
	if dataResponseModel.Page <= 0 {
		dataResponseModel.Page = 1
	}
	return dataResponseModel.Page
}

func (dataResponseModel *Response) GetSort() string {
	if dataResponseModel.Sort == "" {
		dataResponseModel.Sort = "Id " + OrderBy
	}
	return dataResponseModel.Sort
}

func SetPagination(dataResponseModel *Response, limit int, sort string, totalRows int64) {
	dataResponseModel.Limit = limit
	dataResponseModel.Sort = sort
	dataResponseModel.TotalRows = totalRows
	dataResponseModel.SetTotalPages()
}

func Paginate(dataResponseModel *Response) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(dataResponseModel.GetOffset()).Limit(dataResponseModel.GetLimit()).Order(dataResponseModel.GetSort())
	}
}
