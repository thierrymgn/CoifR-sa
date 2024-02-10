package database

import (
	"coifResa"
	"database/sql"
	"fmt"
)

func NewReservationStore(db *sql.DB) *ReservationStore {
	return &ReservationStore{
		db,
	}
}

type ReservationStore struct {
	*sql.DB
}

func (s *ReservationStore) CreateReservation(reservation *coifResa.ReservationItem) error {
	err := s.QueryRow(`
		INSERT INTO reservations (user_id, slot_id) VALUES ($1, $2, $3, $4) RETURNING id
	`, reservation.UserId, reservation.SlotId).Scan(&reservation.ID)

	if err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	return nil
}
