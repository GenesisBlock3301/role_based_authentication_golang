package main

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/configurations"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/configurations/db"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/routes"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/schemas"
	_ "github.com/GenesisBlock3301/role_based_access_boilerplate_go/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
)

func init() {
	initEnv()
	db.ConnectionWithDB()
}

func initEnv() {
	log.Println("Loading env setting....")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No local env file. Using global OS environment variables")
	}
	configurations.SetEnvVariable()
	schemas.SetTableName()
}

// @title           Role Based Authentication API
// @version         1.0
// @description     This is golang full-featured project.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Initialize `gin` router
	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.HandleMethodNotAllowed = true
	routes.RootRouter(router)
	router.Run(":8081")
}
