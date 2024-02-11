package web

import (
	"coifResa"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (h *Handler) GetSlot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		slot, err := h.Store.GetSlot(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetSlotsByHairdresserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "hairdresserId"), 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		slots, err := h.Store.GetSlotsByHairdresserId(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(slots)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) UpdateSlot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slotId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

		slot := &coifResa.SlotItem{ID: slotId}

		err = json.NewDecoder(r.Body).Decode(slot)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.Store.UpdateSlot(slot)
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

func (h *Handler) DeleteSlot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		err = h.Store.DeleteSlot(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "success",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
