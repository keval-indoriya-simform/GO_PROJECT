package models

type Account struct {
	AccountID            int32  `gorm:"primaryKey" json:"-"`
	LinkForDomainAdmin   string `gorm:"type:varchar" json:"link_for_domain_admin,omitempty"`
	AccountNumber        string `gorm:"type:varchar" json:"account_number,omitempty"`
	Name                 string `gorm:"type:varchar" json:"name,omitempty"`
	AccountNumberAliases string `gorm:"type:varchar" json:"account_number_aliases,omitempty"`
	UserName             string `gorm:"type:varchar" json:"user_name,omitempty"`
}
