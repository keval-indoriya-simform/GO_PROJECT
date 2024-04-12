package routers

import (
	"Application/controllers"
	_ "Application/docs"
	"Application/middlewares"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	Route = gin.Default()
	url   = ginSwagger.URL("http://to-do-alb-1758059883.us-east-1.elb.amazonaws.com:8080/docs/doc.json")
)

func init() {

	sessionSecretKey := []byte(strings.ReplaceAll(os.Getenv("SESSION_SECRET_KEY"), `"`, ``))
	sessionStore := cookie.NewStore(sessionSecretKey)

	// SET HTML LOAD PATH
	Route.LoadHTMLGlob("views/*.html")

	// SET STATIC LOAD PATH
	Route.StaticFS("/static", http.Dir("./static"))

	// SET SESSION
	Route.Use(sessions.Sessions("sessionStore", sessionStore))

	// Swagger Path
	Route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	swagger := Route.Group("/docs", middlewares.AuthorizeMiddleware())

	apiGroup := Route.Group("/api/v1")
	// SET VIEW ROUTE PATH
	viewGroup := Route.Group("/network-management-solutions")

	// SET TEMPLATE SAMPLE PATH #TEMPORARY
	sampleGroup := Route.Group("/sample")

	// SAMPLE ROUTES HERE
	sampleGroup.GET("/form", controllers.TemplateFormPageController)
	sampleGroup.GET("/list", controllers.TemplateListPageController)

	// VIEWS ROUTES HERE
	viewGroup.GET("/dashboard", controllers.DashboardController)
	viewGroup.GET("", controllers.LoginPageController)
	viewGroup.GET("login", controllers.LoginPageController)
	viewGroup.GET("register", controllers.RegisterPageController)
	createViewGroup := viewGroup.Group("/create", middlewares.AuthorizeMiddleware())
	listViewGroup := viewGroup.Group("/list", middlewares.AuthorizeMiddleware())
	viewGroup.GET("/logout", middlewares.AuthorizeMiddleware(), controllers.LogoutController)

	// VIEWS ROUTE FOR CREATE HERE
	createViewGroup.GET("/cloud-private-ips", controllers.CreateCloudPrivateIpPageController)
	createViewGroup.GET("/cloud-public-ips", controllers.CreateCloudPublicIpPageController)
	createViewGroup.GET("/customer-locations", controllers.CreateCustomerLocationPageController)
	createViewGroup.GET("/customers", controllers.CreateCustomerPageController)
	createViewGroup.GET("/customer-lans", controllers.CreateCustomerLanPageController)
	createViewGroup.GET("/email-domains", controllers.CreateEmailDomainPageController)
	createViewGroup.GET("/installed-firewalls", controllers.CreateInstalledFirewallPageController)
	createViewGroup.GET("/notes", controllers.CreateNotePageController)
	createViewGroup.GET("/servers", controllers.CreateServerPageController)
	createViewGroup.GET("/softwares", controllers.CreateSoftwarePageController)

	// VIEWS ROUTE FOR LIST HERE
	listViewGroup.GET("/cloud-private-ips", controllers.ListCloudPrivateIpPageController)
	listViewGroup.GET("/cloud-public-ips", controllers.ListCloudPublicIpPageController)
	listViewGroup.GET("/customers", controllers.ListCustomerPageController)
	listViewGroup.GET("/customer-lans", controllers.ListCustomerLanPageController)
	listViewGroup.GET("/customer-locations", controllers.ListCustomerLocationPageController)
	listViewGroup.GET("/email-domains", controllers.ListEmailDomainPageController)
	listViewGroup.GET("/installed-firewalls", controllers.ListInstalledFirewallPageController)
	listViewGroup.GET("/notes", controllers.ListNotePageController)
	listViewGroup.GET("/servers", controllers.ListServerPageController)
	listViewGroup.GET("/softwares", controllers.ListSoftwarePageController)

	// API ROUTES HERE
	apiGroup.GET("", controllers.HealthCheckController)
	customerLocationGroup := apiGroup.Group("/customer-locations", middlewares.AuthorizeMiddleware())
	customerGroup := apiGroup.Group("/customers", middlewares.AuthorizeMiddleware())
	installedFirewallGroup := apiGroup.Group("/installed-firewalls", middlewares.AuthorizeMiddleware())
	notesGroup := apiGroup.Group("/notes", middlewares.AuthorizeMiddleware())
	cloudPublicIpGroup := apiGroup.Group("/cloud-public-ips", middlewares.AuthorizeMiddleware())
	serverGroup := apiGroup.Group("/servers", middlewares.AuthorizeMiddleware())
	userGroup := apiGroup.Group("/users")
	loginGroup := apiGroup.Group("/login")
	emailDomainGroup := apiGroup.Group("/email-domains", middlewares.AuthorizeMiddleware())
	customerLanGroup := apiGroup.Group("/customer-lans", middlewares.AuthorizeMiddleware())
	cloudPrivateIp := apiGroup.Group("/cloud-private-ips", middlewares.AuthorizeMiddleware())
	software := apiGroup.Group("/softwares", middlewares.AuthorizeMiddleware())
	dashboardGroup := apiGroup.Group("/dashboard", middlewares.AuthorizeMiddleware())

	// API ROUTE FOR CUSTOMER LOCATIONS
	customerLocationGroup.POST("", controllers.CreateCustomerLocationController)
	customerLocationGroup.PATCH("", controllers.UpdateCustomerLocationController)
	customerLocationGroup.DELETE("", controllers.DeleteCustomerLocationController)
	customerLocationGroup.GET("", controllers.RetrieveCustomerLocationController)

	// API ROUTE FOR CUSTOMER
	customerGroup.GET("cloud-or-onsites", controllers.RetrieveCloudOrOnsiteController)
	customerGroup.GET("", controllers.RetrieveCustomerController)
	customerGroup.POST("", controllers.CreateCustomerController)
	customerGroup.PATCH("", controllers.UpdateCustomerController)
	customerGroup.DELETE("", controllers.DeleteCustomerController)

	// API ROUTE FOR INSTALLED FIREWALLS
	installedFirewallGroup.POST("", controllers.CreateInstalledFirewallController)
	installedFirewallGroup.PATCH("", controllers.UpdateInstalledFirewallController)
	installedFirewallGroup.DELETE("", controllers.DeleteInstalledFirewallController)
	installedFirewallGroup.GET("", controllers.RetrieveInstalledFirewallController)

	// API ROUTE FOR NOTES
	notesGroup.GET("", controllers.RetrieveNoteController)
	notesGroup.POST("", controllers.CreateNoteController)
	notesGroup.PATCH("", controllers.UpdateNoteController)
	notesGroup.DELETE("", controllers.DeleteNoteController)

	// API ROUTE FOR CLOUD PUBLIC IPS
	cloudPublicIpGroup.POST("", controllers.CreateCloudPublicIPController)
	cloudPublicIpGroup.PATCH("", controllers.UpdateCloudPublicIPController)
	cloudPublicIpGroup.DELETE("", controllers.DeleteCloudPublicIPController)
	cloudPublicIpGroup.GET("", controllers.RetrieveCloudPublicIPController)

	// API ROUTE FOR SERVERS
	serverGroup.POST("", controllers.CreateServerController)
	serverGroup.PATCH("", controllers.UpdateServerController)
	serverGroup.DELETE("", controllers.DeleteServerController)
	serverGroup.GET("", controllers.RetrieveServerController)

	// API ROUTE FOR USERS
	userGroup.GET("", controllers.RetrieveUsersController)
	userGroup.POST("register", controllers.RegisterUserController)
	userGroup.PATCH("", controllers.UpdateUserProfileController)
	userRoleGroup := userGroup.Group("/roles")

	// API ROUTE FOR USER ROLES
	userRoleGroup.POST("", controllers.AddUserRoleController)
	userRoleGroup.PATCH("", controllers.UpdateUserRoleController)

	// API ROUTE FOR EMAIL DOMAINS
	emailDomainGroup.POST("", controllers.CreateEmailDomainController)
	emailDomainGroup.PATCH("", controllers.UpdateEmailDomainController)
	emailDomainGroup.DELETE("", controllers.DeleteEmailDomainController)
	emailDomainGroup.GET("", controllers.RetrieveEmailDomainController)
	emailDomainGroup.GET("email-account-types", controllers.RetrieveEmailAccountTypeController)

	// API ROUTE FOR CUSTOMER LANS
	customerLanGroup.GET("", controllers.RetrieveCustomerLanController)
	customerLanGroup.POST("", controllers.CreateCustomerLanController)
	customerLanGroup.PATCH("", controllers.UpdateCustomerLanController)
	customerLanGroup.DELETE("", controllers.DeleteCustomerLanController)

	// API ROUTE FOR CLOUD PRIVATE IPS
	cloudPrivateIp.POST("", controllers.CreateCloudPrivateIpController)
	cloudPrivateIp.PATCH("", controllers.UpdateCloudPrivateIpController)
	cloudPrivateIp.DELETE("", controllers.DeleteCloudPrivateIpController)
	cloudPrivateIp.GET("", controllers.RetrieveCloudPrivateIpController)

	// API ROUTE FOR SOFTWARES
	software.POST("", controllers.CreateSoftwareController)
	software.PATCH("", controllers.UpdateSoftwareController)
	software.DELETE("", controllers.DeleteSoftwareController)
	software.GET("", controllers.RetrieveSoftwareController)

	// API ROUTE FOR SWAGGER
	swagger.GET("", controllers.SwaggerController)
	// API ROUTE FOR LOGIN
	loginGroup.POST("", controllers.UserLoginController)

	dashboardGroup.GET("", controllers.RetrieveDashboardCount)
}
