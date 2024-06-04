package routes

import (
	"be-eeg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Ping
	app.Get("/api/ping", controllers.GetPing)

	// Bayi
	app.Post("/api/newborn-data", controllers.CreateNewbornData)
}
