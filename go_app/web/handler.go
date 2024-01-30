package web

import (
	database "coifResa/pgsql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store *database.Store) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)
	handler.Get("/users/{id}", handler.GetUser())
	handler.Post("/users", handler.CreateUser())
	handler.Put("/users/{id}", handler.UpdateUser())
	handler.Delete("/users/{id}", handler.DeleteUser())
	handler.Get("/users/username/{username}", handler.GetUserByUsername())
	handler.Get("/users/email/{email}", handler.GetUserByEmail())
	handler.Post("/salon", handler.CreateSalon())

	return handler

}

type Handler struct {
	*chi.Mux
	*database.Store
}
