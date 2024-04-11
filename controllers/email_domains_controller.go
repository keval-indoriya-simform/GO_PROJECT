package controllers

import (
	"Application/models"
	"Application/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateEmailDomainController
// Description To create email domain
// Create Email Domain godoc
//
//	@Summary		Create Email Domain
//	@Description	Add    Email Domain
//	@Tags			Email Domain
//	@Accept			json
//	@Produce		json
//	@Param			EmailDomain	body		models.CreateEmailDomain{account=models.CreateAccount,email_account=models.CreateEmailAccount}	true	"EmailDomain Data"
//	@Success		200			{object}	models.EmailDomain
//	@Router			/email-domains [post]
func CreateEmailDomainController(context *gin.Context) {

	var emailDomainModel models.EmailDomain

	bindJSONError := context.ShouldBindJSON(&emailDomainModel)
	fmt.Println(emailDomainModel)
	context.JSON(

		services.CreateServices(emailDomainModel, "EmailDomain", bindJSONError),
	)

}

// UpdateEmailDomainController
// Description To update email domain
//
//	@Summary		Update Email Domain
//	@Description	Update Email Domain
//	@Tags			Email Domain
//	@Accept			json
//	@Produce		json
//	@Param			EmailDomain		body		models.UpdateEmailDomain{account=models.UpdateAccount,email_account=models.UpdateEmailAccount}	true	"Email Domain"
//	@Param			email_domain_id	query		string			false	"Email Domain ID"
//	@Success		200				{object}	models.Customer
//	@Router			/email-domains [patch]
func UpdateEmailDomainController(context *gin.Context) {
	var emailDomainModel models.EmailDomain
	emailDomainID := context.Query("email_domain_id")
	bindJSONError := context.ShouldBindJSON(&emailDomainModel)
	context.JSON(

		services.UpdateServices(emailDomainID, emailDomainModel, "EmailDomain", bindJSONError),
	)

}

// DeleteEmailDomainController
// Description To delete email domain
//
//	@Summary		Delete Email Domain
//	@Description	Delete Email Domain
//	@Tags			Email Domain
//	@Produce		json
//	@Param			email_domain_id	query		string	false	"Email Domain Id"
//	@Param			USER_ID			header		string	false	"USER ID"
//	@Success		200				{object}	models.EmailDomain
//	@Router			/email-domains [delete]
func DeleteEmailDomainController(context *gin.Context) {

	emailDomainID := context.Query("email_domain_id")
	userID := context.Request.Header.Get("USER_ID")

	context.JSON(

		services.DeleteServices(emailDomainID, userID, "EmailDomain"),
	)

}

// RetrieveEmailDomainController
// Description To get email domain
//
//	@Summary		Retrieve   Email Domain
//	@Description	Retrieve Email Domain
//	@Tags			Email Domain
//	@Produce		json
//	@Param			domain				query		string	false	"Domain"
//	@Param			email_domain_id		query		string	false	"Email Domain Id"
//	@Param			customer_location	query		string	false	"Customer Location"
//	@Param			page				query		int		false	"Page"
//	@Param			order_by			query		string	false	"Order By"
//
//	@Success		200					{object}	models.EmailDomain
//	@Router			/email-domains [get]
func RetrieveEmailDomainController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "EmailDomain"
	queries["model_id"] = "email_domain_id"
	queries["email_domain_id"] = context.Query("email_domain_id")
	queries["domain"] = context.Query("domain")
	queries["customer_location"] = context.Query("customer_location")
	queries["order_by"] = context.Query("order_by")
	queries["page"] = context.Query("page")

	context.JSON(

		services.RetrieveServices(queries),
	)
}

func RetrieveEmailAccountTypeController(context *gin.Context) {
	queries := map[string]string{}
	queries["model_name"] = "EmailAccountType"
	queries["model_id"] = "email_account_type_id"
	queries["set_limit"] = context.Query("set_limit")
	context.JSON(

		services.RetrieveServices(queries),
	)

}
