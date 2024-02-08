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

func (s *HairdresserStore) GetHairdressersBySalonId(salonId int64) ([]*coifResa.HairdresserItem, error) {
	rows, err := s.Query(`
    SELECT id, name, salon_id FROM haidressers WHERE salon_id = $1
    `, salonId)
	if err != nil {
		return nil, fmt.Errorf("failed to get haidresser with salon id %d: %w", salonId, err)
	}
	defer rows.Close()

	var hairdressers []*coifResa.HairdresserItem
	for rows.Next() {
		hairdresser := &coifResa.HairdresserItem{}
		err := rows.Scan(&hairdresser.ID, &hairdresser.Name, &hairdresser.SalonId)
		if err != nil {
			return nil, err
		}
		hairdressers = append(hairdressers, hairdresser)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return hairdressers, nil
}

func (s *HairdresserStore) UpdateHairdresser(hairdresser *coifResa.HairdresserItem) error {
	_, err := s.Exec(`
	UPDATE haidressers SET name = $1, salon_id = $2 WHERE id = $3
	`, hairdresser.Name, hairdresser.SalonId, hairdresser.ID)
	if err != nil {
		return fmt.Errorf("failed to update haidresser with id %d: %w", hairdresser.ID, err)
	}

	return nil
}

func (s *HairdresserStore) DeleteHairdresser(id int64) error {
	_, err := s.Exec(`
	DELETE FROM haidressers WHERE id = $1
	`, id)
	if err != nil {
		return fmt.Errorf("failed to delete haidresser with id %d: %w", id, err)
	}

	return nil
}
