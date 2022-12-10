package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama   string
	Alamat string
	Telp   string
}
