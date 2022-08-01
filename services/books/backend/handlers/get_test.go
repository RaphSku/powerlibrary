package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/handlers"
	"github.com/gorilla/mux"
)

func TestGetBooksS01(t *testing.T) {
	Setup(t)

	request, err := http.NewRequest("GET", "http://localhost/api/v1/books/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetBooks)
	handler.ServeHTTP(recorder, request)

	body := recorder.Body
	var books handlers.Books
	json.NewDecoder(body).Decode(&books)

	for index := range books {
		TearDown(t, books[index].ID)
	}

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestGetBooksByIdS01(t *testing.T) {
	targetId := Setup(t)

	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost/api/v1/books/%v", targetId), nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{"id": fmt.Sprintf("%v", targetId)}
	request = mux.SetURLVars(request, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetBookById)
	handler.ServeHTTP(recorder, request)

	var book handlers.Book
	json.NewDecoder(recorder.Body).Decode(&book)

	TearDown(t, book.ID)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
