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
		INSERT INTO reservations (user_id, slot_id) VALUES ($1, $2) RETURNING id
	`, reservation.UserId, reservation.SlotId).Scan(&reservation.ID)

	if err != nil {
		return fmt.Errorf("failed to create reservation: %w", err)
	}

	return nil
}

func (s *ReservationStore) GetReservation(id int64) (*coifResa.ReservationItem, error) {
	reservation := &coifResa.ReservationItem{}

	err := s.QueryRow(`
        SELECT r.id, r.user_id, r.slot_id, s.start_time, s.end_time, s.hairdresser_id
        FROM reservations r
        INNER JOIN slots s ON r.slot_id = s.id
        WHERE r.id = $1
    `, id).Scan(&reservation.ID, &reservation.UserId, &reservation.SlotId, &reservation.Slot.StartTime, &reservation.Slot.EndTime, &reservation.Slot.HairdresserId)

	if err != nil {
		return nil, fmt.Errorf("failed to get reservation with id %d: %w", id, err)
	}

	return reservation, nil
}

func (s *ReservationStore) GetReservationsByUserId(userId int64) ([]*coifResa.ReservationItem, error) {
	rows, err := s.Query(`
	SELECT id, user_id, slot_id FROM reservations WHERE user_id = $1
	`, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get reservation with user id %d: %w", userId, err)
	}
	defer rows.Close()

	var reservations []*coifResa.ReservationItem
	for rows.Next() {
		reservation := &coifResa.ReservationItem{}
		err := rows.Scan(&reservation.ID, &reservation.UserId, &reservation.SlotId)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (s *ReservationStore) DeleteReservation(id int64) error {
	_, err := s.Exec(`
	DELETE FROM reservations WHERE id = $1
	`, id)

	if err != nil {
		return fmt.Errorf("failed to delete reservation with id %d: %w", id, err)
	}

	return nil
}
