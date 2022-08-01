package utilities_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/RaphSku/powerlibrary/tree/main/services/books/utilities"
)

func TestErrToJsonS01(t *testing.T) {
	err := errors.New("Testing Error Message to JSON")
	gotJson := utilities.ErrorMessageToJson(err, "Validation failed, something unexpected happened")

	response := map[string]string{"message": "Validation failed, something unexpected happened",
		"error": "Testing Error Message to JSON"}
	wantJson, _ := json.Marshal(response)

	if len(gotJson) != len(wantJson) {
		t.Errorf("Json objects length do not match! Want: %v, Got: %v", len(wantJson), len(gotJson))
	}
	for i := range gotJson {
		if gotJson[i] != wantJson[i] {
			t.Errorf("Json objects do not match! Want: %v, Got: %v", wantJson, gotJson)
		}
	}
}

func TestMessageToJsonS01(t *testing.T) {
	gotJson := utilities.MessageToJson("Testing MessageToJson.\n Here we go!")

	response := map[string]string{"message": "Testing MessageToJson.\n Here we go!"}
	wantJson, _ := json.Marshal(response)

	if len(gotJson) != len(wantJson) {
		t.Errorf("Json objects length do not match! Want: %v, Got: %v", len(wantJson), len(gotJson))
	}
	for i := range gotJson {
		if gotJson[i] != wantJson[i] {
			t.Errorf("Json objects do not match! Want: %v, Got: %v", wantJson, gotJson)
		}
	}
}
