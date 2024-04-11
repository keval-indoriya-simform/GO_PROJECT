package models

import (
	"Application/initializers"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	CustomerID         int32              `gorm:"primaryKey" json:"customer_id,omitempty"`
	Name               string             `gorm:"type:varchar(30);not null" json:"name,omitempty"`
	IsActive           *bool              `gorm:"" json:"is_active,omitempty"`
	CloudOrOnsiteID    int32              `gorm:"default:1" json:"cloud_or_onsite_id,omitempty"`
	Voip               *bool              `gorm:"" json:"voip,omitempty"`
	Internet           *bool              `gorm:"" json:"internet,omitempty"`
	Firewall           *bool              `gorm:"" json:"firewall,omitempty"`
	BackupSoftware     string             `gorm:"type:varchar" json:"backup_software,omitempty"`
	HardwareAsAService *bool              `gorm:"" json:"hardware_as_a_service,omitempty"`
	Description        string             `gorm:"type:varchar" json:"description,omitempty"`
	CreatedAt          *time.Time         `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID    int32              `gorm:"not null;" json:"created_by_user_id,omitempty"`
	UpdatedAt          *time.Time         `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID    int32              `gorm:"" json:"updated_by_user_id,omitempty"`
	DeletedAt          *time.Time         `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID    int32              `gorm:"" json:"deleted_by_user_id,omitempty"`
	CloudOrOnsite      CloudOrOnsite      `gorm:"references:CloudOrOnsiteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CustomerLocations  []CustomerLocation `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (customerModel *Customer) TableName() string {
	return "customers"
}

func (customerModel *Customer) CreateCustomer() error {
	createQuery := initializers.DB().Create(&customerModel)
	if customerModel.CustomerID > 0 {
		CreateLog(
			customerModel.CreatedByUserID,
			"Created",
			customerModel.TableName(),
			customerModel.CustomerID,
		)
		return nil
	}
	return createQuery.Error
}

func (customerModel *Customer) UpdateCustomer() error {
	updateQuery := initializers.DB().Where("deleted_at IS NULL").
		Updates(&customerModel)
	if updateQuery.Error != nil {
		return updateQuery.Error
	} else if updateQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		CreateLog(
			customerModel.UpdatedByUserID,
			"Updated",
			customerModel.TableName(),
			customerModel.CustomerID,
		)
		fmt.Println(customerModel.HardwareAsAService)
		return nil
	}
}

func (customerModel *Customer) DeleteCustomer(customerIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND customer_id IN ?", customerIDs).
		Updates(&customerModel)
	if deleteQuery.RowsAffected > 0 {
		var logs []Log
		for index, _ := range customerIDs {
			logs = append(logs, Log{
				UserID:    customerModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: customerModel.TableName(),
				DataID:    customerIDs[index],
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

func RetrieveCustomerRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                      int64
		nameQuery, createdByQuery, customerIDQuery string
	)

	if queries["customer_id"] != "" {
		customerIDQuery = "AND customers.customer_id = " + queries["customer_id"]
	}

	if queries["customer_name"] != "" {
		nameQuery = "AND customers.name ILIKE '%" + queries["customer_name"] + "%'"
	}

	if queries["created_by"] != "" {
		createdByQuery = "AND customers.created_by_user_id = " + queries["created_by"]
	}

	selectedColumn := "customers.customer_id,customers.name, customers.is_active, customers.voip, customers.internet, customers.firewall, " +
		"customers.backup_software, customers.hardware_as_a_service, customers.description," +
		" c.name as cloud_or_onsite, customers.cloud_or_onsite_id"

	if queries["select_column"] != "" {
		selectedColumn = queries["select_column"]
	}

	retrievedRows := initializers.DB().Scopes(pagination).Model(&Customer{}).
		Joins("JOIN cloud_or_onsites as c ON c.cloud_or_onsite_id=customers.cloud_or_onsite_id").
		Where("deleted_at IS NULL " + customerIDQuery + nameQuery + createdByQuery).
		Select(selectedColumn).
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
