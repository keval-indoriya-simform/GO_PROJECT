package models

import (
	"Application/initializers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID     int32      `gorm:"primaryKey;" json:"user_id,omitempty"`
	Name       string     `gorm:"type:varchar;"  json:"name,omitempty"`
	Email      string     `gorm:"type:varchar;not null;unique" json:"email,omitempty"`
	Contact    int        `json:"contact,omitempty"`
	Department string     `gorm:"type:varchar;" json:"department,omitempty"`
	Username   string     `gorm:"type:varchar;not null;unique" json:"username,omitempty"`
	Password   string     `gorm:"type:varchar;not null" json:"password,omitempty"`
	CreatedAt  *time.Time `gorm:"type:timestamp without time zone; not null" json:"created_at,omitempty"`
	UpdatedAt  *time.Time `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	IsActive   *bool      `gorm:"not null;" json:"is_active,omitempty"`
}

func (userModel *User) CreateUser() error {
	createQuery := initializers.DB().Create(&userModel)
	if userModel.UserID > 0 {
		return nil
	}
	return createQuery.Error
}

func (userModel *User) UpdateUser() error {
	updateQuery := initializers.DB().
		Updates(&userModel)
	if updateQuery.Error != nil {
		return updateQuery.Error
	} else if updateQuery.RowsAffected == 0 {
		return NotFoundError
	} else {
		return nil
	}
}

func (userModel *User) CheckCredentials() map[string]interface{} {
	var userCredentials map[string]interface{}
	findQuery := initializers.DB().Model(User{}).Where("users.username = ?", userModel.Username).
		Joins("inner join user_roles on user_roles.user_id = users.user_id").
		Joins("inner join roles on roles.role_id = user_roles.role_id").
		Select("users.*,roles.role_name").First(&userCredentials)
	if findQuery.RowsAffected > 0 {
		if userModel.CheckPasswordHash(userCredentials["password"].(string)) {
			return userCredentials
		}
		return nil
	}
	return nil
}

func (userModel *User) CheckPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(userModel.Password))
	return err == nil
}

func RetrieveUsers(pagination func(db *gorm.DB) *gorm.DB) (Users []map[string]interface{}, err error) {
	var count int64
	findQuery := initializers.DB().Model(User{}).Scopes(pagination).Find(&Users)
	err = findQuery.Error
	count = findQuery.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
