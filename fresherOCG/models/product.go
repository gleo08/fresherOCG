package models

import "gorm.io/gorm"

type Product struct {
	Id          int     `json:"id"`
	BrandId     int     `json:"brand_id"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (p *Product) TableName() string {
	return "product"
}

func (p *Product) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Product{}).Count(&total)
	return total
}

func (p *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []Product

	db.Offset(offset).Limit(limit).Find(&products)

	return products
}
