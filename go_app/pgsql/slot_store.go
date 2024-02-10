package database

import (
	"coifResa"
	"database/sql"
	"fmt"
)

func NewSlotStore(db *sql.DB) *SlotStore {
	return &SlotStore{
		db,
	}
}

type SlotStore struct {
	*sql.DB
}

func (s *SlotStore) CreateSlot(slot *coifResa.SlotItem) error {
	err := s.QueryRow(`
	INSERT INTO slots (start_time, end_time, hairdresser_id) VALUES ($1, $2, $3) RETURNING id
	`, slot.StartTime, slot.EndTime, slot.HairdresserId).Scan(&slot.ID)

	if err != nil {
		return fmt.Errorf("failed to create slot: %w", err)
	}

	return nil
}

func (s *SlotStore) GetSlot(id int64) (*coifResa.SlotItem, error) {
	slot := &coifResa.SlotItem{}

	err := s.QueryRow(`
	SELECT id, start_time, end_time, hairdresser_id FROM slots WHERE id = $1
	`, id).Scan(&slot.ID, &slot.StartTime, &slot.EndTime, &slot.HairdresserId)

	if err != nil {
		return nil, fmt.Errorf("failed to get slot with id %d: %w", id, err)
	}

	return slot, nil
}
