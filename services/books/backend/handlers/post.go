package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectToPSQL()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}
	defer db.Close()

	var book *Book
	json.NewDecoder(r.Body).Decode(&book)

	var newID int
	sqlStatement := `INSERT INTO books(Title, Subtitle, Author, ISBN, Edition, Year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID`
	err = db.QueryRow(sqlStatement, &book.Title, &book.Subtitle, &book.Author, &book.ISBN, &book.Edition, &book.Year).Scan(&newID)
	if err != nil {
		w.WriteHeader(204)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "the query did not match any tuple")
		w.Write(json)
		return
	}

	book.ID = newID
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

func PostBooks(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectToPSQL()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}
	defer db.Close()

	var manyBooks BookSlice
	json.NewDecoder(r.Body).Decode(&manyBooks)

	var insertedBooks Books
	for _, book := range manyBooks.Books {
		var newID int
		sqlStatement := `INSERT INTO books(Title, Subtitle, Author, ISBN, Edition, Year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID`
		db.QueryRow(sqlStatement, &book.Title, &book.Subtitle, &book.Author, &book.ISBN, &book.Edition, &book.Year).Scan(&newID)

		book.ID = newID
		insertedBooks = append(insertedBooks, book)
	}

	jsonBooks, err := json.Marshal(insertedBooks)
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
