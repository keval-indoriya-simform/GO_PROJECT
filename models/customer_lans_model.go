package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type CustomerLan struct {
	CustomerLanID                    int32             `gorm:"primaryKey;autoIncrement;" json:"customer_lan_id,omitempty"`
	CustomerLocationID               int32             `json:"customer_location_id,omitempty"`
	NetworkOnSite                    string            `gorm:"type:varchar;not null" json:"network_on_site,omitempty"`
	NumberOfInternetConnection       int               `gorm:"not null" json:"number_of_internet_connection,omitempty"`
	InternetProvider1ID              int32             `gorm:"not null" json:"internet_provider1_id,omitempty"`
	InternetProvider2ID              int32             `json:"internet_provider2_id,omitempty"`
	PrivateIpAssignmentID            int32             `json:"private_ip_assignment_id,omitempty"`
	FirewallIpv4LanAddress           string            `gorm:"type:varchar" json:"firewall_ipv4_lan_address,omitempty"`
	FirewallIpv6LanAddress           string            `gorm:"type:varchar" json:"firewall_ipv6_lan_address,omitempty"`
	GatewayType                      string            `gorm:"type:varchar" json:"gateway_type,omitempty"`
	EquipmentInstalled               string            `gorm:"type:varchar" json:"equipment_installed,omitempty"`
	BackupDate                       *time.Time        `gorm:"type:date" json:"backup_date,omitempty"`
	VersionOfLastBackup              string            `gorm:"type:varchar;not null" json:"version_of_last_backup,omitempty"`
	Extra                            string            `gorm:"type:varchar" json:"extra,omitempty"`
	LanNotes                         string            `gorm:"type:varchar" json:"lan_notes,omitempty"`
	ScanToFolderLocation             string            `gorm:"type:varchar" json:"scan_to_folder_location,omitempty"`
	ScanToFolderUsernameOrPassword   string            `gorm:"type:varchar" json:"scan_to_folder_username_or_password,omitempty"`
	ScanToEmailSmtpServerPort        string            `gorm:"type:varchar" json:"scan_to_email_smtp_server_port,omitempty"`
	ScanToEmailEmailOrPassword       string            `gorm:"type:varchar" json:"scan_to_email_email_or_password,omitempty"`
	OnSiteBackupServerOrNasType      string            `gorm:"column(type);type:varchar" json:"on_site_backup_server_or_nas_type,omitempty"`
	OnSiteBackupServerOrNasIpAddress string            `gorm:"type:varchar" json:"on_site_backup_server_or_nas_ip_address,omitempty"`
	ManagementServer                 string            `gorm:"type:varchar" json:"management_server,omitempty"`
	ManagementNotes                  string            `gorm:"type:varchar" json:"management_notes,omitempty"`
	ManagementIpAddress              string            `gorm:"type:varchar" json:"management_ip_address,omitempty"`
	WirelessUnit                     string            `gorm:"type:varchar" json:"wireless_unit,omitempty"`
	WirelessIpAddress                string            `gorm:"type:varchar" json:"wireless_ip_address,omitempty"`
	WirelessAdminUsername            string            `gorm:"type:varchar" json:"wireless_admin_username,omitempty"`
	WirelessAdminPassword            string            `gorm:"type:varchar" json:"wireless_admin_password,omitempty"`
	WirelessSsid                     string            `gorm:"type:varchar" json:"wireless_ssid,omitempty"`
	WirelessPassword                 string            `gorm:"type:varchar" json:"wireless_password,omitempty"`
	WirelessConnectionType           string            `gorm:"type:varchar" json:"wireless_connection_type,omitempty"`
	WirelessNotes                    string            `gorm:"type:varchar" json:"wireless_notes,omitempty"`
	SwitchBrandOrModel               string            `gorm:"type:varchar" json:"switch_brand_or_model,omitempty"`
	SwitchCredentials                string            `gorm:"type:varchar" json:"switch_credentials,omitempty"`
	SwitchManage                     string            `gorm:"type:varchar; default:('managed');" json:"switch_manage,omitempty"`
	SwitchIpAddress                  string            `gorm:"type:varchar" json:"switch_ip_address,omitempty"`
	SwitchInstallDate                *time.Time        `gorm:"type:date" json:"switch_install_date,omitempty"`
	SwitchNotes                      string            `gorm:"type:varchar" json:"switch_notes,omitempty"`
	SwitchImageLinks                 string            `gorm:"type:varchar" json:"switch_image_links,omitempty"`
	PrintOrScannerType               string            `gorm:"column(type);type:varchar" json:"print_or_scanner_type,omitempty"`
	PrintOrScannerUserName           string            `gorm:"type:varchar" json:"print_or_scanner_user_name,omitempty"`
	PrintOrScannerPassword           string            `gorm:"type:varchar" json:"print_or_scanner_password,omitempty"`
	PrintOrScannerIpAddress          string            `gorm:"type:varchar" json:"print_or_scanner_ip_address,omitempty"`
	CreatedAt                        *time.Time        `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID                  int32             `gorm:"not null" json:"created_by_user_id,omitempty"`
	UpdatedAt                        *time.Time        `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID                  int32             `json:"updated_by_user_id,omitempty"`
	DeletedAt                        *time.Time        `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID                  int32             `json:"deleted_by_user_id,omitempty"`
	CustomerLocation                 *CustomerLocation `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	InternetProvider1                *InternetProvider `gorm:"references:InternetProviderID;foreignKey:InternetProvider1ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"internet_provider1,omitempty"`
	InternetProvider2                *InternetProvider `gorm:"references:InternetProviderID;foreignKey:InternetProvider2ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"internet_provider2,omitempty"`
	CloudPrivateIp                   *CloudPrivateIp   `gorm:"references:CloudPrivateIPId;foreignKey:PrivateIpAssignmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"cloud_private_ip,omitempty"`
}

