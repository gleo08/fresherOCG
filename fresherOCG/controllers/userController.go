package controllers

import (
	"strconv"

	"github.com/gleo08/fresherOCG/database"
	"github.com/gleo08/fresherOCG/models"
	"github.com/gofiber/fiber/v2"
)

func AllUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func GetUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User

	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

func UpdateUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: id,
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)

}

func DeleteUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	database.DB.Exec("DELETE FROM user WHERE id = ?", id)

	return nil
}
