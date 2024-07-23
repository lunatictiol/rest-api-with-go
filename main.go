package main

import (
	"github.com/gin-gonic/gin"
	"github.comlunatictiol/rest-api-with-go/db"
	"github.comlunatictiol/rest-api-with-go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.ConfigRouting(server)
	server.Run(":8080")

	
}
