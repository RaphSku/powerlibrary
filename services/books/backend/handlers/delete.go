package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db, err := ConnectToPSQL()
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}
	defer db.Close()

	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson("server could not process request")
		w.Write(json)
		return
	}

	var book Book
	var id int
	var title string
	sqlStatement := `DELETE FROM books WHERE id=$1 RETURNING id, title`
	err = db.QueryRow(sqlStatement, &bookID).Scan(&id, &title)
	if err != nil {
		w.WriteHeader(204)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.ErrorMessageToJson(err, "book could not be deleted because it does not exist")
		w.Write(json)
		return
	}

	book.ID = id
	book.Title = title

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
