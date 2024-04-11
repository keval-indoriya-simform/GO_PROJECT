package models

type Role struct {
	RoleID   int32  `gorm:"primaryKey;" json:"role_id,omitempty"`
	RoleName string `gorm:"type:varchar;not null" json:"role_name,omitempty"`
}
