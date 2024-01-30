package database

import (
	"coifResa"
	"database/sql"
	"fmt"
)

func NewSalonStore(db *sql.DB) *SalonStore {
	return &SalonStore{
		db,
	}
}

type SalonStore struct {
	*sql.DB
}

func (s *SalonStore) CreateSalon(salon *coifResa.SalonItem) error {
	err := s.QueryRow(`
		INSERT INTO salons (name, email, address, city, cp, description, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`, salon.Name, salon.Email, salon.Address, salon.City, salon.PosalCode, salon.Description, salon.UserId).Scan(&salon.ID)

	if err != nil {
		return fmt.Errorf("failed to create salon: %w", err)
	}

	return nil
}
