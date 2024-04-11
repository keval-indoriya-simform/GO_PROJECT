package models

import "time"

type UpdateCloudPrivateIp struct {
	Ipv4Assignment  string `gorm:"type:varchar;not null'" json:"ipv4_assignment,omitempty"`
	Ipv6Assignment  string `gorm:"type:varchar" json:"ipv6_assignment,omitempty"`
	Description     string `gorm:"type:text" json:"description,omitempty"`
	AssignedToID    int32  `gorm:"default:1" json:"assigned_to_id,omitempty"`
	UpdatedByUserID int32  `json:"updated_by_user_id,omitempty" swaggerignore:"true"`
}
type UpdateCloudPublicIp struct {
	IpAddress          string `gorm:"type:varchar;not null;" json:"ip_address,omitempty"`
	CustomerLocationID int32  `json:"customer_location_id,omitempty"`
	PostForwardIp      string `gorm:"type:varchar" json:"post_forward_ip,omitempty"`
	CloudVmName        string `gorm:"type:varchar" json:"cloud_vm_name,omitempty"`
	UpdatedByUserID    int32  `json:"updated_by_user_id,omitempty"`
}
type UpdateCustomerLan struct {
	CustomerLocationID               int32      `json:"customer_location_id,omitempty"`
	NetworkOnSite                    string     `gorm:"type:varchar;not null" json:"network_on_site,omitempty"`
	NumberOfInternetConnection       int        `gorm:"not null" json:"number_of_internet_connection,omitempty"`
	PrivateIpAssignmentID            int32      `json:"private_ip_assignment_id,omitempty"`
	FirewallIpv4LanAddress           string     `gorm:"type:varchar" json:"firewall_ipv4_lan_address,omitempty"`
	FirewallIpv6LanAddress           string     `gorm:"type:varchar" json:"firewall_ipv6_lan_address,omitempty"`
	GatewayType                      string     `gorm:"type:varchar" json:"gateway_type,omitempty"`
	EquipmentInstalled               string     `gorm:"type:varchar" json:"equipment_installed,omitempty"`
	BackupDate                       *time.Time `gorm:"type:date" json:"backup_date,omitempty"`
	VersionOfLastBackup              string     `gorm:"type:varchar;not null" json:"version_of_last_backup,omitempty"`
	Extra                            string     `gorm:"type:varchar" json:"extra,omitempty"`
	LanNotes                         string     `gorm:"type:varchar" json:"lan_notes,omitempty"`
	ScanToFolderLocation             string     `gorm:"type:varchar" json:"scan_to_folder_location,omitempty"`
	ScanToFolderUsernameOrPassword   string     `gorm:"type:varchar" json:"scan_to_folder_username_or_password,omitempty"`
	ScanToEmailSmtpServerPort        string     `gorm:"type:varchar" json:"scan_to_email_smtp_server_port,omitempty"`
	ScanToEmailEmailOrPassword       string     `gorm:"type:varchar" json:"scan_to_email_email_or_password,omitempty"`
	OnSiteBackupServerOrNasType      string     `gorm:"column(type);type:varchar" json:"on_site_backup_server_or_nas_type,omitempty"`
	OnSiteBackupServerOrNasIpAddress string     `gorm:"type:varchar" json:"on_site_backup_server_or_nas_ip_address,omitempty"`
	ManagementServer                 string     `gorm:"type:varchar" json:"management_server,omitempty"`
	ManagementNotes                  string     `gorm:"type:varchar" json:"management_notes,omitempty"`
	ManagementIpAddress              string     `gorm:"type:varchar" json:"management_ip_address,omitempty"`

	PrintOrScannerType      string `gorm:"column(type);type:varchar" json:"print_or_scanner_type,omitempty"`
	PrintOrScannerUserName  string `gorm:"type:varchar" json:"print_or_scanner_user_name,omitempty"`
	PrintOrScannerPassword  string `gorm:"type:varchar" json:"print_or_scanner_password,omitempty"`
	PrintOrScannerIpAddress string `gorm:"type:varchar" json:"print_or_scanner_ip_address,omitempty"`
	CreatedByUserID         int32  `gorm:"not null" json:"created_by_user_id,omitempty"`
}
type UpdateSwitches struct {
	SwitchBrandOrModel string     `gorm:"type:varchar" json:"switch_brand_or_model,omitempty"`
	SwitchCredentials  string     `gorm:"type:varchar" json:"switch_credentials,omitempty"`
	SwitchManage       string     `gorm:"type:varchar; default:('managed');" json:"switch_manage,omitempty"`
	SwitchIpAddress    string     `gorm:"type:varchar" json:"switch_ip_address,omitempty"`
	SwitchInstallDate  *time.Time `gorm:"type:date" json:"switch_install_date,omitempty"`
	SwitchNotes        string     `gorm:"type:varchar" json:"switch_notes,omitempty"`
	SwitchImageLinks   string     `gorm:"type:varchar" json:"switch_image_links,omitempty"`
}

