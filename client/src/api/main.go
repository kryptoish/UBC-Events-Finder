package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "application/json")
		
		response := map[string]string{"message": "Hello from Go!"}
        json.NewEncoder(w).Encode(response)
    })

    http.ListenAndServe(":8080", nil)
}