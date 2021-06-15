package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

const LIMIT_VALUE = 5

func Paginate(db *gorm.DB, entity Entity, page int) fiber.Map {
	limit := LIMIT_VALUE

	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)
	lastPage := float64(int(total) / limit)

	return fiber.Map{
		"data": data,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(lastPage),
		},
	}

}
