package models

import (
	"Application/initializers"
	"gorm.io/gorm"
)

type InternetProvider struct {
	InternetProviderID int32      `gorm:"primaryKey;autoIncrement;" json:"-"`
	Name               string     `gorm:"type:varchar" json:"name,omitempty"`
	AccountOrPin       string     `json:"account_or_pin,omitempty"`
	Other              string     `gorm:"type:varchar" json:"other,omitempty"`
	Speed              string     `gorm:"type:varchar" json:"speed,omitempty"`
	PrimaryDns         string     `gorm:"type:varchar" json:"primary_dns,omitempty"`
	SecondaryDns       string     `gorm:"type:varchar" json:"secondary_dns,omitempty"`
	WanConfigIpv4ID    int32      `json:"wan_config_ipv4_id,omitempty"`
	WanConfigIpv6ID    int32      `json:"wan_config_ipv6_id,omitempty"`
	WanConfigIpv4      *WanConfig `gorm:"references:WanConfigID;foreignKey:WanConfigIpv4ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"wan_config_ipv4,omitempty"`
	WanConfigIpv6      *WanConfig `gorm:"references:WanConfigID;foreignKey:WanConfigIpv6ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"wan_config_ipv6,omitempty"`
}

func (internetProviderModel *InternetProvider) CheckInternetProviderID(internetProviderID int32) bool {
	findQuery := initializers.DB().
		Where("internet_provider_id = ?", internetProviderID).
		Select("internet_provider_id, wan_config_ipv4_id, wan_config_ipv6_id").
		First(&internetProviderModel)
	if findQuery.RowsAffected > 0 {
		if internetProviderModel.WanConfigIpv4 != nil {
			internetProviderModel.WanConfigIpv4.WanConfigID = internetProviderModel.WanConfigIpv4ID
		} else {
			internetProviderModel.WanConfigIpv4 = &WanConfig{
				WanConfigID: internetProviderModel.WanConfigIpv4ID,
			}
		}
		if internetProviderModel.WanConfigIpv6 != nil {
			internetProviderModel.WanConfigIpv6.WanConfigID = internetProviderModel.WanConfigIpv6ID
		} else {
			internetProviderModel.WanConfigIpv6 = &WanConfig{
				WanConfigID: internetProviderModel.WanConfigIpv6ID,
			}
		}
		return true
	}
	return false
}

func (internetProviderModel *InternetProvider) CreateInternetProvider() error {
	var createQuery *gorm.DB
	if internetProviderModel.WanConfigIpv4 != nil && internetProviderModel.WanConfigIpv6 != nil {
		createQuery = initializers.DB().Create(&internetProviderModel)
	} else {
		if internetProviderModel.WanConfigIpv4 == nil && internetProviderModel.WanConfigIpv6 != nil {
			createQuery = initializers.DB().Omit("WanConfigIpv4ID").Create(&internetProviderModel)
		} else if internetProviderModel.WanConfigIpv6 == nil && internetProviderModel.WanConfigIpv4 != nil {
			createQuery = initializers.DB().Omit("WanConfigIpv6ID").Create(&internetProviderModel)
		} else {
			createQuery = initializers.DB().Omit("WanConfigIpv4ID,WanConfigIpv6ID").Create(&internetProviderModel)
		}
	}
	if internetProviderModel.InternetProviderID > 0 {
		return nil
	}
	return createQuery.Error
}
