package database

import (
	"coifResa"
	"database/sql"
	"fmt"
)

func NewHaidresserStore(db *sql.DB) *HairdresserStore {
	return &HairdresserStore{
		db,
	}
}

type HairdresserStore struct {
	*sql.DB
}

func (s *HairdresserStore) CreateHairdresser(hairdresser *coifResa.HairdresserItem) error {
	err := s.QueryRow(`
	INSERT INTO haidressers (name, salon_id) VALUES ($1, $2) RETURNING id
	`, hairdresser.Name, hairdresser.SalonId).Scan(&hairdresser.ID)

	if err != nil {
		return fmt.Errorf("failed to create salon: %w", err)
	}

	return nil
}
