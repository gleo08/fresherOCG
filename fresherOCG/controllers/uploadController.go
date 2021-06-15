package controllers

import "github.com/gofiber/fiber/v2"

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()

	if err != nil {
		return err
	}

	files := form.File["image"]
	fileName := ""

	for _, file := range files {
		fileName = file.Filename

		if err := c.SaveFile(file, "./uploads/"+fileName); err != nil {
			return err
		}
	}

	return c.JSON(fiber.Map{
		"url": "http://localhost:8080/api/uploads/" + fileName,
	})
}
