package models

type WanConfig struct {
	WanConfigID int32  `gorm:"primaryKey;autoIncrement;" json:"wan_config_id,omitempty"`
	Firewall    string `gorm:"type:varchar" json:"firewall,omitempty"`
	SubNetMask  string `gorm:"type:varchar" json:"sub_net_mask,omitempty"`
	Gateway     string `gorm:"type:varchar" json:"gateway,omitempty"`
	IpVersion   string `gorm:"type:varchar" json:"ip_version,omitempty"`
	IpRange     string `gorm:"type:varchar" json:"ip_range,omitempty"`
	Comment     string `gorm:"type:varchar" json:"comment,omitempty"`
}
