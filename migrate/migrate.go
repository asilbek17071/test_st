package main

import (
	"fmt"
	"log"

	"github.com/asilbek17071/test_st/initializers"
	"github.com/asilbek17071/test_st/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Phone{})
	fmt.Println("👍 Migration complete")
}
