package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCloudPublicIpPageController
// Description CloudPublicIP create page
func CreateCloudPublicIpPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_cloud_public_ips.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// CreateCloudPrivateIpPageController
// Description CloudPrivateIp create page
func CreateCloudPrivateIpPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_cloud_private_ip.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// CreateCustomerLanPageController
// Description Customer Lan create page
func CreateCustomerLanPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_customer_lan.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// CreateCustomerLocationPageController
// Description Customer Location create page
func CreateCustomerLocationPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_customer_locations.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// CreateCustomerPageController
// Description Customer create page
func CreateCustomerPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_customer.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// CreateEmailDomainPageController
// Description EmailDomain create page
func CreateEmailDomainPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_email_domain.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// CreateInstalledFirewallPageController
// Description Installed Firewall create page
func CreateInstalledFirewallPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_installed_firewalls.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// CreateNotePageController
// Description Note create page
func CreateNotePageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_note.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// CreateServerPageController
// Description Server create page
func CreateServerPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_servers.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// CreateSoftwarePageController
// Description Software create page
func CreateSoftwarePageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"create_software.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

func DashboardController(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"dashboard.html",
		nil,
	)
}

func DefaultPageController(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"default_page.html",
		nil,
	)
}

// ListCloudPublicIpPageController
// Description CloudPublicIP list page
func ListCloudPublicIpPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_cloud_public_ips.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// ListCloudPrivateIpPageController
// Description CloudPrivateIp list page
func ListCloudPrivateIpPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_cloud_private_ip.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// ListCustomerLanPageController
// Description Customer Lan list page
func ListCustomerLanPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_customer_lan.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// ListCustomerLocationPageController
// Description Customer Location list page
func ListCustomerLocationPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_customer_locations.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// ListCustomerPageController
// Description Customer list page
func ListCustomerPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_customer.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// ListEmailDomainPageController
// Description EmailDomain list page
func ListEmailDomainPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_email_domain.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// ListInstalledFirewallPageController
// Description Installed Firewall list page
func ListInstalledFirewallPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_installed_firewalls.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// ListNotePageController
// Description Note list page
func ListNotePageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_note.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// ListServerPageController
// Description Server list page
func ListServerPageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_servers.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		})
}

// ListSoftwarePageController
// Description Software list page
func ListSoftwarePageController(context *gin.Context) {
	userID, _ := context.Get("user_id")
	email, _ := context.Get("email")
	name, _ := context.Get("name")
	role, _ := context.Get("role")
	context.HTML(
		http.StatusOK,
		"list_software.html",
		gin.H{
			"user_id": userID,
			"email":   email,
			"name":    name,
			"role":    role,
		},
	)
}

// LoginPageController
// Description Login page
func LoginPageController(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"login_page.html",
		nil,
	)
}

// RegisterPageController
// Description Register page
func RegisterPageController(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"register_page.html",
		nil,
	)
}

func TemplateFormPageController(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"main.html",
		nil,
	)
}

func TemplateListPageController(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"list_sample.html",
		nil,
	)
}
