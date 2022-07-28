package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

func ValidateRequestBodyMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/book/" {
			requestBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {

			}
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))

			var book *Book
			err = json.NewDecoder(r.Body).Decode(&book)
			if err != nil {

			}

			err = validateBook(book)
			if err != nil {
				w.WriteHeader(400)
				w.Header().Set("Content-Type", "application/json")
				response := make(map[string]string)
				response["error"] = fmt.Sprintf("Validation Error occured with the following message: %s", err)
				json, err := json.Marshal(response)
				if err != nil {
					log.Fatalf("JSON Marshal did not work. Err: %s", err)
				}
				w.Write(json)
				return
			}

			r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBytes))
		}

		next.ServeHTTP(w, r)
	})
}

func validateBook(book *Book) error {
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
