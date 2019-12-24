package controllers

func (server *Server) InitializeRoutes() {
	v1 := server.Router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", server.GetUsers)
			users.PUT("/:id", server.UpdateUser)
			users.GET("/:id", server.GetUser)
			users.POST("/", server.CreateUser)
			users.DELETE("/:id", server.DeleteUser)
		}
	}
}
