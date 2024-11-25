package main

import (
	// "fmt"
	// "log"

	// "github.com/gofiber/fiber/v2"
	"encoding/json"
    "net/http"
)


func main() {
	// fmt.Print("Hello World")

	// app := fiber.New()

	// app.Get("/healthcheck", func(c *fiber.Ctx) error {
	// 	return c.SendString("OK")
	// })

	http.HandleFunc("/api/greeting", func(w http.ResponseWriter, r *http.Request) {
        response := map[string]string{"message": "Hello from Go!"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    })

    http.ListenAndServe(":8080", nil)

	// log.Fatal(app.Listen(":4000"))
}