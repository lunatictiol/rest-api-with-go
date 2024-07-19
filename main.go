package main

import (
	"github.com/gin-gonic/gin"
	"lunatictiol.com/resApi/db"
	"lunatictiol.com/resApi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.ConfigRouting(server)
	server.Run(":8080")
}
