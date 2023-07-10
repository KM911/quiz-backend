package router

import (
	"github.com/gofiber/fiber/v2"
	"quiz/service"
)

func InitRouter(app *fiber.App) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	quizGroup := app.Group("/quiz")

	quizGroup.Get("/choice/num", service.GetChoiceNum)
	quizGroup.Get("/choice/id/:id", service.GetChoiceById)
}
