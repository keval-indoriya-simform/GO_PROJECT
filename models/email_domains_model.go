package models

import (
	"Application/initializers"
	"gorm.io/gorm"
	"time"
)

type EmailDomain struct {
	EmailDomainID      int32             `gorm:"primaryKey;" json:"email_domain_id"`
	CustomerLocationID int32             `gorm:"type:varchar" json:"customer_location_id,omitempty"`
	DomainRegistrar    string            `gorm:"type:varchar" json:"domain_registrar,omitempty"`
	AccountID          int32             `json:"account_id,omitempty"`
	Password           string            `gorm:"type:varchar" json:"password,omitempty"`
	Pin                string            `gorm:"type:varchar" json:"pin,omitempty"`
	Domain             string            `gorm:"type:varchar;not null" json:"domain"`
	ARecord1           string            `gorm:"type:varchar" json:"a_record_1,omitempty"`
	ARecord2           string            `gorm:"type:varchar" json:"a_record_2,omitempty"`
	ARecord3           string            `gorm:"type:varchar" json:"a_record_3,omitempty"`
	ARecord4           string            `gorm:"type:varchar" json:"a_record_4,omitempty"`
	EmailAccountID     int32             `json:"email_account_id,omitempty"`
	MxRecord1          string            `gorm:"type:varchar" json:"mx_record_1,omitempty"`
	MxRecord2          string            `gorm:"type:varchar" json:"mx_record_2,omitempty"`
	WebsiteIpOrAlias   string            `gorm:"type:varchar" json:"website_ip_or_alias,omitempty"`
	WebMailPopImap     string            `gorm:"type:varchar" json:"web_mail_pop_imap,omitempty"`
	WebMailExchange    string            `gorm:"type:varchar" json:"web_mail_exchange,omitempty"`
	PhoneSettingNote   string            `gorm:"type:varchar" json:"phone_setting_note,omitempty"`
	Notes              string            `gorm:"type:varchar" json:"notes,omitempty"`
	CreatedAt          *time.Time        `gorm:"type:timestamp without time zone" json:"created_at,omitempty"`
	CreatedByUserID    int32             `gorm:"not null;" json:"created_by_user_id,omitempty"`
	DeletedAt          *time.Time        `gorm:"type:timestamp without time zone" json:"deleted_at,omitempty"`
	DeletedByUserID    int32             `gorm:"" json:"deleted_by_user_id,omitempty"`
	UpdatedAt          *time.Time        `gorm:"type:timestamp without time zone" json:"updated_at,omitempty"`
	UpdatedByUserID    int32             `gorm:"" json:"updated_by_user_id,omitempty"`
	CustomerLocation   *CustomerLocation `gorm:"references:CustomerLocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Account            *Account          `gorm:"references:AccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"account,omitempty"`
	EmailAccount       *EmailAccount     `gorm:"references:EmailAccountID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"email_account,omitempty"`
}

func (emailDomainModel *EmailDomain) TableName() string {
	return "email_domains"
}

func (emailDomainModel *EmailDomain) CreateEmailDomain() error {
	var (
		emailTemp   *EmailAccount
		createQuery *gorm.DB
	)
	if emailDomainModel.EmailAccount != nil {
		err := emailDomainModel.EmailAccount.CreateEmailAccount()
		if err != nil {
			return err
		}
		emailTemp = emailDomainModel.EmailAccount
		emailDomainModel.EmailAccount = nil
	}

	if emailDomainModel.CustomerLocationID > 0 && emailDomainModel.EmailAccount != nil {
		if emailDomainModel.Account != nil {
			createQuery = initializers.DB().Create(&emailDomainModel)
		} else {
			createQuery = initializers.DB().Omit("AccountID").Create(&emailDomainModel)
		}
	} else {
		if emailDomainModel.CustomerLocationID == 0 && emailDomainModel.EmailAccount != nil {
			if emailDomainModel.Account != nil {
				createQuery = initializers.DB().Omit("CustomerLocationID").Create(&emailDomainModel)
			} else {
				createQuery = initializers.DB().Omit("AccountID,CustomerLocationID").Create(&emailDomainModel)
			}
		} else if emailDomainModel.CustomerLocationID == 0 && emailDomainModel.EmailAccount == nil {
			if emailDomainModel.Account != nil {
				createQuery = initializers.DB().Omit("CustomerLocationID,EmailAccountID").Create(&emailDomainModel)
			} else {
				createQuery = initializers.DB().Omit("AccountID,CustomerLocationID,EmailAccountID").Create(&emailDomainModel)
			}
		} else if emailDomainModel.CustomerLocationID > 0 && emailDomainModel.EmailAccount == nil {
			if emailDomainModel.Account != nil {
				createQuery = initializers.DB().Omit("EmailAccountID").Create(&emailDomainModel)
			} else {
				createQuery = initializers.DB().Omit("AccountID,EmailAccountID").Create(&emailDomainModel)

			}
		}
	}
	if emailDomainModel.EmailDomainID > 0 {
		CreateLog(
			emailDomainModel.CreatedByUserID,
			"Created",
			emailDomainModel.TableName(),
			emailDomainModel.EmailDomainID,
		)
		emailDomainModel.EmailAccount = emailTemp
		return nil
	}
	return createQuery.Error

}

