package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type CloudPrivateIp struct {
	CloudPrivateIPId int32      `gorm:"type:integer;primaryKey" json:"cloud_private_ip_id,omitempty"`
	Ipv4Assignment   string     `gorm:"type:varchar;not null'" json:"ipv4_assignment,omitempty"`
	Ipv6Assignment   string     `gorm:"type:varchar" json:"ipv6_assignment,omitempty"`
	Description      string     `gorm:"type:text" json:"description,omitempty"`
	AssignedToID     int32      `gorm:"default:1" json:"assigned_to_id,omitempty"`
	CreatedAt        *time.Time `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID  int32      `gorm:"not null;" json:"created_by_user_id,omitempty"`
	UpdatedAt        *time.Time `gorm:"type:timestamp without time zone" json:"updated_at,omitempty" swaggerignore:"true"`
	UpdatedByUserID  int32      `json:"updated_by_user_id,omitempty" swaggerignore:"true"`
	DeletedAt        *time.Time `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty" swaggerignore:"true"`
	DeletedByUserID  int32      `json:"deleted_by_user_id,omitempty"swaggerignore:"true"`
	User             User       `gorm:"references:UserID;foreignKey:AssignedToID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (cloudPrivateIpModel *CloudPrivateIp) TableName() string {
	return "cloud_private_ips"
}

func (cloudPrivateIpModel *CloudPrivateIp) CreateCloudPrivateIp() error {
	err := initializers.DB().Create(&cloudPrivateIpModel).Error
	if err != nil {
		return err
	} else {
		CreateLog(cloudPrivateIpModel.CreatedByUserID, "Created", cloudPrivateIpModel.TableName(), cloudPrivateIpModel.CloudPrivateIPId)
		return nil
	}
}

func (cloudPrivateIpModel *CloudPrivateIp) UpdateCloudPrivateIp() error {
	db := initializers.DB().Where("deleted_at IS NULL").Updates(&cloudPrivateIpModel)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return NotFoundError
	} else {
		CreateLog(cloudPrivateIpModel.UpdatedByUserID, "Updated", cloudPrivateIpModel.TableName(), cloudPrivateIpModel.CloudPrivateIPId)
		return nil
	}
}

func (cloudPrivateIpModel *CloudPrivateIp) DeleteCloudPrivateIp(cloudPrivateIPIds []int32) error {
	rows := initializers.DB().Where("deleted_at IS NULL AND cloud_private_ip_id IN ?", cloudPrivateIPIds).Updates(&cloudPrivateIpModel)
	if rows.RowsAffected > 0 {
		var logs []Log
		for index, _ := range cloudPrivateIPIds {
			logs = append(logs, Log{
				UserID:    cloudPrivateIpModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: cloudPrivateIpModel.TableName(),
				DataID:    cloudPrivateIPIds[index],
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

func RetrieveCloudPrivateIpRecords(queries map[string]string,
	paginate func(db *gorm.DB) *gorm.DB) (records []map[string]interface{},
	err error) {
	var count int64
	var cloudPrivateIpIdQuery string

	if queries["cloud_private_ip_id"] != "" {
		cloudPrivateIpIdQuery = " AND cloud_private_ips.cloud_private_ip_id = " + queries["cloud_private_ip_id"]
	}

	selectedColumn := "customers.name,cloud_public_ips.cloud_vm_name,cloud_private_ips.ipv4_assignment," +
		"cloud_private_ips.ipv6_assignment,cloud_private_ips.cloud_private_ip_id,cloud_private_ips.description, " +
		"cloud_private_ips.assigned_to_id"

	if queries["select_column"] != "" {
		if queries["append_select"] == "true" {
			selectedColumn += "," + queries["select_column"]
		} else {
			selectedColumn = queries["select_column"]
		}
	}

	rows := initializers.DB().Model(&Customer{}).Scopes(paginate).
		Select(selectedColumn).
		Joins("right join customer_locations using (customer_id)").
		Joins("right join cloud_public_ips using (customer_location_id)").
		Joins("right join cloud_private_ips on cloud_public_ips.post_forward_ip=cloud_private_ips.ipv4_assignment").
		Where("cloud_public_ips.deleted_at is null and cloud_private_ips.deleted_at is null" + cloudPrivateIpIdQuery).
		Find(&records)
	err = rows.Error
	count = rows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
