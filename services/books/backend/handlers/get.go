package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectToPSQL()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM books`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "books could not be retrieved")
		w.Write(json)
		return
	}
	defer rows.Close()

	var books Books
	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Subtitle,
			&book.Author, &book.ISBN, &book.Edition, &book.Year)
		if err != nil {
			w.WriteHeader(500)
			w.Header().Set("Content-Type", "application/json")
			json := utilities.MessageToJson("server could not process request")
			w.Write(json)
			return
		}

		books = append(books, &book)
	}

	err = rows.Err()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "books could not be retrieved")
		w.Write(json)
		return
	}

	jsonBooks, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBooks)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectToPSQL()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}
	defer db.Close()

	id := mux.Vars(r)["id"]
	sqlStatement := `SELECT * FROM books WHERE id=$1`
	row, err := db.Query(sqlStatement, id)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "books could not be retrieved")
		w.Write(json)
		return
	}
	defer row.Close()

	var book Book
	for row.Next() {
		err = row.Scan(&book.ID, &book.Title, &book.Subtitle,
			&book.Author, &book.ISBN, &book.Edition, &book.Year)
		if err != nil {
			w.WriteHeader(500)
			w.Header().Set("Content-Type", "application/json")
			json := utilities.MessageToJson("server could not process request")
			w.Write(json)
			return
		}
	}

	err = row.Err()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "books could not be retrieved")
		w.Write(json)
		return
	}

	jsonBook, err := json.Marshal(book)
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBook)
}
