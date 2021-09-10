package routes

import "github.com/gofiber/fiber/v2"

func Run() {
	app := fiber.New()
	router := app.Group("/v1")

	setupRoutes(&router)

	app.Listen(":8000")
}

func setupRoutes(router *fiber.Router) {
	initAuthRoutes(router)
}
