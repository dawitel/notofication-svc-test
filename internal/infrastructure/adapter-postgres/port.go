package adapterpostgres

import (
	"database/sql"
	"errors"
	"os"
)

type Db struct {
	Conn *sql.DB
}

func New() (*Db, error) {
	connString := os.Getenv("CONN_STRING")
	if connString == "" {
		return nil, errors.New("connection string is missing")
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Db{
		Conn: db,
	}, nil
}
