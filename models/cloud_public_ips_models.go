package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type CloudPublicIp struct {
	CloudPublicIpID    int32            `gorm:"primaryKey" json:"cloud_public_ip_id,omitempty"`
	IpAddress          string           `gorm:"type:varchar;not null;" json:"ip_address,omitempty"`
	CustomerLocationID int32            `json:"customer_location_id,omitempty"`
	PostForwardIp      string           `gorm:"type:varchar" json:"post_forward_ip,omitempty"`
	CloudVmName        string           `gorm:"type:varchar" json:"cloud_vm_name,omitempty"`
	CreatedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID    int32            `gorm:"not null;" json:"created_by_user_id,omitempty"`
	UpdatedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID    int32            `json:"updated_by_user_id,omitempty"`
	DeletedAt          *time.Time       `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID    int32            `json:"deleted_by_user_id,omitempty"`
	CustomerLocation   CustomerLocation `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (cloudPublicIpModel *CloudPublicIp) TableName() string {
	return "cloud_public_ips"
}

func (cloudPublicIpModel *CloudPublicIp) CreateCloudPublicIp() error {
	createQuery := initializers.DB().Create(&cloudPublicIpModel)
	if cloudPublicIpModel.CloudPublicIpID > 0 {
		CreateLog(
			cloudPublicIpModel.CreatedByUserID,
			"Created",
			cloudPublicIpModel.TableName(),
			cloudPublicIpModel.CloudPublicIpID,
		)
		return nil
	}
	return createQuery.Error
}

func (cloudPublicIpModel *CloudPublicIp) UpdateCloudPublicIp() error {
	if cloudPublicIpModel.CheckCloudPublicIpID(cloudPublicIpModel.CloudPublicIpID) {
		updateQuery := initializers.DB().Where("deleted_at IS NULL").Updates(&cloudPublicIpModel)
		if updateQuery.RowsAffected > 0 {
			CreateLog(
				cloudPublicIpModel.UpdatedByUserID,
				"Updated",
				cloudPublicIpModel.TableName(),
				cloudPublicIpModel.CloudPublicIpID,
			)
			return nil
		}
		return updateQuery.Error
	}
	return NotFoundError
}

func (cloudPublicIpModel *CloudPublicIp) DeleteCloudPublicIp(cloudPublicIpIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND cloud_public_ip_id IN ?", cloudPublicIpIDs).Updates(&cloudPublicIpModel)
	if deleteQuery.RowsAffected > 0 {
		var logs []Log
		for index, _ := range cloudPublicIpIDs {
			logs = append(logs, Log{
				UserID:    cloudPublicIpModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: cloudPublicIpModel.TableName(),
				DataID:    cloudPublicIpIDs[index],
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

func (cloudPublicIpModel *CloudPublicIp) CheckCloudPublicIpID(cloudPublicIpId int32) bool {
	findQuery := initializers.DB().
		Select("cloud_public_ip_id").
		Where("cloud_public_ip_id = ? AND deleted_at IS NULL ", cloudPublicIpId).
		First(&cloudPublicIpModel)
	if findQuery.RowsAffected > 0 {
		return true
	}
	return false
}

func RetrieveCloudPublicIpRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                                                                         int64
		cloudPublicIpIDQuery, ipAddressQuery, customerLocationQuery, cloudVmQuery, postForwardIpQuery string
	)
	if queries["cloud_public_ip_id"] != "" {
		cloudPublicIpIDQuery = " AND cloud_public_ips.cloud_public_ip_id = " + queries["cloud_public_ip_id"]
	}

	if queries["ip_address"] != "" {
		ipAddressQuery = " AND cloud_public_ips.ip_address ILIKE '%" + queries["ip_address"] + "%'"
	}

	if queries["location"] != "" {
		customerLocationQuery = " AND customer_locations.name ILIKE '%" + queries["location"] + "%'"
	}

	if queries["cloud_vm_name"] != "" {
		cloudVmQuery = " AND cloud_public_ips.cloud_vm_name ILIKE '%" + queries["cloud_vm_name"] + "%'"
	}

	if queries["post_forward_ip"] != "" {
		postForwardIpQuery = " AND cloud_public_ips.post_forward_ip ILIKE '%" + queries["post_forward_ip"] + "%'"
	}

	retrievedRows := initializers.DB().
		Model(&CloudPublicIp{}).Scopes(pagination).
		Select("cloud_public_ips.*, customer_locations.name AS customer_location").
		Joins("inner join customer_locations USING(customer_location_id)").
		Where("cloud_public_ips.deleted_at IS NULL AND customer_locations.deleted_at IS NULL " +
			cloudPublicIpIDQuery + ipAddressQuery + cloudVmQuery + customerLocationQuery +
			postForwardIpQuery).
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
