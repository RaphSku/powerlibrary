package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
	"github.com/gorilla/mux"
)

func PutBook(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectToPSQL()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}
	defer db.Close()

	targetBookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}

	var targetBook *Book
	json.NewDecoder(r.Body).Decode(&targetBook)
	targetBook.ID = targetBookID

	var updatedBook Book
	sqlStatement := `UPDATE books SET title=$1, subtitle=$2, author=$3, isbn=$4, edition=$5, year=$6 WHERE id=$7 RETURNING id, title, subtitle,
					 author, isbn, edition, year`
	err = db.QueryRow(sqlStatement, &targetBook.Title, &targetBook.Subtitle, &targetBook.Author,
		&targetBook.ISBN, &targetBook.Edition, &targetBook.Year, &targetBook.ID).Scan(&updatedBook.ID, &updatedBook.Title, &updatedBook.Subtitle,
		&updatedBook.Author, &updatedBook.ISBN, &updatedBook.Edition, &updatedBook.Year)
	if err != nil {
		w.WriteHeader(204)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "book could not be updated because it does not exist")
		w.Write(json)
		return
	}

	jsonBook, err := json.Marshal(updatedBook)
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
