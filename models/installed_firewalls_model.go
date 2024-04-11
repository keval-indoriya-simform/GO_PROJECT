package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type InstalledFirewall struct {
	InstalledFirewallID    int32             `gorm:"primaryKey;" json:"installed_firewall_id,omitempty"`
	CustomerLocationID     int32             `json:"customer_location_id,omitempty"`
	CustomerLanOnSite      string            `gorm:"type:varchar;" json:"customer_lan_on_site,omitempty"`
	Brand                  string            `gorm:"type:varchar;" json:"brand,omitempty"`
	Equipment              string            `gorm:"type:varchar;" json:"equipment,omitempty"`
	InstalledDate          *time.Time        `gorm:"type:date" json:"installed_date,omitempty"`
	InternetProvider1ID    int32             `json:"internet_provider1_id,omitempty"`
	InternetProvider2ID    int32             `json:"internet_provider2_id,omitempty"`
	FirewallIpv4LanAddress string            `gorm:"type:varchar;" json:"firewall_ipv4_lan_address,omitempty"`
	FirewallIpv6LanAddress string            `gorm:"type:varchar;" json:"firewall_ipv6_lan_address,omitempty"`
	CurrentVersion         string            `gorm:"type:varchar;" json:"current_version,omitempty"`
	VersionBackup          string            `gorm:"type:varchar;" json:"version_backup,omitempty"`
	BackupDate             *time.Time        `gorm:"type:date" json:"backup_date,omitempty"`
	Extra                  string            `gorm:"type:varchar;" json:"extra,omitempty"`
	CreatedAt              *time.Time        `gorm:"type:timestamp without time zone; not null" json:"created_at,omitempty"`
	CreatedByUserID        int32             `gorm:"not null;" json:"created_by_user_id,omitempty"`
	UpdatedAt              *time.Time        `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID        int32             `json:"updated_by_user_id,omitempty"`
	DeletedAt              *time.Time        `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID        int32             `json:"deleted_by_user_id,omitempty"`
	CustomerLocation       CustomerLocation  `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	InternetProvider1      *InternetProvider `gorm:"references:InternetProviderID;foreignKey:InternetProvider1ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"internet_provider1,omitempty"`
	InternetProvider2      *InternetProvider `gorm:"references:InternetProviderID;foreignKey:InternetProvider2ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"internet_provider2,omitempty"`
}

func (installedFirewallModel *InstalledFirewall) TableName() string {
	return "installed_firewalls"
}

func (installedFirewallModel *InstalledFirewall) CreateInstalledFirewall() error {
	var (
		tempProvider1 *InternetProvider
		tempProvider2 *InternetProvider
		createQuery   *gorm.DB
	)
	createInternetProvider1Error := installedFirewallModel.InternetProvider1.CreateInternetProvider()
	if createInternetProvider1Error != nil {
		return createInternetProvider1Error
	}
	installedFirewallModel.InternetProvider1ID = installedFirewallModel.InternetProvider1.InternetProviderID
	tempProvider1 = installedFirewallModel.InternetProvider1
	installedFirewallModel.InternetProvider1 = nil
	if installedFirewallModel.InternetProvider2 == nil {
		createQuery = initializers.DB().Omit("InternetProvider2ID").Create(&installedFirewallModel)
	} else {
		createInternetProvider2Error := installedFirewallModel.InternetProvider2.CreateInternetProvider()
		if createInternetProvider2Error != nil {
			return createInternetProvider2Error
		}
		installedFirewallModel.InternetProvider2ID = installedFirewallModel.InternetProvider2.InternetProviderID
		tempProvider2 = installedFirewallModel.InternetProvider2
		installedFirewallModel.InternetProvider2 = nil
		createQuery = initializers.DB().Create(&installedFirewallModel)
	}
	if installedFirewallModel.InstalledFirewallID > 0 {
		CreateLog(
			installedFirewallModel.CreatedByUserID,
			"Created",
			installedFirewallModel.TableName(),
			installedFirewallModel.InstalledFirewallID,
		)
		installedFirewallModel.InternetProvider1 = tempProvider1
		if tempProvider2 != nil {
			installedFirewallModel.InternetProvider2 = tempProvider2
		}
		return nil
	}
	return createQuery.Error
}

func (installedFirewallModel *InstalledFirewall) UpdateInstalledFirewall() error {
	if installedFirewallModel.CheckInstalledFirewallID(installedFirewallModel.InstalledFirewallID) {
		installedFirewallModel.InternetProvider1.CheckInternetProviderID(installedFirewallModel.InternetProvider1ID)
		installedFirewallModel.InternetProvider2.CheckInternetProviderID(installedFirewallModel.InternetProvider2ID)
		updateQuery := initializers.DB().Where("deleted_at IS NULL").
			Updates(&installedFirewallModel)
		if updateQuery.RowsAffected > 0 {
			CreateLog(
				installedFirewallModel.UpdatedByUserID,
				"Updated",
				installedFirewallModel.TableName(),
				installedFirewallModel.InstalledFirewallID,
			)
			return nil
		}
		return updateQuery.Error
	}
	return NotFoundError
}