func (emailDomainModel *EmailDomain) UpdateEmailDomain() error {
	if emailDomainModel.GetEmailDomainId(emailDomainModel.EmailDomainID) {
		updateQuery := initializers.DB().Where("deleted_at IS NULL").
			Updates(&emailDomainModel)
		if updateQuery.RowsAffected > 0 {
			CreateLog(
				emailDomainModel.UpdatedByUserID,
				"Updated",
				emailDomainModel.TableName(),
				emailDomainModel.EmailDomainID,
			)
			return nil
		} else {
			return updateQuery.Error
		}
	}
	return NotFoundError
}

func (emailDomainModel *EmailDomain) GetEmailDomainId(emailDomainID int32) bool {
	findQuery := initializers.DB().Model(EmailDomain{}).Select("email_domain_id,account_id,email_account_id").
		Where("email_domain_id = ? AND deleted_at IS NULL", emailDomainID).First(&emailDomainModel)
	emailDomainModel.Account.AccountID = emailDomainModel.AccountID
	emailDomainModel.EmailAccount.EmailAccountID = emailDomainModel.EmailAccountID
	return findQuery.RowsAffected > 0
}

func (emailDomainModel *EmailDomain) DeleteEmailDomain(emailDomainIDs []int32) error {
	deleteQuery := initializers.DB().Where("deleted_at IS NULL AND email_domain_id IN ?", emailDomainIDs).Updates(&emailDomainModel)
	if deleteQuery.RowsAffected > 0 {
		var logs []Log
		for index, _ := range emailDomainIDs {
			logs = append(logs, Log{
				UserID:    emailDomainModel.DeletedByUserID,
				LogType:   "Deleted",
				TableName: emailDomainModel.TableName(),
				DataID:    emailDomainIDs[index],
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

func RetrieveEmailDomainRecords(queries map[string]string,
	pagination func(db *gorm.DB) *gorm.DB) (records []map[string]interface{}, err error) {
	var (
		count                                                  int64
		emailDomainIDQuery, domainQuery, customerLocationQuery string
	)

	if queries["email_domain_id"] != "" {
		emailDomainIDQuery = " AND email_domains.email_domain_id = " + queries["email_domain_id"]
	}

	if queries["domain"] != "" {
		domainQuery = " AND email_domains.domain='" + queries["domain"] + "'"
	}
	if queries["customer_location"] != "" {
		customerLocationQuery = " AND customer_locations.name ILIKE '%" + queries["customer_location"] + "%'"
	}

	retrievedRows := initializers.DB().Scopes(pagination).Model(EmailDomain{}).Joins("left join customer_locations on " +
		"customer_locations.customer_location_id=email_domains.customer_location_id").
		Joins("left join accounts on accounts.account_id=email_domains.account_id").
		Joins("left join email_accounts on email_accounts.email_account_id=email_domains.email_account_id").
		Joins("left join email_account_types " +
			"on email_account_types.email_account_type_id=email_accounts.email_account_type_id").
		Where("email_domains.deleted_at is NULL" + emailDomainIDQuery + domainQuery + customerLocationQuery).
		Select("email_domains.email_domain_id,customer_locations.name as customer_location,email_domains.domain_registrar," +
			"accounts.link_for_domain_admin,accounts.account_number," +
			"accounts.account_number_aliases,accounts.name as account_user_name," +
			"email_domains.password as password,email_domains.pin as pin," +
			"email_domains.domain,email_domains.a_record1,email_domains.a_record2,email_domains.a_record3," +
			"email_domains.a_record4,email_accounts.email_hosting,email_accounts.link_for_email_admin," +
			"email_accounts.account_number as email_account_number,email_account_types.email_account_type," +
			"email_accounts.user_name as email_user_name,email_accounts.password as email_account_password," +
			"email_accounts.pin as email_pin,email_domains.mx_record1," +
			"email_domains.mx_record2,email_domains.website_ip_or_alias,email_domains.web_mail_pop_imap," +
			"email_domains.web_mail_exchange,email_domains.notes,email_account_types.email_account_type_id," +
			"email_domains.phone_setting_note,email_domains.phone_setting_note,accounts.name as account_name").
		Find(&records)
	err = retrievedRows.Error
	count = retrievedRows.RowsAffected
	if count == 0 && err == nil {
		err = NotFoundError
	}
	return
}
