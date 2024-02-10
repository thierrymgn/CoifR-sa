package database

import (
	"coifResa"
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewUserStore(db),
		NewSalonStore(db),
		NewHaidresserStore(db),
		NewSlotStore(db),
		NewReservationStore(db),
	}
}

type Store struct {
	coifResa.UserStoreInterface
	coifResa.SalonStoreInterface
	coifResa.HairdresserStoreInterface
	coifResa.SlotStoreInterface
	coifResa.ReservationStoreInterface
}
