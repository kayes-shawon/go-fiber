package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kayes-shawon/go-fiber/pkg/db"
	"github.com/kayes-shawon/go-fiber/pkg/models"
	"strconv"
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

	err := dbCon.Model(student).Order("id").Select()
	if err != nil {
		return err
	}
	err = c.JSON(student)
	if err != nil {
		return err
	}
	return nil
}

func GetStudentDetails(c *fiber.Ctx) error {
	student := new(models.Student)
	dbCon := db.ConnectDB()
	id := c.Params("id")
	err := dbCon.Model(student).Where("id = ?", id).Select()
	if err != nil {
		return err
	}
	err = c.JSON(student)
	if err != nil {
		return err
	}

	return nil
}


func UpdateStudent(c *fiber.Ctx) error {
	dbCon := db.ConnectDB()
	student := new(models.Student)
	err := c.BodyParser(student)
	if err != nil {
		return err
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	student.Id = idInt

	_, err = dbCon.Model(student).Where("id = ?", idInt).Update()
	if err != nil {
		return err
	}
	err = c.JSON(student)
	if err != nil{
		return err
	}
	return nil
}

func DeleteStudent(c *fiber.Ctx) error {
	dbCon := db.ConnectDB()
	student := new(models.Student)
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = dbCon.Model(student).Where("id = ?", idInt).Delete()
	if err != nil {
		return err
	}
	err = c.SendString(fmt.Sprintf("ID %s has been deleted", id))
	if err != nil {
		return err
	}
	return nil
}