package main

import (
	"technical-test/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	quoteHandler := handler.NewQuoteHandler()

	server.Post("/quote", quoteHandler.HandlePostQuote)
	server.Listen(":3002")
}
