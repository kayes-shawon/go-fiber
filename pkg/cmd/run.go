package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/api"
)

func RunServer()  {
	app := fiber.New()

	app.Post("/student/", func(c *fiber.Ctx) error {
		return api.CreateStudent(c)
	})

	app.Get("/student/list/", func(c *fiber.Ctx) error {
		return api.GetStudentList(c)
	})

	app.Get("/student/:id", func(c *fiber.Ctx) error {
		return api.GetStudentDetails(c)
	})

	app.Post("/student/update/:id", func(c *fiber.Ctx) error {
		return api.UpdateStudent(c)
	})

	app.Delete("/student/delete/:id", func(c *fiber.Ctx) error {
		return api.DeleteStudent(c)
	})


	app.Listen(":8080")
}
