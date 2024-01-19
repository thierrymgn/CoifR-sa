package database

import (
	"coifResa"
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
	}
}

type Store struct {
	coifResa.UserStoreInterface
}