package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type Software struct {
	SoftwareId         int32            `gorm:"primaryKey;integer" json:"software_id,omitempty"`
	CustomerLocationID int32            `gorm:"not null;" json:"customer_location_id,omitempty"`
	Name               string           `gorm:"type:varchar" json:"name,omitempty"`
	Version            string           `gorm:"type:varchar" json:"version,omitempty"`
	LicenseKey         string           `gorm:"type:varchar;not null" json:"license_key,omitempty"`
	ServerOrVM         string           `gorm:"type:varchar" json:"server_or_vm,omitempty"`
	OtherLicenseInfo   string           `gorm:"type:varchar" json:"other_license_info,omitempty"`
	InstallDate        *time.Time       `gorm:"type:date" json:"install_date,omitempty"`
	ExpiryDate         *time.Time       `gorm:"type:date" json:"expiry_date,omitempty"`
	Notes              string           `gorm:"type:varchar" json:"notes,omitempty"`
	CreatedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID    int32            `gorm:"not null;" json:"created_by_user_id,omitempty"`
	UpdatedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID    int32            `json:"updated_by_user_id,omitempty"`
	DeletedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID    int32            `json:"deleted_by_user_id,omitempty"`
	CustomerLocation   CustomerLocation `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (softwareModel *Software) TableName() string {
	return "softwares"
}

func (softwareModel *Software) CreateSoftware() error {
	err := initializers.DB().Create(&softwareModel).Error
	if err != nil {
		return err
	} else {
		CreateLog(softwareModel.CreatedByUserID, "Created", softwareModel.TableName(), softwareModel.SoftwareId)
		return nil
	}
}

func (softwareModel *Software) UpdateSoftware() error {
	db := initializers.DB().Where("deleted_at IS NULL").Updates(&softwareModel)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return NotFoundError
	} else {
		CreateLog(softwareModel.UpdatedByUserID, "Updated", softwareModel.TableName(), softwareModel.SoftwareId)
		return nil
	}
}

func (softwareModel *Software) DeleteSoftware(softwareIDs []int32) error {
	rows := initializers.DB().Where("deleted_at IS NULL AND software_id IN ?", softwareIDs).Updates(&softwareModel)
	if rows.RowsAffected > 0 {
		var logs []Log
		for index, _ := range softwareIDs {
			logs = append(logs, Log{
				UserID:    softwareModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: softwareModel.TableName(),
				DataID:    softwareIDs[index],
			})
		}
		CreateBatchLogs(logs)
		return nil
	} else if rows.Error == nil && rows.RowsAffected == 0 {
		return NotFoundError
	} else {
		return rows.Error
	}
}

func RetrieveSoftwareRecords(queries map[string]string,
	paginate func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count           int64
		softwareIDQuery string
	)

	if queries["software_id"] != "" {
		softwareIDQuery = " AND softwares.software_id = " + queries["software_id"]
	}

	selectedColumns := "customers.customer_id, customer_locations.name as customer_locations," +
		"customers.name customer_name,softwares.*"
	if queries["select_column"] != "" {
		if queries["append_select"] == "true" {
			selectedColumns += "," + queries["select_column"]
		} else {
			selectedColumns = queries["select_column"]
		}
	}

	rows := initializers.DB().Model(&Customer{}).Scopes(paginate).
		Select(selectedColumns).
		Joins("right join customer_locations on customers.customer_id=customer_locations.customer_id").
		Joins("right join softwares on customer_locations.customer_location_id=softwares.customer_location_id").
		Where("softwares.deleted_at IS NULL" + softwareIDQuery).
		Find(&records)
	err = rows.Error
	count = rows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
