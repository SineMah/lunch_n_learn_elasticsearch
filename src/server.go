//go:build server
// +build server

package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"mincedmind.com/elasticsearch/routes"
	"os"
)

func setUpRoutes(app *fiber.App) {
	app.Post("/search/:index", routes.Search)
	app.Get("/", routes.Hello)
}

func main() {
	godotenv.Load()

	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
