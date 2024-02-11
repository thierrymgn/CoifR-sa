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

func (s *SlotStore) GetSlotsByHairdresserId(hairdresserId int64) ([]*coifResa.SlotItem, error) {
	rows, err := s.Query(`
	SELECT id, start_time, end_time, hairdresser_id FROM slots WHERE hairdresser_id = $1
	`, hairdresserId)
	if err != nil {
		return nil, fmt.Errorf("failed to get slot with hairdresser id %d: %w", hairdresserId, err)
	}
	defer rows.Close()

	var slots []*coifResa.SlotItem
	for rows.Next() {
		slot := &coifResa.SlotItem{}
		err := rows.Scan(&slot.ID, &slot.StartTime, &slot.EndTime, &slot.HairdresserId)
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}

	return slots, nil
}

func (s *SlotStore) UpdateSlot(slot *coifResa.SlotItem) error {
	_, err := s.Exec(`
	UPDATE slots SET start_time = $1, end_time = $2 WHERE id = $3
	`, slot.StartTime, slot.EndTime, slot.ID)

	if err != nil {
		return fmt.Errorf("failed to update slot with id %d: %w", slot.ID, err)
	}

	return nil
}

func (s *SlotStore) DeleteSlot(id int64) error {
	_, err := s.Exec(`
	DELETE FROM slots WHERE id = $1
	`, id)

	if err != nil {
		return fmt.Errorf("failed to delete slot with id %d: %w", id, err)
	}

	return nil
}
