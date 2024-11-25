package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func main() {
	fmt.Print("Hello World")

	app := fiber.New()

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	log.Fatal(app.Listen(":4000"))
}