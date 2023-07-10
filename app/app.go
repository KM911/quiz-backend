package app

import (
	"github.com/gofiber/fiber/v2"
	"quiz/router"
)

func Start() {
	app := fiber.New()
	//app.Static("/", "./dist")
	router.InitRouter(app)
	app.Listen(":3000")
}
