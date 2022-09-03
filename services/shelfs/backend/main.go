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

func main() {
	err := godotenv.Load("properties.env")
	if err != nil {
		log.Fatalf("error occured with the message: %s", err)
	}

	handler.LoadProperties()

	http.HandleFunc("/shelf", func(w http.ResponseWriter, r *http.Request) {
		result := handler.ExecuteQuery(r.URL.Query().Get("query"), handler.Schema)
		json.NewEncoder(w).Encode(result)
	})

	port := os.Getenv("SERVER_PORT")

	fmt.Printf("Server is running on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
