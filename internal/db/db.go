package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	_, err = db.Exec(`
	CREATE TABLE users(
		id BIGINT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		user_type VARCHAR(255) CHECK (user_type IN('client', 'salon', 'admin')) NOT NULL
	);

	CREATE TABLE salons(
		id BIGINT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		address VARCHAR(255) NOT NULL,
		city VARCHAR(255) NOT NULL,
		postal_code VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		user_id BIGINT REFERENCES users(id)
	);

	CREATE TABLE hairdressers(
		id BIGINT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		salon_id BIGINT REFERENCES salons(id)
	);

	CREATE TABLE slots(
		id BIGINT PRIMARY KEY,
		start_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
		end_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
		hairdresser_id BIGINT REFERENCES hairdressers(id)
	);

	CREATE TABLE reservations(
		id BIGINT PRIMARY KEY,
		user_id BIGINT REFERENCES users(id),
		slot_id BIGINT REFERENCES slots(id)
	);
`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
