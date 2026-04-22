package database

import (
	"database/sql"
	"fmt"
)

func NewPostgresConnection() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=tu_password dbname=config_service sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Conexión a PostgreSQL establecida exitosamente")

	return db, nil
}
