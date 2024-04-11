package models

import (
	"Application/initializers"
	"gorm.io/gorm"
)

func RetrieveDashboardCount(pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var count int64

	retrievedRows := initializers.DB().Scopes(pagination).
		Raw("SELECT * FROM public.get_counts()").
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
