package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/db"
	"github.com/kayes-shawon/go-fiber/pkg/models"
)

func CreateStudent(c *fiber.Ctx) error {
	student := new(models.Student)
	//err := json.Unmarshal(c.Body(), &student)
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("%+v",student)

	err := c.BodyParser(student)
	if err != nil {
		return err
	}

	dbCon := db.ConnectDB()
	_, err = dbCon.Model(student).Insert()
	if err != nil {
		return err
	}
	err = c.JSON(student)
	if err != nil {
		return err
	}
	return nil
}

func GetStudentList(c *fiber.Ctx) error {
	student := new([]models.Student)
	dbCon := db.ConnectDB()

	err := dbCon.Model(student).Select()
	if err != nil {
		return err
	}
	err = c.JSON(student)
	if err != nil {
		return err
	}
	return nil
}
