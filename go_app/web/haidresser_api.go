package web

import (
	"coifResa"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CreateHairdresser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hairdresser := &coifResa.HairdresserItem{}

		err := json.NewDecoder(r.Body).Decode(hairdresser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.Store.CreateHairdresser(hairdresser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(struct {
			Status      string                    `json:"status"`
			Message     string                    `json:"message"`
			Hairdresser *coifResa.HairdresserItem `json:"hairdresser"`
		}{
			Status:      "success",
			Message:     "Coiffeur créé avec succès",
			Hairdresser: hairdresser,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetHairdresser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		hairdresser, err := h.Store.GetHairdresser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(hairdresser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
