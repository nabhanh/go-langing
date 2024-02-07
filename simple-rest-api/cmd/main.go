// package name doesn't have to be the same as the folder name or the file name
package main

import (
	"log"
	"nabhanh/simple-rest-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// func helloworld (c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!")
// }

func main() {
	// Entry point of the application
	// Create a new fiber app
	app := fiber.New()

	// Define a route
	// the reason G is capitalized is because it's a public function
	// app.Get("/ping", func(c *fiber.Ctx) error {
	// 	return c.SendString("pong")
	// })

	// // Alternative way to define a route
	// app.Get("/hello", helloworld)

	app.Post("/api/products", handlers.CreateProduct)

	// Start the server on port 3000

	log.Fatal(app.Listen(":3000")) // super weird syntax, but it's the same as app.Listen(":3000") but it also logs the error if there is one and exits the program with os.Exit(1)
}
