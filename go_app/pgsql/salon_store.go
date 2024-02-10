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
		INSERT INTO salons (name, email, address, city, postal_code, description, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`, salon.Name, salon.Email, salon.Address, salon.City, salon.PostalCode, salon.Description, salon.UserId).Scan(&salon.ID)

	if err != nil {
		return fmt.Errorf("failed to create salon: %w", err)
	}

	return nil
}

func (s *SalonStore) GetSalon(id int64) (*coifResa.SalonItem, error) {
	salon := &coifResa.SalonItem{}

	err := s.QueryRow(`
		SELECT id, name, email, address, city, postal_code, description, user_id FROM salons WHERE id = $1
	`, id).Scan(&salon.ID, &salon.Name, &salon.Email, &salon.Address, &salon.City, &salon.PostalCode, &salon.Description, &salon.UserId)

	if err != nil {
		return nil, fmt.Errorf("failed to get salon with id %d: %w", id, err)
	}

	return salon, nil
}

func (s *SalonStore) GetSalonsByUserId(userId int64) ([]*coifResa.SalonItem, error) {
	rows, err := s.Query(`
	SELECT id, name, email, address, city, postal_code, description, user_id FROM salons WHERE user_id = $1
	`, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get salon with user id %d: %w", userId, err)
	}
	defer rows.Close()

	var salons []*coifResa.SalonItem
	for rows.Next() {
		salon := &coifResa.SalonItem{}
		err := rows.Scan(&salon.ID, &salon.Name, &salon.Email, &salon.Address, &salon.City, &salon.PostalCode, &salon.Description, &salon.UserId)
		if err != nil {
			return nil, err
		}
		salons = append(salons, salon)
	}

	return salons, nil
}
