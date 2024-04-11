package models

import (
	"Application/initializers"
	"gorm.io/gorm"
)

type EmailAccount struct {
	EmailAccountID     int32             `gorm:"primaryKey" json:"-"`
	EmailHosting       string            `gorm:"type:varchar" json:"email_hosting,omitempty"`
	LinkForEmailAdmin  string            `gorm:"type:varchar" json:"link_for_email_admin,omitempty"`
	AccountNumber      string            `gorm:"type:varchar" json:"account_number,omitempty"`
	EmailAccountTypeID int32             `gorm:"not null;"  json:"email_account_type_id,omitempty"`
	UserName           string            `gorm:"type:varchar" json:"user_name,omitempty"`
	Password           string            `gorm:"type:varchar" json:"password,omitempty"`
	Pin                string            `gorm:"type:varchar" json:"pin,omitempty"`
	EmailAccountType   *EmailAccountType `gorm:"references:EmailAccountTypeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"email_account_type,omitempty"`
}

func (emailAccountModel *EmailAccount) CreateEmailAccount() error {
	var createQuery *gorm.DB

	if emailAccountModel.EmailAccountTypeID != 0 {
		createQuery = initializers.DB().Create(&emailAccountModel)
	} else {
		createQuery = initializers.DB().Omit("EmailAccountTypeID").Create(&emailAccountModel)
	}
	if emailAccountModel.EmailAccountID > 0 {
		return nil
	}
	return createQuery.Error
}
