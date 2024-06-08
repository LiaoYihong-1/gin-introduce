package main

import (
	"gin/database"
	"gin/routes"
	"github.com/gin-gonic/gin"
)

func startServer() {
	database.InitDatabase()
	defer database.CloseDB()
	ginServer := gin.Default()
	routes.Routes(ginServer)
	ginServer.Run(":8080")
}

func main() {
	startServer()
}
