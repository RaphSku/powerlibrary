package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db := ConnectToPSQL()
	defer db.Close()

	sqlStatement := `SELECT * FROM books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var booksSlice Books
	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Subtitle,
			&book.Author, &book.ISBN, &book.Edition, &book.Year)
		if err != nil {
			panic(err)
		}

		booksSlice = append(booksSlice, &book)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	jsonBooks, err := json.Marshal(booksSlice)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	db := ConnectToPSQL()
	defer db.Close()

	id := mux.Vars(r)["id"]
	sqlStatement := `SELECT * FROM books WHERE id=$1`
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	var book Book
	for row.Next() {
		err = row.Scan(&book.ID, &book.Title, &book.Subtitle,
			&book.Author, &book.ISBN, &book.Edition, &book.Year)
		if err != nil {
			panic(err)
		}
	}

	err = row.Err()
	if err != nil {
		panic(err)
	}

	jsonBook, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBook)
}
