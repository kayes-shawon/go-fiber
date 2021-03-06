package api

import (
	"github.com/alexandrevicenzi/unchained"
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/db"
	"github.com/kayes-shawon/go-fiber/pkg/models"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil  {
		return err
	}
	dbCon := db.ConnectDB()
	hash, err := unchained.MakePassword(user.Password, unchained.GetRandomString(12), "default")
	if err !=  nil {
		return err
	}
	user.Password = hash
	_, err = dbCon.Model(user).Insert()
	if err != nil {
		return err
	}
	err = c.JSON(user)
	if err != nil {
		return err
	}
	//dbCon.Model(user)
	return nil
}
