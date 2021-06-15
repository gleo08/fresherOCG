package controllers

import (
	"strconv"

	"github.com/gleo08/fresherOCG/database"
	"github.com/gleo08/fresherOCG/models"
	"github.com/gofiber/fiber/v2"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.JSON(err)
	}

	database.DB.Create(&product)

	return c.JSON(product)
}

func GetProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product models.Product

	database.DB.Where("id = ?", id).First(&product)
	return c.JSON(product)
}

func UpdateProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{
		Id: id,
	}

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(product)

	return c.JSON(product)

}

func DeleteProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := database.DB.Exec("DELETE FROM product WHERE id = ?", id)

	if err != nil {
		return c.JSON(err)
	}
	return nil
}
