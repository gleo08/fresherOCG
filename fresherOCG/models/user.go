package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Name     string `json:"name"`
	Role     int    `json:"role"`
}

func (u *User) TableName() string {
	return "user"
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (u *User) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&User{}).Count(&total)
	return total
}

func (u *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Offset(offset).Limit(limit).Find(&users)

	return users
}
