package models

import (
	"Application/initializers"
	"time"
)

type Log struct {
	LogID     int32      `gorm:"primaryKey;" json:"log_id"`
	UserID    int32      `gorm:"not null;" json:"user_id"`
	LogType   string     `gorm:"type:varchar;not null;" json:"log_type"`
	TableName string     `gorm:"type:varchar;not null;" json:"table_name"`
	DataID    int32      `gorm:"not null;" json:"data_id"`
	CreatedAt *time.Time `gorm:"type:timestamp without time zone; not null" json:"created_at"`
}

func CreateLog(userID int32, logType, tableName string, dataID int32) {
	initializers.DB().Create(&Log{
		UserID:    userID,
		LogType:   logType,
		TableName: tableName,
		DataID:    dataID,
	})
}

func CreateBatchLogs(logs []Log) {
	initializers.DB().CreateInBatches(logs, len(logs))
}
