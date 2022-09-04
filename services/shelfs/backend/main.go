package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RaphSku/powerlibrary/tree/main/services/shelfs/handler"
	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")

	result := handler.ExecuteQuery(r.URL.Query().Get("query"), handler.Schema)
	json.NewEncoder(w).Encode(result)
}

func main() {
	err := godotenv.Load("properties.env")
	if err != nil {
		log.Fatalf("error occured with the message: %s", err)
	}

	handler.LoadProperties()

	http.HandleFunc("/shelf", Handler)

	port := os.Getenv("SERVER_PORT")

	fmt.Printf("Server is running on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
