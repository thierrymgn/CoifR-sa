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
	/* USER */
	handler.Get("/users/{id}", handler.GetUser())
	handler.Post("/users", handler.CreateUser())
	handler.Put("/users/{id}", handler.UpdateUser())
	handler.Delete("/users/{id}", handler.DeleteUser())
	handler.Get("/users/username/{username}", handler.GetUserByUsername())
	handler.Get("/users/email/{email}", handler.GetUserByEmail())
	/* SALON */
	handler.Post("/salons", handler.CreateSalon())
	/* HAIRDRESSER */
	handler.Post("/hairdressers", handler.CreateHairdresser())
	handler.Get("/hairdressers/{id}", handler.GetHairdresser())
	handler.Get("/hairdressers/salon/{salonId}", handler.GetHairdressersBySalonId())
	handler.Put("/hairdressers/{id}", handler.UpdateHairdresser())
	handler.Delete("/hairdressers/{id}", handler.DeleteHairdresser())
	/* SLOT */
	handler.Post("/slots", handler.CreateSlot())
	handler.Get("/slots/{id}", handler.GetSlot())
	handler.Get("/slots/hairdresser/{hairdresserId}", handler.GetSlotsByHairdresserId())
	handler.Put("/slots/{id}", handler.UpdateSlot())

	return handler

}

type Handler struct {
	*chi.Mux
	*database.Store
}
