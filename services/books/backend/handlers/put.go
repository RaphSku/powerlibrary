package handlers

import (
	"encoding/json"
	"net/http"
)

func PutBook(w http.ResponseWriter, r *http.Request) {
	db := ConnectToPSQL()
	defer db.Close()

	var targetBook *Book
	json.NewDecoder(r.Body).Decode(&targetBook)

	var updatedBook Book
	var newID int
	var updatedTitle string
	var updatedSubtitle string
	var updatedAuthor string
	var updatedISBN string
	var updatedEdition int
	var updatedYear int
	sqlStatement := `UPDATE books SET title=$1, subtitle=$2, author=$3, isbn=$4, edition=$5, year=$6 WHERE id=$7 RETURNING id, title, subtitle, 
					 author, isbn, edition, year`
	err := db.QueryRow(sqlStatement, &targetBook.Title, &targetBook.Subtitle, &targetBook.Author,
		&targetBook.ISBN, &targetBook.Edition, &targetBook.Year, &targetBook.ID).Scan(&newID, &updatedTitle, &updatedSubtitle,
		&updatedAuthor, &updatedISBN, &updatedEdition, &updatedYear)
	if err != nil {
		panic(err)
	}

	updatedBook.ID = newID
	updatedBook.Title = updatedTitle
	updatedBook.Subtitle = updatedSubtitle
	updatedBook.Author = updatedAuthor
	updatedBook.ISBN = updatedISBN
	updatedBook.Edition = updatedEdition
	updatedBook.Year = updatedYear
	jsonBook, err := json.Marshal(updatedBook)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBook)
}
