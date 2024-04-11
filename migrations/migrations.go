package main

import (
	"Application/initializers"
	"Application/models"
	"log"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
}

type AlterUniqueConstraint struct {
	TableName      string
	ConstraintName string
	Columns        string
}

type CreateFunction struct {
	FunctionName                      string
	ReturnType                        string
	Language                          string
	MaxCost                           string
	VolatileParallel                  string
	MaxRows                           string
	ExecutionBodyIncludingBeginAndEnd string
}

func addUniqueConstraint(constraints ...AlterUniqueConstraint) {
	for _, constraint := range constraints {
		alterUniqueConstraintError := initializers.DB().Exec("ALTER TABLE " + constraint.TableName + " ADD CONSTRAINT " + constraint.ConstraintName + " UNIQUE (" + constraint.Columns + ")").Error
		if alterUniqueConstraintError != nil {
			log.Fatal(alterUniqueConstraintError)
		}
	}
}

func AddFunction(functions ...CreateFunction) {
	for _, function := range functions {
		var returnQuery, languageQuery, costQuery, volatileParallelQuery, maxRowsQuery string
		if function.ReturnType != "" {
			returnQuery = " RETURNS " + function.ReturnType
		}
		if function.Language != "" {
			languageQuery = " LANGUAGE " + function.Language
		} else {
			languageQuery = " LANGUAGE 'sql'"
		}
		if function.MaxCost != "" {
			costQuery = " COST " + function.MaxCost
		} else {
			costQuery = " COST 100"
		}
		if function.VolatileParallel != "" {
			volatileParallelQuery = " VOLATILE PARALLEL " + function.VolatileParallel
		} else {
			volatileParallelQuery = " VOLATILE PARALLEL UNSAFE"
		}
		if function.MaxRows != "" {
			maxRowsQuery = " ROWS " + function.MaxRows
		} else {
			maxRowsQuery = " ROWS 1000"
		}

		AddFunctionError := initializers.DB().Exec("CREATE OR REPLACE FUNCTION " + function.FunctionName +
			returnQuery + languageQuery + costQuery + volatileParallelQuery + maxRowsQuery + " AS $BODY$ " +
			function.ExecutionBodyIncludingBeginAndEnd + " $BODY$;").Error
		if AddFunctionError != nil {
			log.Fatal(AddFunctionError)
		}
	}
}

