package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/handlers"
	"github.com/gorilla/mux"
)

func TestPutBookS01(t *testing.T) {
	targetId := Setup(t)

	body := []byte(`{"title": "Testing", "subtitle": "Unit Test - The Easy Way", "author": "George M.", 
				     "isbn": "247-2257225794", "edition": 1, "year": 2020, "shelf_name": "Violet", "shelf_level": 2}`)
	request, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost/api/v1/books/%v", targetId), bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{"id": fmt.Sprintf("%v", targetId)}
	request = mux.SetURLVars(request, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PutBook)
	handler.ServeHTTP(recorder, request)

	var book *handlers.Book
	json.NewDecoder(recorder.Body).Decode(&book)

	TearDown(t, targetId)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v, want %v", status, http.StatusOK)
	}

	comparator := [][]string{{"Testing", book.Title}, {"George M.", book.Author}, {"247-2257225794", book.ISBN}}
	for index := range comparator {
		if comparator[index][0] != comparator[index][1] {
			t.Errorf("the following attribute does not match: got %v, want %v", comparator[index][1], comparator[index][0])
		}
	}

	if book.Subtitle != "Unit Test - The Easy Way" && book.Edition != 1 && book.Year != 2020 {
		t.Error("the following attributes did not get updated: subtitle or/and edition or/and year")
		t.Errorf("got %v, want %v\ngot %v, want %v\ngot %v, want %v", book.Subtitle, "Unit Test - The Easy Way", book.Edition, 1, book.Year, 2020)
	}
}
