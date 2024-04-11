package models

import (
	"Application/initializers"
	"gorm.io/gorm"
)

type EmailAccountType struct {
	EmailAccountTypeID int32  `gorm:"primaryKey" json:"email_account_type_id,omitempty"`
	EmailAccountType   string `gorm:"type:varchar" json:"email_account_type,omitempty"`
}

func RetrieveEmailAccountTypeRecords(pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{},
	err error) {
	var count int64
	retrievedRows := initializers.DB().Model(EmailAccountType{}).Scopes(pagination).Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
