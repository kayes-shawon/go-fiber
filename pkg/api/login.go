package api

import (
	"github.com/alexandrevicenzi/unchained"
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/db"
	"github.com/kayes-shawon/go-fiber/pkg/models"
	"github.com/kayes-shawon/go-fiber/pkg/utils"
)

func UserLogin(c *fiber.Ctx) error {
	user := new(models.User)

	err := c.BodyParser(user)
	if err != nil {
		return err
	}
	password := user.Password
	dbCon := db.ConnectDB()
	err = dbCon.Model(user).Where("user_name = ?", user.UserName).Select()
	if err != nil {
		return err
	}
	encodedPassword := user.Password
	valid, err := unchained.CheckPassword(password, encodedPassword)
	if err != nil {
		return err
	}
	if !valid {
		err = c.JSON(map[string]interface{} {"message": "Your password is not correct"})
		if err != nil {
			return err
		}
		return nil
	}

	payload := map[string]interface{} {
		"user_name" : user.UserName,
	}

	token, err := utils.Encode(payload)
	if err != nil {
		return err
	}

	data := map[string]interface{} {
		"access_token" : string(token),
	}

	err = c.JSON(data)
	if err != nil {
		return err
	}

	return nil
}
