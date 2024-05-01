package main

import (
	"Application/models"
	"Application/routers"
	"log"
	"os"
	"strconv"
	"strings"
)

func init() {

	// SET PAGINATION LIMIT

	models.PageLimit, _ = strconv.Atoi(strings.ReplaceAll(os.Getenv("PAGE_LIMIT"), `"`, ``))
	models.OrderBy = strings.ReplaceAll(os.Getenv("ORDER_BY"), `"`, ``)
}

//	@title			Network Management Solutions API
//	@version		1.0
//	@description	This is a sample server .
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// licence.name Apache 2.0
// licence.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host		192.168.49.2:31471
// @BasePath	/api/v1
// @schemes	http
func main() {

	// START SERVER
	if serverRunError := routers.Route.Run(":" + strings.ReplaceAll(os.Getenv("SERVER_RUN_PORT"), `"`, ``)); serverRunError != nil {
		log.Fatal(serverRunError)
	}

}
