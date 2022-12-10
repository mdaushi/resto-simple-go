package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   int
	User     User
	MenuID   int
	Menu     Menu
	Quantity int
}
