package validation_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/validation"
)

func TestValidationMwS01(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	handlerToTest := validation.ValidateRequestBodyMw(nextHandler)

	jsonString := []byte(`{"title": "Testing", "subtitle": "How To Unit Test", "author": "George M.", "isbn": "247-2547684792", "edition": 2, "year": 2018}`)
	request := httptest.NewRequest("POST", "http://testing/api/v1/book/", bytes.NewBuffer(jsonString))
	handlerToTest.ServeHTTP(httptest.NewRecorder(), request)
}

func TestValidationMwS02(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	handlerToTest := validation.ValidateRequestBodyMw(nextHandler)

	jsonString := []byte(`{"books": [{"title": "Testing", "subtitle": "How To Unit Test", "author": "George M.",
									  "isbn": "247-2547684792", "edition": 2, "year": 2018},
						   			 {"title": "Testing V2", "subtitle": "How To Unit Test", "author": "George M.", "isbn": "247-2547684794",
									  "edition": 1, "year": 2019}]}`)
	request := httptest.NewRequest("POST", "http://testing/api/v1/books/", bytes.NewBuffer(jsonString))
	handlerToTest.ServeHTTP(httptest.NewRecorder(), request)
}
