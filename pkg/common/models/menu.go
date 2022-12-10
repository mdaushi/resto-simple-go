package models

import (
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Nama  string
	Harga int64
	Stok  int
}
