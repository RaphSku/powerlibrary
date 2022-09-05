package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
)

func CheckShelf(name string) (bool, string) {
	response, err := http.Get(fmt.Sprintf(`http://10.104.78.169:8081/shelf?query={shelf(name:%q){name}}`, name))
	if err != nil {
		return false, "Query for shelf name failed!"
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, "Response Body could not be read!"
	}

	bodyString := string(body)
	match, err := regexp.MatchString(`"name":\\K"\w*"`, bodyString)
	if err != nil {
		return false, "Regex Pattern Matching did not work!"
	}

	return match, "Shelf Name is empty!"
}

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

	ok, msg := CheckShelf(book.ShelfName)
	if !ok {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json := utilities.MessageToJson(fmt.Sprintf("a problem arose: %s", msg))
		w.Write(json)
		return
	}

	var newID int
	sqlStatement := `INSERT INTO books(Title, Subtitle, Author, ISBN, Edition, Year, ShelfName, ShelfLevel) 
					 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ID`
	err = db.QueryRow(sqlStatement, &book.Title, &book.Subtitle, &book.Author, &book.ISBN, &book.Edition,
		&book.Year, &book.ShelfName, &book.ShelfLevel).Scan(&newID)
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
		ok, msg := CheckShelf(book.ShelfName)
		if !ok {
			w.WriteHeader(500)
			w.Header().Set("Content-Type", "application/json")
			json := utilities.MessageToJson(fmt.Sprintf("a problem arose: %s", msg))
			w.Write(json)
			return
		}

		var newID int
		sqlStatement := `INSERT INTO books(Title, Subtitle, Author, ISBN, Edition, Year, ShelfName, ShelfLevel) 
						 VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING ID`
		db.QueryRow(sqlStatement, &book.Title, &book.Subtitle, &book.Author, &book.ISBN, &book.Edition, &book.Year,
			&book.ShelfName, &book.ShelfLevel).Scan(&newID)

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