func (installedFirewallModel *InstalledFirewall) CheckInstalledFirewallID(installedFirewallID int32) bool {
	findQuery := initializers.DB().
		Where("installed_firewall_id = ? AND deleted_at IS NULL", installedFirewallID).
		Select("installed_firewall_id,internet_provider1_id,internet_provider2_id").
		First(&installedFirewallModel)
	return findQuery.RowsAffected > 0
}

func (installedFirewallModel *InstalledFirewall) DeleteInstalledFirewall(installedFirewallIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND installed_firewall_id IN ?", installedFirewallIDs).
		Updates(&installedFirewallModel)
	if deleteQuery.RowsAffected > 0 {
		var logs []Log
		for index := range installedFirewallIDs {
			logs = append(logs, Log{
				UserID:    installedFirewallModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: installedFirewallModel.TableName(),
				DataID:    installedFirewallIDs[index],
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

func RetrieveInstalledFirewallRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                                                                                         int64
		installedFirewallIdQuery, versionBackupQuery, firewallWan1Ipv4Query, customerLocationIDQuery, backUpDateQuery string
	)

	if queries["installed_firewall_id"] != "" {
		installedFirewallIdQuery = " AND installed_firewalls.installed_firewall_id = " + queries["installed_firewall_id"]
	}

	if queries["version_backup"] != "" {
		versionBackupQuery = " AND installed_firewalls.version_backup ILIKE '%" + queries["version_backup"] + "%'"
	}

	if queries["firewall_wan1_ipv4"] != "" {
		firewallWan1Ipv4Query = " AND wan_configs1_ipv4.firewall = '" + queries["firewall_wan1_ipv4"] + "'"
	}

	if queries["customer_location_id"] != "" {
		customerLocationIDQuery = " AND installed_firewalls.customer_location_id = " + queries["customer_location_id"]
	}

	if queries["backup_date"] != "" {
		backUpDateQuery = " AND installed_firewalls.backup_date = '" + queries["backup_date"] + "'"
	}

	selectedColumns := "customer_locations.name as customer_locations,installed_firewalls.customer_lan_on_site, " +
		"installed_firewalls.brand,installed_firewalls.current_version,installed_firewalls.version_backup," +
		"installed_firewalls.backup_date,installed_firewalls.equipment,installed_firewalls.installed_date," +
		"wan_configs1_ipv4.firewall as firewall_wan1_ipv4, " +
		"wan_configs1_ipv4.sub_net_mask as sub_net_mask_wan1_ipv4, " +
		"wan_configs1_ipv4.gateway as gateway_wan1_ipv4," +
		"wan_configs1_ipv6.firewall as firewall_wan1_ipv6, " +
		"wan_configs1_ipv6.sub_net_mask as sub_net_mask_wan1_ipv6," +
		" wan_configs1_ipv6.gateway as gateway_wan1_ipv6," +
		"wan_configs2_ipv4.firewall as firewall_wan2_ipv4, " +
		"wan_configs2_ipv4.sub_net_mask as sub_net_mask_wan2_ipv4, " +
		"wan_configs2_ipv4.gateway as gateway_wan2_ipv4," +
		"wan_configs2_ipv6.firewall as firewall_wan2_ipv6, " +
		"wan_configs2_ipv6.sub_net_mask as sub_net_mask_wan2_ipv6, " +
		"wan_configs2_ipv6.gateway as gateway_wan2_ipv6," +
		"installed_firewalls.firewall_ipv4_lan_address," +
		"installed_firewalls.firewall_ipv6_lan_address,installed_firewalls.extra"

	if queries["select_column"] != "" {
		if queries["append_select"] == "true" {
			selectedColumns += "," + queries["select_column"]
		} else {
			selectedColumns = queries["select_column"]
		}
	}

	retrievedRows := initializers.DB().Scopes(pagination).Model(InstalledFirewall{}).
		Joins("INNER JOIN customer_locations " +
			"on customer_locations.customer_location_id = installed_firewalls.customer_location_id").
		Joins("INNER JOIN internet_providers as internet_providers1 " +
			"on internet_providers1.internet_provider_id = installed_firewalls.internet_provider1_id").
		Joins("INNER JOIN internet_providers as internet_providers2 " +
			"on internet_providers2.internet_provider_id = installed_firewalls.internet_provider2_id").
		Joins("INNER JOIN wan_configs as wan_configs1_ipv4 " +
			"on wan_configs1_ipv4.wan_config_id = internet_providers1.wan_config_ipv4_id").
		Joins("INNER JOIN wan_configs as wan_configs1_ipv6 " +
			"on wan_configs1_ipv6.wan_config_id = internet_providers1.wan_config_ipv6_id").
		Joins("INNER JOIN wan_configs as wan_configs2_ipv4 " +
			"on wan_configs2_ipv4.wan_config_id = internet_providers2.wan_config_ipv4_id").
		Joins("INNER JOIN wan_configs as wan_configs2_ipv6 " +
			"on wan_configs2_ipv6.wan_config_id = internet_providers2.wan_config_ipv6_id").
		Where("customer_locations.deleted_at IS NULL " +
			"AND installed_firewalls.deleted_at IS NULL" + installedFirewallIdQuery + versionBackupQuery +
			firewallWan1Ipv4Query + customerLocationIDQuery + backUpDateQuery).
		Select(selectedColumns).
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
