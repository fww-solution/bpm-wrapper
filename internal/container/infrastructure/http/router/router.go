package router

import (
	"bpm-wrapper/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Initialize(app *fiber.App, ctrl *controller.Controller) *fiber.App {
	app.Get("/", monitor.New(monitor.Config{Title: "fww-bpm-wrapper metrics page"}))

	Api := app.Group("/api")

	v1 := Api.Group("/private/v1")

	// bpm
	v1.Post("/workflow", ctrl.SaveWorkflow)

	// passanger
	v1.Put("/passenger", ctrl.UpdatePassenger)

	return app

}