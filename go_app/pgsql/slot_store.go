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
