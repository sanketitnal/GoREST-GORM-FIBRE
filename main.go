package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanketitnal/gobasicrest/database"
	"github.com/sanketitnal/gobasicrest/services/user"
)

func initilizeServices(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	user.RegisterRoutes(app)
}

func constructor() {
	database.ConnectAllDatabases()
}

func main() {
	constructor()

	app := fiber.New()
	initilizeServices(app)
	app.Listen(":80")
}
