package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RaphSku/powerlibrary/tree/main/services/shelfs/handler"
)

func main() {
	http.HandleFunc("/shelf", func(w http.ResponseWriter, r *http.Request) {
		result := handler.ExecuteQuery(r.URL.Query().Get("query"), handler.Schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