type UpdateWireless struct {
	WirelessUnit           string `gorm:"type:varchar" json:"wireless_unit,omitempty"`
	WirelessIpAddress      string `gorm:"type:varchar" json:"wireless_ip_address,omitempty"`
	WirelessAdminUsername  string `gorm:"type:varchar" json:"wireless_admin_username,omitempty"`
	WirelessAdminPassword  string `gorm:"type:varchar" json:"wireless_admin_password,omitempty"`
	WirelessSsid           string `gorm:"type:varchar" json:"wireless_ssid,omitempty"`
	WirelessPassword       string `gorm:"type:varchar" json:"wireless_password,omitempty"`
	WirelessConnectionType string `gorm:"type:varchar" json:"wireless_connection_type,omitempty"`
	WirelessNotes          string `gorm:"type:varchar" json:"wireless_notes,omitempty"`
}
type UpdateInternetProvider struct {
	Name         string `gorm:"type:varchar" json:"name,omitempty"`
	AccountOrPin string `json:"account_or_pin,omitempty"`
	Other        string `gorm:"type:varchar" json:"other,omitempty"`
	Speed        string `gorm:"type:varchar" json:"speed,omitempty"`
	PrimaryDns   string `gorm:"type:varchar" json:"primary_dns,omitempty"`
	SecondaryDns string `gorm:"type:varchar" json:"secondary_dns,omitempty"`
}
type UpdateWanConfig struct {
	Firewall   string `gorm:"type:varchar" json:"firewall,omitempty"`
	SubNetMask string `gorm:"type:varchar" json:"sub_net_mask,omitempty"`
	Gateway    string `gorm:"type:varchar" json:"gateway,omitempty"`
	IpVersion  string `gorm:"type:varchar" json:"ip_version,omitempty"`
	IpRange    string `gorm:"type:varchar" json:"ip_range,omitempty"`
	Comment    string `gorm:"type:varchar" json:"comment,omitempty"`
}
type UpdateCustomerLocation struct {
	CustomerID      int32  `gorm:"not null;" json:"customer_id,omitempty"`
	IsPrimary       *bool  `gorm:"not null" json:"is_primary,omitempty"`
	Name            string `gorm:"type:varchar; not null;" json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	UpdatedByUserID int32  `json:"updated_by_user_id,omitempty"`
}
type UpdateCustomer struct {
	Name               string `gorm:"type:varchar(30);not null" json:"name,omitempty"`
	IsActive           *bool  `gorm:"" json:"is_active,omitempty"`
	CloudOrOnsiteID    int32  `gorm:"default:1" json:"cloud_or_onsite_id,omitempty"`
	Voip               *bool  `gorm:"" json:"voip,omitempty"`
	Internet           *bool  `gorm:"" json:"internet,omitempty"`
	Firewall           *bool  `gorm:"" json:"firewall,omitempty"`
	BackupSoftware     string `gorm:"type:varchar" json:"backup_software,omitempty"`
	HardwareAsAService *bool  `gorm:"" json:"hardware_as_a_service,omitempty"`
	Description        string `gorm:"type:varchar" json:"description,omitempty"`
	UpdatedByUserID    int32  `gorm:"" json:"updated_by_user_id,omitempty"`
}

type UpdateEmailDomain struct {
	CustomerLocationID int32  `gorm:"type:varchar" json:"customer_location_id,omitempty"`
	DomainRegistrar    string `gorm:"type:varchar" json:"domain_registrar,omitempty"`
	Password           string `gorm:"type:varchar" json:"password,omitempty"`
	Pin                string `gorm:"type:varchar" json:"pin,omitempty"`
	Domain             string `gorm:"type:varchar;not null" json:"domain"`
	ARecord1           string `gorm:"type:varchar" json:"a_record_1,omitempty"`
	ARecord2           string `gorm:"type:varchar" json:"a_record_2,omitempty"`
	ARecord3           string `gorm:"type:varchar" json:"a_record_3,omitempty"`
	ARecord4           string `gorm:"type:varchar" json:"a_record_4,omitempty"`
	MxRecord1          string `gorm:"type:varchar" json:"mx_record_1,omitempty"`
	MxRecord2          string `gorm:"type:varchar" json:"mx_record_2,omitempty"`
	WebsiteIpOrAlias   string `gorm:"type:varchar" json:"website_ip_or_alias,omitempty"`
	WebMailPopImap     string `gorm:"type:varchar" json:"web_mail_pop_imap,omitempty"`
	WebMailExchange    string `gorm:"type:varchar" json:"web_mail_exchange,omitempty"`
	PhoneSettingNote   string `gorm:"type:varchar" json:"phone_setting_note,omitempty"`
	Notes              string `gorm:"type:varchar" json:"notes,omitempty"`
	UpdatedByUserID    int32  `json:"updated_by_user_id,omitempty"`
}
type UpdateAccount struct {
	LinkForDomainAdmin   string `gorm:"type:varchar" json:"link_for_domain_admin,omitempty"`
	AccountNumber        string `gorm:"type:varchar" json:"account_number,omitempty"`
	Name                 string `gorm:"type:varchar" json:"name,omitempty"`
	AccountNumberAliases string `gorm:"type:varchar" json:"account_number_aliases,omitempty"`
	UserName             string `gorm:"type:varchar" json:"user_name,omitempty"`
}
type UpdateEmailAccount struct {
	EmailAccountID     int32  `gorm:"primaryKey" json:"-"`
	EmailHosting       string `gorm:"type:varchar" json:"email_hosting,omitempty"`
	LinkForEmailAdmin  string `gorm:"type:varchar" json:"link_for_email_admin,omitempty"`
	AccountNumber      string `gorm:"type:varchar" json:"account_number,omitempty"`
	EmailAccountTypeID int32  `gorm:"not null;"  json:"email_account_type_id,omitempty"`
	UserName           string `gorm:"type:varchar" json:"user_name,omitempty"`
	Password           string `gorm:"type:varchar" json:"password,omitempty"`
	Pin                string `gorm:"type:varchar" json:"pin,omitempty"`
}
type UpdateInstalledFirewall struct {
	CustomerLocationID     int32      `json:"customer_location_id,omitempty"`
	CustomerLanOnSite      string     `gorm:"type:varchar;" json:"customer_lan_on_site,omitempty"`
	Brand                  string     `gorm:"type:varchar;" json:"brand,omitempty"`
	Equipment              string     `gorm:"type:varchar;" json:"equipment,omitempty"`
	InstalledDate          *time.Time `gorm:"type:date" json:"installed_date,omitempty"`
	FirewallIpv4LanAddress string     `gorm:"type:varchar;" json:"firewall_ipv4_lan_address,omitempty"`
	FirewallIpv6LanAddress string     `gorm:"type:varchar;" json:"firewall_ipv6_lan_address,omitempty"`
	CurrentVersion         string     `gorm:"type:varchar;" json:"current_version,omitempty"`
	VersionBackup          string     `gorm:"type:varchar;" json:"version_backup,omitempty"`
	BackupDate             *time.Time `gorm:"type:date" json:"backup_date,omitempty"`
	Extra                  string     `gorm:"type:varchar;" json:"extra,omitempty"`
	UpdatedByUserID        int32      `json:"updated_by_user_id,omitempty"`
}

type UpdateNote struct {
	Subject            string `gorm:"type:varchar;not null" json:"subject,omitempty"`
	CustomerLocationID int32  `json:"customer_location_id,omitempty"`
	Attachment         string `gorm:"type:varchar" json:"attachment,omitempty"`
	Note               string `gorm:"type:varchar" json:"note,omitempty"`
	AssignedToID       int32  `gorm:"default:1" json:"assigned_to_id,omitempty"`
	UpdatedByUserID    int32  `json:"updated_by_user_id,omitempty"`
}

type UpdateServer struct {
	HostName              string     `gorm:"type:varchar" json:"host_name,omitempty"`
	CustomerLocationID    int32      `json:"customer_location_id,omitempty"`
	HardwareAsAService    *bool      `gorm:"column(hardware_as_a_service)" json:"hardware_as_a_service,omitempty"`
	OsPlatform            string     `gorm:"type:varchar" json:"os_platform,omitempty"`
	ServiceTag            string     `gorm:"type:varchar" json:"service_tag,omitempty"`
	ExpressionServiceCode int        `json:"expression_service_code,omitempty"`
	Location              string     `gorm:"type:varchar" json:"location,omitempty"`
	Warranty              string     `gorm:"type:varchar" json:"warranty,omitempty"`
	Type                  string     `gorm:"type:varchar;" json:"type,omitempty"`
	PowerConnectType      string     `gorm:"type:varchar; not null" json:"power_connect_type,omitempty"`
	PurchaseDate          *time.Time `gorm:"type:date" json:"purchase_date,omitempty"`
	ExpirationDate        *time.Time `gorm:"type:date" json:"expiration_date,omitempty"`
	DaysLeft              string     `gorm:"type:varchar" json:"days_left,omitempty"`
	Ownership             string     `gorm:"type:varchar" json:"ownership,omitempty"`
	OrderNumber           string     `gorm:"type:varchar" json:"order_number,omitempty"`
	Description           string     `gorm:"type:varchar" json:"description,omitempty"`
	Idrac                 string     `gorm:"type:varchar" json:"idrac,omitempty"`
	UpdatedByUserID       int32      `json:"updated_by_user_id,omitempty"`
}
type UpdateSoftware struct {
	CustomerLocationID int32      `gorm:"not null;" json:"customer_location_id,omitempty"`
	Name               string     `gorm:"type:varchar" json:"name,omitempty"`
	Version            string     `gorm:"type:varchar" json:"version,omitempty"`
	LicenseKey         string     `gorm:"type:varchar;not null" json:"license_key,omitempty"`
	ServerOrVM         string     `gorm:"type:varchar" json:"server_or_vm,omitempty"`
	OtherLicenseInfo   string     `gorm:"type:varchar" json:"other_license_info,omitempty"`
	InstallDate        *time.Time `gorm:"type:date" json:"install_date,omitempty"`
	ExpiryDate         *time.Time `gorm:"type:date" json:"expiry_date,omitempty"`
	Notes              string     `gorm:"type:varchar" json:"notes,omitempty"`
	UpdatedByUserID    int32      `json:"updated_by_user_id,omitempty"`
}
