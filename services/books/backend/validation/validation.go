package validation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/handlers"
	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
)

func ValidateRequestBodyMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/book/" {
			requestBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				json := utilities.ErrorMessageToJson(err, "request body could not be read")
				w.Write(json)
				return
			}
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))

			var book *handlers.Book
			err = json.NewDecoder(r.Body).Decode(&book)
			if err != nil {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				json := utilities.ErrorMessageToJson(err, "request body could not be decoded into book")
				w.Write(json)
				return
			}

			err = validateBook(book)
			if err != nil {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				json := utilities.ErrorMessageToJson(err, "validation failed")
				w.Write(json)
				return
			}

			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))
		} else if r.URL.Path == "/api/v1/books/" {
			requestBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				json := utilities.ErrorMessageToJson(err, "request body could not be read")
				w.Write(json)
				return
			}
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))

			var books handlers.BookSlice
			err = json.NewDecoder(r.Body).Decode(&books)
			if err != nil {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				json := utilities.ErrorMessageToJson(err, "request body could not be decoded into books")
				w.Write(json)
				return
			}

			validationFailed := false
			var lastEmittedError error
			errorMessage := ""
			for _, book := range books.Books {
				err = validateBook(book)
				if err != nil {
					errorMessage = errorMessage + "err\n"
					lastEmittedError = err
					validationFailed = true
				}
			}

			if validationFailed {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				json := utilities.ErrorMessageToJson(lastEmittedError, "validation failed")
				w.Write(json)
				return
			}

			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))
		}

		next.ServeHTTP(w, r)
	})
}

func validateBook(book *handlers.Book) error {
	title := book.Title
	uni := []rune(title)
	if len(uni) > 40 || len(uni) == 0 {
		return errors.New(`title cannot be longer than 40 character and cannot have length 0,
			do not forget that the title is required`)
	}

	subtitle := book.Subtitle
	uni = []rune(subtitle)
	if len(uni) > 30 || len(uni) == 0 {
		return errors.New(`subtitle cannot be longer than 30 characters and cannot have length 0,
			do not forget that the subtitle is required`)
	}

	author := book.Author
	uni = []rune(author)
	if len(uni) > 20 || len(uni) == 0 {
		return errors.New(`author cannot be longer than 20 characters and cannot have length 0,
			do not forget that the author is required`)
	}

	isbn := book.ISBN
	match, _ := regexp.MatchString("\\d{3}-\\d{10}", isbn)
	uni = []rune(isbn)
	if len(uni) != 14 && !match || len(uni) == 0 {
		return errors.New(`isbn has to be 14 characters long and has to match the pattern of an ISBN-13 number,
			do not forget that the isbn is required`)
	}

	edition := book.Edition
	if edition <= 0 || edition > 20 {
		return errors.New(`edition has to be larger than 0, and cannot be larger than 20,
			do not forget that the edition is required`)
	}

	year := book.Year
	yr, _, _ := time.Now().Date()
	if year <= 0 || year > yr {
		return errors.New(`the book was very likely not released before the year 0 and cannot exceed our current year,
			do not forget that the year is required`)
	}

	return nil
}
