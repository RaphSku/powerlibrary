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
	password = "test1"
	dbname   = "powerlibrary"
)

func ConnectToPSQL() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	err = db.Ping()

	return db, err
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

type BookSlice struct {
	Books []*Book `json:"books"`
}
