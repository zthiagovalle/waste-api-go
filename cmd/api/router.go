package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zthiagovalle/waste-api-go/cmd/api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListWastes)
}
