package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/api"
)

func RunServer()  {
	app := fiber.New()

	app.Get("/student/list/", func(c *fiber.Ctx) error {
		return api.GetStudentList(c)
	})

	app.Post("/student/", func(c *fiber.Ctx) error {
		return api.CreateStudent(c)
	})

	app.Get("/student/:id", func(c *fiber.Ctx) error {
		return c.SendString("Student Details")
	})

	app.Post("/student/update/:id", func(c *fiber.Ctx) error {
		return c.SendString("Student Update")
	})

	app.Delete("/student/delete/:id", func(c *fiber.Ctx) error {
		return c.SendString("Student Delete")
	})


	app.Listen(":8080")
}
