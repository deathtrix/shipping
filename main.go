package main

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

var packSizes = []int{250, 500, 1000, 2000, 5000}

func main() {
	// Show file and line number in logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Create fiber app and use the optimized goccy json library instead of default
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Define routes - handlers defined in endpoints.go
	app.
		Use(logger.New()).
		Static("/", "./static").
		Get("/health", func(c fiber.Ctx) error {
			return c.SendString("ok")
		}).
		Group("/api/v1").
		Get("/packages/:items", getPackages).
		Get("/sizes", getPackageSizes).
		Put("/sizes", setPackageSizes)

	app.Listen(":8080")
}
