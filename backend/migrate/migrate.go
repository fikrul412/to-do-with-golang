package main

import (
	"backend/initializers"
	"backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Todo{}, &models.Category{})
}
