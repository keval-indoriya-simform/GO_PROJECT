package models

import (
	"Application/initializers"
	"time"
)

type UserRole struct {
	UserRoleID      int32      `gorm:"primaryKey;" json:"user_role_id,omitempty"`
	RoleID          int32      `gorm:"not null" json:"role_id,omitempty"`
	UserID          int32      `gorm:"not null" json:"user_id,omitempty"`
	CreatedByUserID int32      `gorm:"not null" json:"created_by_user_id,omitempty"`
	UpdatedByUserID int32      `json:"updated_by_user_id,omitempty"`
	CreatedAt       *time.Time `gorm:"type:timestamp without time zone; not null" json:"created_at,omitempty"`
	UpdatedAt       *time.Time `gorm:"type:timestamp without time zone;" json:"updated_at,omitempty"`
	Role            Role       `gorm:"references:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	User            User       `gorm:"references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (userRoleModel *UserRole) CreateUserRole() error {
	if roleID := RetrieveRoleID(userRoleModel.CreatedByUserID); (roleID == 1 || roleID == 2) && userRoleModel.RoleID != 1 {
		createQuery := initializers.DB().Create(&userRoleModel)
		if userRoleModel.UserRoleID > 0 {
			return nil
		}
		return createQuery.Error
	}
	return InvalidUserPermissionError
}

func (userRoleModel *UserRole) UpdateUserRole() error {
	if roleID := RetrieveRoleID(userRoleModel.UpdatedByUserID); (roleID == 1 || roleID == 2) && userRoleModel.RoleID != 1 {
		updateQuery := initializers.DB().
			Updates(&userRoleModel)
		if updateQuery.Error != nil {
			return updateQuery.Error
		} else if updateQuery.RowsAffected == 0 {
			return NotFoundError
		} else {
			return nil
		}
	}
	return InvalidUserPermissionError
}

func RetrieveRoleID(userID int32) (roleID int32) {
	initializers.DB().Model(UserRole{}).
		Where("user_id = ?", userID).
		Select("role_id").
		First(&roleID)
	return
}
