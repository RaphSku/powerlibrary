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

func Setup(t *testing.T) int {
	body := []byte(`{"title": "Testing", "subtitle": "How To Unit Test Advanced", "author": "George M.", "isbn": "247-2257225794", "edition": 2, "year": 2018}`)
	request, err := http.NewRequest("POST", "http://localhost/api/v1/book/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostBook)
	handler.ServeHTTP(recorder, request)

	var targetBook handlers.Book
	json.NewDecoder(recorder.Body).Decode(&targetBook)

	return targetBook.ID
}

func TestDeleteBook(t *testing.T) {
	targetId := Setup(t)
	request, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost/api/v1/books/%v", targetId), nil)
	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{"id": fmt.Sprintf("%v", targetId)}
	request = mux.SetURLVars(request, vars)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteBook)
	handler.ServeHTTP(recorder, request)

	var targetBook handlers.Book
	json.NewDecoder(recorder.Body).Decode(&targetBook)
	TearDown(t, targetBook.ID)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := fmt.Sprintf(`{"id":%v,"title":"Testing","subtitle":"","author":"","isbn":"","edition":0,"year":0}`, targetBook.ID)
	got := fmt.Sprintf(`{"id":%v,"title":"%s","subtitle":"%s","author":"%s","isbn":"%s","edition":%v,"year":%v}`, targetBook.ID, targetBook.Title,
		targetBook.Subtitle, targetBook.Author, targetBook.ISBN, targetBook.Edition, targetBook.Year)
	if got != expected {
		t.Errorf("unexpected body: got %v, want %v", got, expected)
	}
}
