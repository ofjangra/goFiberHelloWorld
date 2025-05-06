package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	envErr := godotenv.Load(".env")
	if envErr != nil {
		panic("Error loading .env file")
	}

	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port
	}
	server.Listen(":" + port)
}
