package web

import (
	"coifResa"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateSlot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slot := &coifResa.SlotItem{}

		err := json.NewDecoder(r.Body).Decode(slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.Store.CreateSlot(slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(struct {
			Status string             `json:"status"`
			Slot   *coifResa.SlotItem `json:"slot"`
		}{
			Status: "success",
			Slot:   slot,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
