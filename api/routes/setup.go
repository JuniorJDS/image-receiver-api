package routes

import (
	"image-receiver-api/api/routes/endpoints"
	"image-receiver-api/config"

	"github.com/gofiber/fiber/v2"
)

var settings = config.GetSettings()

func Routes(app *fiber.App) {
	r := app.Group(settings["API_V1"])

	h := endpoints.NewHello()
	uploadRoute := endpoints.NewUploadRoute()

	r.Get("/hello", h.GetHello)
	r.Post("/upload", uploadRoute.Upload)
}
