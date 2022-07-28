package handlers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "raphael"
	password = "leber51a"
	dbname   = "powerlibrary"
)

func ConnectToPSQL() *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Author   string `json:"author"`
	ISBN     string `json:"isbn"`
	Edition  int    `json:"edition"`
	Year     int    `json:"year"`
}

type Books []*Book
