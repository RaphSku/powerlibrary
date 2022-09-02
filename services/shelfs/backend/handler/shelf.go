package handler

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "raphael"
	password = "test2"
	dbname   = "powerlibrary-shelf"
)

func ConnectToPSQL() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	err = db.Ping()

	return db, err
}

type Shelf struct {
	ID       int64  `json:"id"`
	Room     string `json:"room"`
	Location string `json:"location"`
}

type Shelfs []*Shelf
