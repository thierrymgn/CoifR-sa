package database

import (
	"coifResa"
	"database/sql"
	"fmt"
)

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

type UserStore struct {
	*sql.DB
}

func (s *UserStore) CreateUser(user *coifResa.UserItem) error {
	err := s.QueryRow(`
        INSERT INTO users (username, password, email, user_type) VALUES ($1, $2, $3, $4) RETURNING id
    `, user.Username, user.Password, user.Email, user.UserType).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserStore) GetUser(id int64) (*coifResa.UserItem, error) {
	user := &coifResa.UserItem{}

	err := s.QueryRow(`
        SELECT id, username, password, email, user_type FROM users WHERE id = $1
    `, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.UserType)

	if err != nil {
		return nil, fmt.Errorf("failed to get user with id %d: %w", id, err)
	}

	return user, nil
}

func (s *UserStore) GetUserBy(field, value string) (*coifResa.UserItem, error) {
	user := &coifResa.UserItem{}

	query := fmt.Sprintf(`
        SELECT id, username, password, email, user_type FROM users WHERE %s = $1
    `, field)

	err := s.QueryRow(query, value).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.UserType)

	if err != nil {
		return nil, fmt.Errorf("unable to get user by %s: %w", field, err)
	}

	return user, nil
}

func (s *UserStore) GetUserByUsername(username string) (*coifResa.UserItem, error) {
	return s.GetUserBy("username", username)
}

func (s *UserStore) GetUserByEmail(email string) (*coifResa.UserItem, error) {
	return s.GetUserBy("email", email)
}

func (s *UserStore) UpdateUser(user *coifResa.UserItem) error {
	_, err := s.Exec(`
        UPDATE users SET username = $1, password = $2, email = $3, user_type = $4 WHERE id = $5
    `, user.Username, user.Password, user.Email, user.UserType, user.ID)

	if err != nil {
		return fmt.Errorf("failed to update user with id %d: %w", user.ID, err)
	}

	return nil
}

func (s *UserStore) DeleteUser(id int64) error {
	_, err := s.Exec(`
        DELETE FROM users WHERE id = $1
    `, id)

	if err != nil {
		return fmt.Errorf("failed to delete user with id %d: %w", id, err)
	}

	return nil
}
