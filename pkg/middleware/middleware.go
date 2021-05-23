package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/utils"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	fmt.Println(token)

	if len(token) < 1 {
		err := c.SendString("no token found")
		if err != nil {
			return err
		}
		return nil
	}
	valid, err := utils.Decode(token)
	if err != nil {
		return err
	}
	if !valid {
		err := c.SendString("authorization failed")
		if err != nil {
			return err
		}
		return nil
	}

	return c.Next()
}
