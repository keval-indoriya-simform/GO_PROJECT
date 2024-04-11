package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"log"
	"time"
)

type CustomerLocation struct {
	CustomerLocationID int32      `gorm:"primaryKey;" json:"customer_location_id,omitempty"`
	CustomerID         int32      `gorm:"not null;" json:"customer_id,omitempty"`
	IsPrimary          *bool      `gorm:"not null" json:"is_primary,omitempty"`
	Name               string     `gorm:"type:varchar; not null;" json:"name,omitempty"`
	CreatedAt          *time.Time `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID    int32      `gorm:"not null;" json:"created_by_user_id,omitempty"`
	Description        string     `json:"description,omitempty"`
	UpdatedAt          *time.Time `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID    int32      `json:"updated_by_user_id,omitempty"`
	DeletedAt          *time.Time `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID    int32      `json:"deleted_by_user_id,omitempty"`
}

func (customerLocationModel *CustomerLocation) TableName() string {
	return "customer_locations"
}

func (customerLocationModel *CustomerLocation) CreateCustomerLocation() error {
	if !customerLocationModel.CheckCustomerPrimaryLocationExists() {
		*customerLocationModel.IsPrimary = true
	} else {
		*customerLocationModel.IsPrimary = false
	}
	log.Println(customerLocationModel.IsPrimary)
	createQuery := initializers.DB().Create(&customerLocationModel)
	if customerLocationModel.CustomerLocationID > 0 {
		CreateLog(
			customerLocationModel.CreatedByUserID,
			"Created",
			customerLocationModel.TableName(),
			customerLocationModel.CustomerLocationID,
		)
		return nil
	}
	return createQuery.Error
}

func (customerLocationModel *CustomerLocation) UpdateCustomerLocation() error {
	updateQuery := initializers.DB().Where("deleted_at IS NULL").
		Updates(&customerLocationModel)
	if updateQuery.Error != nil {
		return updateQuery.Error
	} else if updateQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		CreateLog(
			customerLocationModel.UpdatedByUserID,
			"Updated",
			customerLocationModel.TableName(),
			customerLocationModel.CustomerLocationID,
		)
		return nil
	}
}

func (customerLocationModel *CustomerLocation) DeleteCustomerLocation(customerLocationIDs []int32) error {
	customerLocationID := customerLocationModel.GetIDWithPrimaryLocationNotExists(customerLocationIDs)
	if len(customerLocationID) != 0 {
		deleteQuery := initializers.DB().Where("deleted_at IS NULL AND customer_location_id IN ?", customerLocationID).
			Updates(&customerLocationModel)
		if deleteQuery.RowsAffected > 0 {
			var logs []Log
			for index, _ := range customerLocationIDs {
				logs = append(logs, Log{
					UserID:    customerLocationModel.DeletedByUserID,
					LogType:   "Deleted",
					TableName: customerLocationModel.TableName(),
					DataID:    customerLocationIDs[index],
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
	return IsPrimaryKeyError
}

func (customerLocationModel *CustomerLocation) CheckCustomerPrimaryLocationExists() bool {
	var count int64
	initializers.DB().Model(CustomerLocation{}).
		Where("customer_id = ?", customerLocationModel.CustomerID).
		Count(&count)
	return count > 0
}

func (customerLocationModel *CustomerLocation) GetIDWithPrimaryLocationNotExists(customerLocationIDs []int32) (foundIDs []int32) {
	initializers.DB().Model(CustomerLocation{}).Where("customer_location_id in ? AND is_primary=false",
		customerLocationIDs).Select("customer_location_id").Find(&foundIDs)
	return foundIDs
}

func RetrieveCustomerLocationRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                                                  int64
		customerLocationIDQuery, locationQuery, customerQuery, assignedToQuery string
	)

	if queries["customer_location_id"] != "" {
		customerLocationIDQuery = " AND customer_locations.customer_location_id = " + queries["customer_location_id"]
	}

	if queries["customer_location_name"] != "" {
		locationQuery = " AND customer_locations.name ILIKE '%" + queries["customer_location_name"] + "%'"
	}

	if queries["customer_name"] != "" {
		customerQuery = " AND customers.name ILIKE '%" + queries["customer_name"] + "%'"
	}

	if queries["assigned_to"] != "" {
		assignedToQuery = " AND customer_locations.created_by_user_id = '" + queries["assigned_to"] + "'"
	}

	selectedColumn := "customer_locations.name,customers.name as customers"

	if queries["select_column"] != "" {
		if queries["append_select"] == "true" {
			selectedColumn += "," + queries["select_column"]
		} else {
			selectedColumn = queries["select_column"]
		}
	}

	retrievedRows := initializers.DB().Scopes(pagination).Model(CustomerLocation{}).
		Joins("INNER JOIN customers on customers.customer_id = customer_locations.customer_id").
		Where("customer_locations.deleted_at IS NULL AND customers.deleted_at IS NULL" + customerLocationIDQuery +
			locationQuery + customerQuery + assignedToQuery).
		Select(selectedColumn).
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
