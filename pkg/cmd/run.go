package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/api"
	"github.com/kayes-shawon/go-fiber/pkg/middleware"
)

func UnRestricted(app *fiber.App)  {
	userGroup := app.Group("/user", func(c *fiber.Ctx) error {
		return c.Next()
	})

	userGroup.Post("/login/", func(c *fiber.Ctx) error {
		return api.UserLogin(c)
	})

	userGroup.Post("/login-refresh/", func(c *fiber.Ctx) error {
		return api.UserLoginRefresh(c)
	})

	userGroup.Post("/create", func(c *fiber.Ctx) error {
		return api.CreateUser(c)
	})
}

func AuthRequired(app *fiber.App) {
	studentGroup := app.Group("/student", middleware.Auth)

	studentGroup.Post("/create", api.CreateStudent)

	studentGroup.Get("/list/", api.GetStudentList)

	studentGroup.Get("/:id", api.GetStudentDetails)

	studentGroup.Post("/update/:id", api.UpdateStudent)

	studentGroup.Delete("/delete/:id", api.DeleteStudent)
}

func RunServer()  {
	app := fiber.New()

	UnRestricted(app)

	AuthRequired(app)


	app.Listen(":8080")
}
