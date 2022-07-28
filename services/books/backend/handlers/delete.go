package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	db := ConnectToPSQL()
	defer db.Close()

	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}

	var book Book
	var id int
	var title string
	sqlStatement := `DELETE FROM books WHERE id=$1 RETURNING id, title`
	err = db.QueryRow(sqlStatement, &bookID).Scan(&id, &title)
	if err != nil {
		panic(err)
	}

	book.ID = id
	book.Title = title

	jsonBook, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBook)
}
