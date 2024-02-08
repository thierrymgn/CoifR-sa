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

func (s *HairdresserStore) GetHairdresser(id int64) (*coifResa.HairdresserItem, error) {
	hairdresser := &coifResa.HairdresserItem{}

	err := s.QueryRow(`
	SELECT id, name, salon_id FROM haidressers WHERE id = $1
	`, id).Scan(&hairdresser.ID, &hairdresser.Name, &hairdresser.SalonId)

	if err != nil {
		return nil, fmt.Errorf("failed to get haidresser with id %d: %w", id, err)
	}

	return hairdresser, nil
}
