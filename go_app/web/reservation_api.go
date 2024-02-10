package web

import (
	"coifResa"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateReservation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reservation := &coifResa.ReservationItem{}

		err := json.NewDecoder(r.Body).Decode(reservation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.Store.CreateReservation(reservation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(struct {
			Status      string                    `json:"status"`
			Reservation *coifResa.ReservationItem `json:"reservation"`
		}{
			Status:      "success",
			Reservation: reservation,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
