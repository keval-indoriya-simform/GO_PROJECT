package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	ServerID              int32            `gorm:"primaryKey" json:"server_id,omitempty"`
	HostName              string           `gorm:"type:varchar" json:"host_name,omitempty"`
	CustomerLocationID    int32            `json:"customer_location_id,omitempty"`
	HardwareAsAService    *bool            `gorm:"column(hardware_as_a_service)" json:"hardware_as_a_service,omitempty"`
	OsPlatform            string           `gorm:"type:varchar" json:"os_platform,omitempty"`
	ServiceTag            string           `gorm:"type:varchar" json:"service_tag,omitempty"`
	ExpressionServiceCode int              `json:"expression_service_code,omitempty"`
	Location              string           `gorm:"type:varchar" json:"location,omitempty"`
	Warranty              string           `gorm:"type:varchar" json:"warranty,omitempty"`
	Type                  string           `gorm:"type:varchar;" json:"type,omitempty"`
	PowerConnectType      string           `gorm:"type:varchar; not null" json:"power_connect_type,omitempty"`
	PurchaseDate          *time.Time       `gorm:"type:date" json:"purchase_date,omitempty"`
	ExpirationDate        *time.Time       `gorm:"type:date" json:"expiration_date,omitempty"`
	DaysLeft              string           `gorm:"type:varchar" json:"days_left,omitempty"`
	Ownership             string           `gorm:"type:varchar" json:"ownership,omitempty"`
	OrderNumber           string           `gorm:"type:varchar" json:"order_number,omitempty"`
	Description           string           `gorm:"type:varchar" json:"description,omitempty"`
	Idrac                 string           `gorm:"type:varchar" json:"idrac,omitempty"`
	CreatedAt             *time.Time       `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID       int32            `gorm:"not null;" json:"created_by_user_id,omitempty"`
	UpdatedAt             *time.Time       `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID       int32            `json:"updated_by_user_id,omitempty"`
	DeletedAt             *time.Time       `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID       int32            `json:"deleted_by_user_id,omitempty"`
	CustomerLocation      CustomerLocation `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (serverModel *Server) TableName() string {
	return "servers"
}

func (serverModel *Server) CreateServer() error {
	createQuery := initializers.DB().Create(&serverModel)
	if serverModel.ServerID > 0 {
		CreateLog(serverModel.CreatedByUserID, "Created", serverModel.TableName(), serverModel.ServerID)
		return nil
	}
	return createQuery.Error
}

func (serverModel *Server) UpdateServer() error {
	updateQuery := initializers.DB().Where("deleted_at IS NULL").Updates(&serverModel)
	if updateQuery.Error != nil {
		return updateQuery.Error
	} else if updateQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		CreateLog(serverModel.UpdatedByUserID, "Updated", serverModel.TableName(), serverModel.ServerID)
		return nil
	}
}

func (serverModel *Server) DeleteServer(serverIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND server_id IN ?", serverIDs).Updates(&serverModel)
	if deleteQuery.RowsAffected > 0 {
		var logs []Log
		for index := range serverIDs {
			logs = append(logs, Log{
				UserID:    serverModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: serverModel.TableName(),
				DataID:    serverIDs[index],
			})
		}
		CreateBatchLogs(logs)
		return nil
	} else if deleteQuery.Error == nil && deleteQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		return deleteQuery.Error
	}
}

func RetrieveServerRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                                                              int64
		serverIdQuery, serviceTagQuery, equipmentTypeQuery, typeQuery, expirationDateQuery string
	)

	if queries["server_id"] != "" {
		serverIdQuery = " AND servers.server_id = " + queries["server_id"]
	}

	if queries["service_tag"] != "" {
		serviceTagQuery = " AND servers.service_tag ILIKE '%" + queries["service_tag"] + "%'"
	}
	if queries["equipment_type"] != "" {
		equipmentTypeQuery = " AND servers.power_connect_type ILIKE '%" + queries["equipment_type"] + "%'"
	}
	if queries["type"] != "" {
		typeQuery = " AND servers.type ILIKE '%" + queries["type"] + "%'"
	}
	if queries["expiration_date"] != "" {
		expirationDateQuery = " AND servers.expiration_date = " + queries["expiration_date"]
	}
	retrievedRows := initializers.DB().
		Model(&Server{}).Scopes(pagination).
		Select("servers.server_id, servers.host_name, servers.customer_location_id, " +
			"servers.hardware_as_a_service," +
			"servers.os_platform, servers.service_tag, servers.expression_service_code, " +
			"servers.warranty, servers.type, servers.description, servers.idrac," +
			"servers.power_connect_type, servers.purchase_date, servers.expiration_date, servers.days_left," +
			"servers.ownership, servers.order_number, servers.location, customers.name as customers").
		Where("servers.deleted_at IS NULL AND customer_locations.deleted_at IS NULL " + serverIdQuery +
			equipmentTypeQuery + serviceTagQuery + typeQuery + expirationDateQuery).
		Joins("inner join customer_locations " +
			"ON servers.customer_location_id = customer_locations.customer_location_id").
		Joins("inner join customers ON customers.customer_id= customer_locations.customer_id").
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
