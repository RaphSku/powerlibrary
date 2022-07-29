package utilities

import (
	"encoding/json"
	"fmt"
	"log"
)

func ErrorMessageToJson(err error, message string) []byte {
	response := make(map[string]string)
	response["message"] = message
	response["error"] = fmt.Sprintf("%s", err)
	json, err := json.Marshal(response)
	if err != nil {
		log.Printf("JSON Marshal did not work. Err: %s", err)
	}

	return json
}

func MessageToJson(message string) []byte {
	response := make(map[string]string)
	response["message"] = message
	json, err := json.Marshal(response)
	if err != nil {
		log.Printf("JSON Marshal did not work. Err: %s", err)
	}

	return json
}
