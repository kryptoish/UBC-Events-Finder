package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/api/greeting", func(w http.ResponseWriter, r *http.Request) {
        response := map[string]string{"message": "Hello from Go!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })

    http.ListenAndServe(":8080", nil)
}