func main() {
	err := initializers.DB().Migrator().DropTable(
		&models.Role{},
		&models.User{},
		&models.UserRole{},
		&models.CloudOrOnsite{},
		&models.Customer{},
		&models.CustomerLocation{},
		&models.CloudPublicIp{},
		&models.WanConfig{},
		&models.InternetProvider{},
		&models.Account{},
		&models.CloudPrivateIp{},
		&models.CustomerLan{},
		&models.Software{},
		&models.EmailAccountType{},
		&models.EmailAccount{},
		&models.EmailDomain{},
		&models.Server{},
		&models.InstalledFirewall{},
		&models.Note{},
		&models.Log{},
	)

	err = initializers.DB().AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.UserRole{},
		&models.CloudOrOnsite{},
		&models.Customer{},
		&models.CustomerLocation{},
		&models.CloudPublicIp{},
		&models.WanConfig{},
		&models.InternetProvider{},
		&models.Account{},
		&models.CloudPrivateIp{},
		&models.CustomerLan{},
		&models.Software{},
		&models.EmailAccountType{},
		&models.EmailAccount{},
		&models.EmailDomain{},
		&models.Server{},
		&models.InstalledFirewall{},
		&models.Note{},
		&models.Log{},
	)
	if err != nil {
		return
	}

	// Creating a Unique Constraint for multiple fields together
	addUniqueConstraint(
		AlterUniqueConstraint{ //CASE STUDY PROJECT VALUE UNIQUE CONSTRAINTS
			TableName:      "customer_locations",
			ConstraintName: "customer_locations_un",
			Columns:        "name, customer_id",
		},
	)

	AddFunction(
		CreateFunction{
			FunctionName: "get_counts()",
			ReturnType:   "TABLE(field_name text, field_count bigint)",
			Language:     "plpgsql",
			ExecutionBodyIncludingBeginAndEnd: `begin RETURN QUERY (SELECT 'active_customer_location', ` +
				`count(customer_location_id) FROM public.customer_locations where deleted_at IS NULL)` +
				` union ` +
				`(SELECT 'total_customers',count(*) FROM public.customers)` +
				` union ` +
				`(SELECT 'active_customers_' || is_active, count(is_active) FROM public.customers group by is_active)` +
				` union ` +
				`(SELECT 'cloud_or_onsites_' || cloud_or_onsites.name, count(cloud_or_onsite_id) FROM ` +
				`public.customers join cloud_or_onsites using(cloud_or_onsite_id) group by cloud_or_onsites.name)` +
				` union ` +
				`(SELECT 'voip_' || VOIP, count(VOIP) FROM public.customers group by VOIP)` +
				` union ` +
				`(SELECT 'internet_' || internet, count(internet) FROM public.customers group by internet)` +
				` union ` +
				`(SELECT 'firewall_' || firewall, count(firewall) FROM public.customers group by firewall)` +
				` union ` +
				`(SELECT 'gatewat_type', count(installed_firewall_id) FROM public.installed_firewalls)` +
				` union ` +
				`(SELECT 'gatewat_type_' || brand, count(brand) FROM public.installed_firewalls group by brand)` +
				` union ` +
				`(SELECT 'equipment_' || equipment, count(equipment) FROM public.installed_firewalls group by equipment)` +
				` union ` +
				`(SELECT 'current_version_' || current_version, count(current_version) FROM ` +
				`public.installed_firewalls group by current_version)` +
				` union ` +
				`(SELECT 'internet_provider_' || internet_providers.name, count(internet_providers.name) FROM ` +
				`public.internet_providers group by internet_providers.name)` +
				` union ` +
				`(SELECT 'email_hosting_' || email_hosting, count(email_hosting) FROM ` +
				`public.email_accounts group by email_hosting)` +
				` union ` +
				`(SELECT 'software_'  || softwares.name, count(softwares.name) FROM ` +
				`public.softwares group by softwares.name); end;`,
		},
	)

	createDummyData()

}
func createDummyData() {
	var (
		currentTime = time.Now()
		boolValue   = true
	)

	initializers.DB().Create(
		&models.Role{
			RoleName: "Super Admin",
		})
	initializers.DB().Create(

		&models.Role{
			RoleName: "Admin",
		})
	initializers.DB().Create(
		&models.Role{
			RoleName: "Employee",
		})
	initializers.DB().Create(
		&models.User{
			Email:     "admin@gmail.com",
			Username:  "admin",
			Password:  "$2a$14$8AeQTIuvRzMm9oUQ1/LUe.EWIADjJSVPqi6xRuOuy.3eIHq3FtRjG",
			CreatedAt: &currentTime,
			UpdatedAt: nil,
			IsActive:  &boolValue,
		})

	initializers.DB().Create(
		&models.User{
			Name:       "Abhishek Mali",
			Email:      "abhishek.m@simformsolutions.com",
			Department: "Tech",
			Username:   "abhi",
			Password:   "$2a$14$iya6WWujzFcaOgJx3VDC1OhK9ZYc6FHFfhpaPuui8iRtm4pHAvuGq",
			CreatedAt:  &currentTime,
			UpdatedAt:  nil,
			IsActive:   &boolValue,
		})

	initializers.DB().Create(
		&models.User{
			Name:       "NMS",
			Email:      "nmsd@gmail.com",
			Department: "Tech",
			Username:   "nms",
			Password:   "$2a$14$JYkzvxv11dJHqnhfjxs3dOWM40mWv5UPOqIYP62f3BI7bFmnpopu6",
			CreatedAt:  &currentTime,
			UpdatedAt:  nil,
			IsActive:   &boolValue,
		})

	initializers.DB().Create(
		&models.UserRole{
			RoleID:          1,
			UserID:          1,
			CreatedByUserID: 1,
			CreatedAt:       &currentTime,
		})

	initializers.DB().Create(
		&models.UserRole{
			RoleID:          2,
			UserID:          2,
			CreatedByUserID: 1,
			CreatedAt:       &currentTime,
		})

	initializers.DB().Create(
		&models.UserRole{
			RoleID:          3,
			UserID:          3,
			CreatedByUserID: 1,
			CreatedAt:       &currentTime,
		})
	initializers.DB().Create(
		&models.CloudOrOnsite{
			Name: "On Site No Server",
		})

	initializers.DB().Create(
		&models.CloudOrOnsite{
			Name: "On Site",
		})

	initializers.DB().Create(
		&models.CloudOrOnsite{
			Name: "Hybrid",
		})

	initializers.DB().Create(
		&models.CloudOrOnsite{
			Name: "Cloud",
		})

	initializers.DB().Create(
		&models.CloudOrOnsite{
			Name: "VM Hosting Only",
		})

	initializers.DB().Create(&models.Customer{
		Name:               "john doe",
		IsActive:           &boolValue,
		CloudOrOnsiteID:    1,
		Voip:               &boolValue,
		Internet:           &boolValue,
		Firewall:           &boolValue,
		BackupSoftware:     "",
		HardwareAsAService: &boolValue,
		Description:        "",
		CreatedAt:          &currentTime,
		CreatedByUserID:    1,
	})

	initializers.DB().Create(&models.CustomerLocation{
		CustomerID:      1,
		Name:            "simform solution 6th floor",
		CreatedAt:       &currentTime,
		CreatedByUserID: 1,
		IsPrimary:       &boolValue,
	})

	initializers.DB().Create(&models.CloudPublicIp{
		IpAddress:          "196.172.0.4",
		CustomerLocationID: 1,
		PostForwardIp:      "198.124.1.55",
		CloudVmName:        "amazon workspace",
		CreatedAt:          &currentTime,
		CreatedByUserID:    1,
	})

	initializers.DB().Create(&models.WanConfig{
		Firewall:   "",
		SubNetMask: "255.255.255.224",
		Gateway:    "127.0.0.1",
		IpVersion:  "4",
		IpRange:    "14.99.102.225 - 14.99.102.226",
		Comment:    "no comments",
	})

	initializers.DB().Create(&models.WanConfig{
		Firewall:   "",
		SubNetMask: "2001:db8:3333:4444:5555:6666:7777:8888/64",
		Gateway:    "2001:db8:3333:4444::",
		IpVersion:  "6",
		IpRange:    "2001:0db8:85a3:0000:0000:0000:0000:0000 - 2001:0db8:85a3:0000:ffff:ffff:ffff:ffff",
		Comment:    "no comments",
	})

	internetProvider1 := &models.InternetProvider{
		Name:            "Spectrum Fiber",
		AccountOrPin:    "323129941 PIN #: 5336",
		Other:           "other",
		Speed:           "30x30 meg",
		PrimaryDns:      "8.8.8.8",
		SecondaryDns:    "8.8.4.4",
		WanConfigIpv4ID: 1,
		WanConfigIpv6ID: 2,
	}
	internetProvider2 := &models.InternetProvider{
		Name:            "Spectrum Fiber Pro",
		AccountOrPin:    "323129952 PIN #: 5662",
		Other:           "other",
		Speed:           "50x50 meg",
		PrimaryDns:      "75.75.75.75",
		SecondaryDns:    "75.75.75.76",
		WanConfigIpv4ID: 1,
		WanConfigIpv6ID: 2,
	}
	account := &models.Account{
		LinkForDomainAdmin:   "https://hosting.digitalspaceportal.net/portal/",
		AccountNumber:        "35829125",
		Name:                 "ACME Steaks",
		AccountNumberAliases: "sdsdds",
		UserName:             "Petrina Manages",
	}
	cloudPrivateIp := &models.CloudPrivateIp{
		Ipv4Assignment:  "10.0.24.014",
		Ipv6Assignment:  "2607:f758:5080:fc::27",
		Description:     "hello",
		CreatedAt:       &currentTime,
		CreatedByUserID: 1,
	}
	initializers.DB().Create(internetProvider1)
	initializers.DB().Create(internetProvider2)
	initializers.DB().Create(account)
	initializers.DB().Create(cloudPrivateIp)
	initializers.DB().Create(
		&models.CustomerLan{
			CustomerLocationID:               1,
			NetworkOnSite:                    "192.168.52.0/24",
			NumberOfInternetConnection:       1,
			InternetProvider1ID:              1,
			InternetProvider2ID:              2,
			PrivateIpAssignmentID:            1,
			FirewallIpv4LanAddress:           "12.161.34.34",
			FirewallIpv6LanAddress:           "2001:1890:156d:7300::4",
			GatewayType:                      "SonicWall (ACME)",
			EquipmentInstalled:               "SonicWall Customer Equipment (ACME)",
			BackupDate:                       &currentTime,
			VersionOfLastBackup:              "n/a",
			Extra:                            "Firewall User: Done Customer provider ipSec VPN",
			LanNotes:                         "This is Lan Note",
			SwitchBrandOrModel:               "Cisco",
			SwitchCredentials:                "root_admin",
			SwitchManage:                     "Managed",
			SwitchIpAddress:                  "172.16.254.1",
			SwitchInstallDate:                &currentTime,
			SwitchNotes:                      "This is switch dummy data",
			SwitchImageLinks:                 "https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.metapoint.in%2Fassets%2Fupload_images%2Fproduct%2F1629888737_2.png&tbnid=AB2GLl3giplQSM&vet=12ahUKEwicm7HKscL_AhXWNrcAHToJB4EQMygAegUIARDfAQ..i&imgrefurl=https%3A%2F%2Fwww.metapoint.in%2Fproduct%2Fcisco-sg550x-48-k9-eu&docid=wkDrxBuUDSU-NM&w=900&h=600&q=cisco%20switch&client=ubuntu-sn&ved=2ahUKEwicm7HKscL_AhXWNrcAHToJB4EQMygAegUIARDfAQ,https://www.google.com/imgres?imgurl=https%3A%2F%2Fm.media-amazon.com%2Fimages%2FI%2F713vsWHSnML.jpg&tbnid=Aur9U4liYKcXfM&vet=12ahUKEwicm7HKscL_AhXWNrcAHToJB4EQ94IIKAJ6BQgBEOQB..i&imgrefurl=https%3A%2F%2Fwww.amazon.in%2FCisco-Sg350-28P-28-Port-Gigabit-SG35028PK9NA%2Fdp%2FB01HYA397Y&docid=r5Cn99emQ4GdQM&w=2560&h=1800&q=cisco%20switch&client=ubuntu-sn&ved=2ahUKEwicm7HKscL_AhXWNrcAHToJB4EQ94IIKAJ6BQgBEOQB,https://www.google.com/imgres?imgurl=https%3A%2F%2Fwww.ycict.net%2Fwp-content%2Fuploads%2Fsites%2F5%2F2019%2F04%2FCisco-Catalyst-2960-X-Series-Switches-3.jpg&tbnid=wr_Nqc3wFHXXkM&vet=12ahUKEwicm7HKscL_AhXWNrcAHToJB4EQMygHegUIARDvAQ..i&imgrefurl=https%3A%2F%2Fwww.ycict.net%2Fmrj%2Fproducts%2Fcisco-catalyst-2960-x-series-switches%2F&docid=s14MAq8N117GDM&w=600&h=600&q=cisco%20switch&client=ubuntu-sn&ved=2ahUKEwicm7HKscL_AhXWNrcAHToJB4EQMygHegUIARDvAQ",
			WirelessUnit:                     "Ubiquity Wireless AP",
			WirelessIpAddress:                "192.158.1.38",
			WirelessAdminUsername:            "nms_admin",
			WirelessAdminPassword:            "admin@123",
			WirelessSsid:                     "TDST2200H0156",
			WirelessPassword:                 "password",
			WirelessConnectionType:           "LAN",
			WirelessNotes:                    "hello notes",
			OnSiteBackupServerOrNasType:      "Differential",
			OnSiteBackupServerOrNasIpAddress: "192.158.1.38",
			ManagementServer:                 "RDP: ada.nmsknows.com",
			ManagementNotes:                  "Hello management",
			ManagementIpAddress:              "192.158.1.38",
			ScanToFolderLocation:             `\\amofs\JDrive\Scans`,
			ScanToFolderUsernameOrPassword:   `AMO\scanuser - !AMOsc@nuser`,
			ScanToEmailSmtpServerPort:        "smtp.gmail.com",
			ScanToEmailEmailOrPassword:       "U:Communitycoalition207@gmail.com P: 8256712Ed!",
			PrintOrScannerType:               "Optical",
			PrintOrScannerUserName:           "admin",
			PrintOrScannerPassword:           "123456",
			PrintOrScannerIpAddress:          "192.158.1.38",
			CreatedAt:                        &currentTime,
			CreatedByUserID:                  1,
		})

	initializers.DB().Create(
		&models.Software{
			CustomerLocationID: 1,
			Name:               "Quickbooks Enterprise",
			Version:            "23.0",
			LicenseKey:         "8644 5606 5564 491",
			ServerOrVM:         "BBCD1 / BBDB1",
			OtherLicenseInfo:   "Product Key 885197",
			InstallDate:        &currentTime,
			ExpiryDate:         &currentTime,
			Notes:              "Installed to replace buggy QB Ent. 22",
			CreatedAt:          &currentTime,
			CreatedByUserID:    1,
		})

	initializers.DB().Create(
		&models.EmailAccountType{
			EmailAccountTypeID: 1,
			EmailAccountType:   "Exchange",
		})

	initializers.DB().Create(
		&models.EmailAccountType{
			EmailAccountTypeID: 2,
			EmailAccountType:   "pop/imap",
		})

	initializers.DB().Create(
		&models.EmailAccountType{
			EmailAccountTypeID: 3,
			EmailAccountType:   "Hybrid (Exchange/POP)",
		})

	initializers.DB().Create(
		&models.EmailAccountType{
			EmailAccountTypeID: 4,
			EmailAccountType:   "Open Exchange",
		})

	initializers.DB().Create(
		&models.EmailAccount{
			EmailHosting:       "Godaddy",
			LinkForEmailAdmin:  "admin@godaddy.com",
			AccountNumber:      "2536789",
			EmailAccountTypeID: 1,
			UserName:           "nms",
			Password:           "admin@123",
			Pin:                "7878",
		})

	initializers.DB().Create(
		&models.EmailDomain{
			CustomerLocationID: 1,
			DomainRegistrar:    "Godaddy",
			AccountID:          1,
			Password:           "jhxnyn",
			Pin:                "7878",
			Domain:             "nms",
			ARecord1:           "temp1@nms.com",
			ARecord2:           "temp2@nms.com",
			ARecord3:           "temp3@nms.com",
			ARecord4:           "temp4@nms.com",
			EmailAccountID:     1,
			MxRecord1:          "temp5@nms.com",
			MxRecord2:          "temp6@nms.com",
			WebsiteIpOrAlias:   "www.nms.com",
			WebMailPopImap:     "IMAP",
			WebMailExchange:    "N/A",
			PhoneSettingNote:   "N/A",
			Notes:              "N/A",
			CreatedAt:          &currentTime,
			CreatedByUserID:    1,
		})

	initializers.DB().Create(
		&models.Server{
			HostName:              "nms-main",
			CustomerLocationID:    1,
			HardwareAsAService:    &boolValue,
			OsPlatform:            "FS OS",
			ServiceTag:            "G09B182",
			ExpressionServiceCode: 34844148722,
			Location:              "Flex",
			Warranty:              "Dell Basic 3 Years",
			Type:                  "S5850-48S6Q",
			PowerConnectType:      "Switch",
			PurchaseDate:          &currentTime,
			ExpirationDate:        &currentTime,
			DaysLeft:              "",
			Ownership:             "nms",
			OrderNumber:           "",
			Description:           "N/A",
			Idrac:                 "",
			CreatedAt:             &currentTime,
			CreatedByUserID:       1,
		})

	initializers.DB().Create(
		&models.InstalledFirewall{
			CustomerLocationID:     1,
			CustomerLanOnSite:      "192.168.52.0/24",
			Brand:                  "OPNSense",
			Equipment:              "Blackbox 4 port Model:FW4C",
			InstalledDate:          &currentTime,
			InternetProvider1ID:    1,
			InternetProvider2ID:    2,
			FirewallIpv4LanAddress: "10.0.21.1",
			FirewallIpv6LanAddress: "2607:f758:5080:21::1",
			CurrentVersion:         "23.1.9",
			VersionBackup:          "23.1",
			BackupDate:             &currentTime,
			Extra:                  "Firewall user. Done",
			CreatedAt:              &currentTime,
			CreatedByUserID:        1,
		})

	initializers.DB().Create(
		&models.Note{
			Subject:            "BINAT before OpenVPN tunnel",
			CustomerLocationID: 1,
			Attachment:         "BINAT_Entry_for_VPN.png ",
			Note:               "Use the procedure in this guide to set up a site-to-site VPN connection with Access Server and a site-to-site connector using an OpenVPN client.",
			AssignedToID:       0,
			CreatedAt:          &currentTime,
			CreatedByUserID:    1,
		})
}
