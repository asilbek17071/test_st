package main

import (
	"log"

	"github.com/asilbek17071/test_st/controllers"
	"github.com/asilbek17071/test_st/initializers"
	"github.com/asilbek17071/test_st/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PhoneController      controllers.PhoneController
	PhoneRouteController routes.PhoneRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PhoneController = controllers.NewPhoneController(initializers.DB)
	PhoneRouteController = routes.NewRoutePhoneController(PhoneController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/user")
	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PhoneRouteController.PhoneRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
