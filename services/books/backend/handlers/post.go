package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	db := ConnectToPSQL()
	defer db.Close()

	var book *Book
	json.NewDecoder(r.Body).Decode(&book)
	fmt.Println(book)

	var newID int
	sqlStatement := `INSERT INTO books(Title, Subtitle, Author, ISBN, Edition, Year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID`
	err := db.QueryRow(sqlStatement, &book.Title, &book.Subtitle, &book.Author, &book.ISBN, &book.Edition, &book.Year).Scan(&newID)
	if err != nil {
		panic(err)
	}

	book.ID = newID
	jsonBook, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBook)
}

func PostBooks(w http.ResponseWriter, r *http.Request) {
	db := ConnectToPSQL()
	defer db.Close()

	var manyBooks Books
	json.NewDecoder(r.Body).Decode(&manyBooks)

	var insertedBooks Books
	for _, book := range manyBooks {
		var newID int
		sqlStatement := `INSERT INTO books(Title, Subtitle, Author, ISBN, Edition, Year) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID`
		err := db.QueryRow(sqlStatement, &book.Title, &book.Subtitle, &book.Author, &book.ISBN, &book.Edition, &book.Year).Scan(&newID)
		if err != nil {
			panic(err)
		}

		book.ID = newID
		insertedBooks = append(insertedBooks, book)
	}

	jsonBooks, err := json.Marshal(insertedBooks)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBooks)
}
