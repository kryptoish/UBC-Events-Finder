package auth

import (
    "fmt"
    "net/http"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    if code == "" {
        http.Error(w, "Authorization code not found", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Authorization Code: %s", code)
}