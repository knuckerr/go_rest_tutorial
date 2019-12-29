package controllers

import (
	"github.com/knuckerr/go_rest/api/middlewares"
)

func (server *Server) InitializeRoutes() {
	// PUBLIC SERVICES
	v1 := server.Router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", server.Login)
		}
	}
	// AUTH REQUIRE
	auth_v1 := server.Router.Group("/v1")
	auth_v1.Use(middlewares.AuthenticationRequired())
	{
		auth := auth_v1.Group("/auth")
		{
			auth.POST("/refresh", server.RefreshToken)
		}
		users := auth_v1.Group("/users")
		{
			users.GET("/", server.GetUsers)
			users.PUT("/:id", server.UpdateUser)
			users.GET("/:id", server.GetUser)
			users.POST("/", server.CreateUser)
			users.DELETE("/:id", server.DeleteUser)
		}
	}

}
