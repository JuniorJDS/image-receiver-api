package app

import (
	"image-receiver-api/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MakeApp() *fiber.App {
	log.SetFlags(log.Ltime | log.Lshortfile)

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	allowOriginsCors := ""

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowOriginsCors,
		AllowHeaders:     "*",
		AllowMethods:     "*",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Disposition",
	}))

	routes.Routes(app)

	return app
}
