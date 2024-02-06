package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)
func main() {
	db.InitDB()
	server := gin.Default()
	api := server.Group("/api/v1")

	api.GET("/events", routes.GetEvents)
	api.GET("/events/:id", routes.GetEvent)

	authenticated := api.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", routes.CreateEvent)
	authenticated.PUT("/events/:id", routes.UpdateEvent)
	authenticated.DELETE("/events/:id", routes.DeleteEvent)
	authenticated.POST("/events/:id/register", routes.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", routes.CancelRegistration)

	api.POST("/signup", routes.Signup)
	api.POST("/login", routes.Login)

	server.Run(":3000")
}
