package controllers

import (
	"github.com/go-chi/chi"
	"github.com/knuckerr/go_rest/api/middlewares"
)

func (server *Server) InitializeRoutes() {
	// PUBLIC SERVICES
	server.Router.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.With(middlewares.AuthenticationRequired).Get("/", server.GetUsers)
			r.Post("/", server.CreateUser)
			r.Route("/{id}", func(r chi.Router) {
				r.Use(middlewares.AuthenticationRequired)
				r.Get("/", server.GetUser)
				r.Put("/", server.UpdateUser)
				r.Delete("/", server.DeleteUser)
			})
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", server.Login)
			r.With(middlewares.AuthenticationRequired).Post("/refresh", server.RefreshToken)
		})
	})
}
