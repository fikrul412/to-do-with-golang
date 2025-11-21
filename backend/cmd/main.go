package main

import (
	"github.com/gin-gonic/gin"
	"backend/initializers"
	"backend/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
}