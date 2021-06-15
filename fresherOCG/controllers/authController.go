package controllers

import (
	"strconv"
	"time"

	"github.com/gleo08/fresherOCG/database"
	"github.com/gleo08/fresherOCG/models"
	"github.com/gleo08/fresherOCG/util"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user := models.User{
		Email: data["email"],
		Name:  data["name"],
		Role:  2,
	}
	user.SetPassword(data["password"])
	var resultQuery models.User
	database.DB.Where("email = ?", data["email"]).First(&resultQuery)
	if resultQuery.Id != 0 {
		return c.JSON(fiber.Map{
			"message": "Email is already in use",
		})
	} else {
		database.DB.Create(&user)
		return c.JSON(user)
	}

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		panic(err)
	}

	var user models.User

	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		return c.JSON(fiber.Map{
			"message": "NOT FOUND",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return c.JSON(fiber.Map{
			"message": "Password is incorrect",
		})
	}

	token, err := util.GenerateJwt(strconv.Itoa(user.Id))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}

func UpdateInfo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(err)
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)
	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id:    userId,
		Name:  data["name"],
		Email: data["email"],
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(err)
	}

	cookie := c.Cookies("jwt")

	id, _ := util.ParseJwt(cookie)

	userId, _ := strconv.Atoi(id)

	user := models.User{
		Id: userId,
	}

	user.SetPassword(data["password"])

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)

}
