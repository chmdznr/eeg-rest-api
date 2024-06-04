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
	app.Post("/api/newborn-eeg", controllers.CreateNewbornEEG)
	app.Post("/api/newborn-cv", controllers.CreateNewbornCv)

	// Ibu hamil
	app.Post("/api/pregnant-data", controllers.CreatePregnantData)
	app.Post("/api/pregnant-eeg", controllers.CreatePregnantEEG)
}
