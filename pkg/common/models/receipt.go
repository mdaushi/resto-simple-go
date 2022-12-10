package models

import (
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	OrderID int
	Order   Order
}
