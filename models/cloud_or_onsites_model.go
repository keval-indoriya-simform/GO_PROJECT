package models

import (
	"Application/initializers"
	"gorm.io/gorm"
)

type CloudOrOnsite struct {
	CloudOrOnsiteID int32  `gorm:"primaryKey" json:"cloud_or_onsite_id,omitempty"`
	Name            string `gorm:"type:varchar(30)" json:"name,omitempty"`
}

func RetrieveCloudOrOnsiteRecords(pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var count int64
	retrievedRows := initializers.DB().Model(CloudOrOnsite{}).Scopes(pagination).Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