func (customerLanModel *CustomerLan) TableName() string {
	return "customer_lans"
}

func (customerLanModel *CustomerLan) CreateCustomerLan() error {
	var (
		tempProvider1, tempProvider2 *InternetProvider
		createQuery                  *gorm.DB
	)
	createInternetProvider1Error := customerLanModel.InternetProvider1.CreateInternetProvider()
	if createInternetProvider1Error != nil {
		return createInternetProvider1Error
	}
	customerLanModel.InternetProvider1ID = customerLanModel.InternetProvider1.InternetProviderID
	tempProvider1 = customerLanModel.InternetProvider1
	customerLanModel.InternetProvider1 = nil
	if customerLanModel.InternetProvider2 == nil {
		createQuery = initializers.DB().Omit("InternetProvider2ID").Create(&customerLanModel)
	} else {
		createInternetProvider2Error := customerLanModel.InternetProvider2.CreateInternetProvider()
		if createInternetProvider2Error != nil {
			return createInternetProvider2Error
		}
		customerLanModel.InternetProvider2ID = customerLanModel.InternetProvider2.InternetProviderID
		tempProvider2 = customerLanModel.InternetProvider2
		customerLanModel.InternetProvider2 = nil
		createQuery = initializers.DB().Create(&customerLanModel)
	}
	if customerLanModel.CustomerLanID > 0 {
		CreateLog(
			customerLanModel.CreatedByUserID,
			"Created",
			customerLanModel.TableName(),
			customerLanModel.CustomerLanID,
		)
		customerLanModel.InternetProvider1 = tempProvider1
		if tempProvider2 != nil {
			customerLanModel.InternetProvider2 = tempProvider2
		}
		return nil
	}
	return createQuery.Error
}

func (customerLanModel *CustomerLan) UpdateCustomerLan() error {
	if customerLanModel.CheckCustomerLanID(customerLanModel.CustomerLanID) {
		updateQuery := initializers.DB().
			Updates(customerLanModel)
		if updateQuery.RowsAffected > 0 {
			CreateLog(
				customerLanModel.UpdatedByUserID,
				"Updated",
				customerLanModel.TableName(),
				customerLanModel.CustomerLanID,
			)
			return nil
		}
		return updateQuery.Error
	}
	return NotFoundError
}

func (customerLanModel *CustomerLan) CheckCustomerLanID(customerLanID int32) bool {
	findQuery := initializers.DB().Model(CustomerLan{}).
		Where("customer_lan_id = ? AND deleted_at IS NULL", customerLanID).
		Select("customer_lan_id, customer_location_id, internet_provider1_id, internet_provider2_id, " +
			"private_ip_assignment_id").
		First(&customerLanModel)
	if findQuery.RowsAffected > 0 {
		customerLanModel.InternetProvider1.CheckInternetProviderID(customerLanModel.InternetProvider1ID)
		customerLanModel.InternetProvider2.CheckInternetProviderID(customerLanModel.InternetProvider2ID)
		return true
	}
	return false
}

