package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID           uint    `json:gorm:"primary_key"`
	Username     string  `json:gorm:"column:username"`
	Email        string  `json:gorm:"column:email;unique_index"`
	Hp           string  `json:gorm:"column:hp"`
	Image        *string `json:gorm:"column:image"`
	PasswordHash string  `json:gorm:"column:password;not null"`
}
type Product struct {
	gorm.Model
	ID       uint    `json:gorm:"primary_key"`
	Name     string  `json:gorm:"column:name"`
	Image    *string `json:gorm:"column:image"`
	Harga    string  `json:gorm:"column:harga"`
	Kategori string  `json:gorm:"column:kategori"`
}
