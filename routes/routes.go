package routes

import (
	"github.com/gin-gonic/gin"
	"lunatictiol.com/resApi/routes/middlewares"
)

func ConfigRouting(server *gin.Engine) {
	server.GET("/hello", getHello)

	//events

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Aunthenticate)
	authenticated.GET("/events", getEvents)
	authenticated.POST("/events", addEvent)
	authenticated.GET("/event/:id", getEventById)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

	//user
	server.POST("/signup", signup)
	server.POST("/login", login)

}