func (customerLanModel *CustomerLan) DeleteCustomerLan(customerLanIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND customer_lan_id IN ?", customerLanIDs).
		Updates(&customerLanModel)
	if deleteQuery.RowsAffected > 0 {
		CreateLog(
			customerLanModel.DeletedByUserID,
			"Deleted",
			customerLanModel.TableName(),
			customerLanModel.CustomerLocationID,
		)
		var logs []Log
		for index, _ := range customerLanIDs {
			logs = append(logs, Log{
				UserID:    customerLanModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: customerLanModel.TableName(),
				DataID:    customerLanIDs[index],
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

func RetrieveCustomerLanRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                                                                   int64
		customerLanIDQuery, customerLocationQuery, customerLanNetworkOnSiteQuery, assignToQuery string
	)

	if queries["customer_lan_id"] != "" {
		customerLanIDQuery = " AND customer_lans.customer_lan_id = " + queries["customer_lan_id"]
	}

	if queries["customer_location"] != "" {
		customerLocationQuery = "AND customer_locations.name Ilike '%" + queries["customer_location"] + "%'"
	}
	if queries["customer_lan_network_on_site"] != "" {
		customerLanNetworkOnSiteQuery = "AND customer_lans.network_on_site Ilike '%" + queries["customer_lan_network_on_site"] + "%'"
	}
	if queries["assign_to"] != "" {
		assignToQuery = "AND customer_lans.created_by_user_id = '" + queries["assign_to"] + "'"
	}

	selectedColumns := "customer_lans.customer_lan_id,customer_locations.name as customer_location, " +
		"customer_lans.network_on_site as network_on_site, " +
		"customer_lans.number_of_internet_connection," +
		"internet_providers1.name as internet_providers1_name, " +
		"internet_providers2.name as internet_providers2_name," +
		"customer_lans.management_server, " +
		"internet_providers1.account_or_pin as internet_providers1_account_or_pin," +
		" internet_providers1.speed as internet_providers1_speed, internet_providers2.account_or_pin as " +
		"internet_providers2_account_or_pin, internet_providers2.speed as internet_providers2_speed," +
		"wan_configs1_ipv4.firewall as firewall_wan1_ipv4, " +
		"wan_configs1_ipv4.sub_net_mask as sub_net_mask_wan1_ipv4, " +
		"wan_configs1_ipv4.gateway as gateway_wan1_ipv4," +
		"wan_configs1_ipv6.firewall as firewall_wan1_ipv6, " +
		"wan_configs1_ipv6.sub_net_mask as sub_net_mask_wan1_ipv6, " +
		"wan_configs1_ipv6.gateway as gateway_wan1_ipv6," +
		"wan_configs2_ipv4.firewall as firewall_wan2_ipv4, " +
		"wan_configs2_ipv4.sub_net_mask as sub_net_mask_wan2_ipv4, " +
		"wan_configs2_ipv4.gateway as gateway_wan2_ipv4," +
		"wan_configs2_ipv6.firewall as firewall_wan2_ipv6," +
		" wan_configs2_ipv6.sub_net_mask as sub_net_mask_wan2_ipv6, " +
		"wan_configs2_ipv6.gateway as gateway_wan2_ipv6, " +
		"customer_lans.equipment_installed, internet_providers1.primary_dns as internet_providers1_primary_dns, " +
		"internet_providers1.secondary_dns as internet_providers1_secondary_dns," +
		" internet_providers2.primary_dns as internet_providers2_primary_dns," +
		" internet_providers2.secondary_dns as internet_providers2_secondary_dns, " +
		"customer_lans.firewall_ipv4_lan_address, customer_lans.firewall_ipv6_lan_address, " +
		"customer_lans.gateway_type, customer_lans.wireless_unit, customer_lans.backup_date," +
		" customer_lans.version_of_last_backup, customer_lans.extra, customer_lans.print_or_scanner_ip_address, " +
		"customer_lans.scan_to_folder_location, customer_lans.scan_to_folder_username_or_password, " +
		"cloud_private_ips.ipv4_assignment, customer_lans.scan_to_email_smtp_server_port, " +
		"customer_lans.scan_to_email_email_or_password, customer_lans.print_or_scanner_user_name, " +
		"customer_lans.print_or_scanner_password"

	if queries["select_column"] != "" {
		if queries["append_select"] == "true" {
			selectedColumns += "," + queries["select_column"]
		} else {
			selectedColumns = queries["select_column"]
		}
	}

	retrievedRows := initializers.DB().Model(CustomerLan{}).Scopes(pagination).Where("customer_lans.deleted_at IS NULL " +
		customerLanIDQuery + customerLocationQuery + customerLanNetworkOnSiteQuery + assignToQuery + "").
		Joins("INNER JOIN customer_locations " +
			"on customer_locations.customer_location_id = customer_lans.customer_location_id").
		Joins("INNER JOIN internet_providers as internet_providers1 " +
			"on internet_providers1.internet_provider_id = customer_lans.internet_provider1_id").
		Joins("INNER JOIN internet_providers as internet_providers2 " +
			"on internet_providers2.internet_provider_id = customer_lans.internet_provider2_id").
		Joins("INNER JOIN wan_configs as wan_configs1_ipv4 " +
			"on wan_configs1_ipv4.wan_config_id = internet_providers1.wan_config_ipv4_id").
		Joins("INNER JOIN wan_configs as wan_configs1_ipv6 " +
			"on wan_configs1_ipv6.wan_config_id = internet_providers1.wan_config_ipv6_id").
		Joins("INNER JOIN wan_configs as wan_configs2_ipv4 " +
			"on wan_configs2_ipv4.wan_config_id = internet_providers2.wan_config_ipv4_id").
		Joins("INNER JOIN wan_configs as wan_configs2_ipv6 " +
			"on wan_configs2_ipv6.wan_config_id = internet_providers2.wan_config_ipv6_id").
		Joins("INNER JOIN cloud_private_ips " +
			"on cloud_private_ips.cloud_private_ip_id = customer_lans.private_ip_assignment_id").
		Select(selectedColumns).
		Find(&records)

	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
