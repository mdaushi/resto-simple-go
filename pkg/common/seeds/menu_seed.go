package seeds

import (
	"fmt"

	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"gorm.io/gorm"
)

func menusSeed() []models.Menu {
	nasgor := models.Menu{
		Nama:  "Nasi Goreng",
		Harga: 10000,
		Stok:  10,
	}

	naskung := models.Menu{
		Nama:  "Nasi Kuning Ayam",
		Harga: 15000,
		Stok:  4,
	}

	menus := []models.Menu{nasgor, naskung}

	return menus
}

func ApplyMenusSeed(db *gorm.DB) {
	menus := menusSeed()
	for _, menu := range menus {
		db.Create(&menu)
	}

	fmt.Println("Successfully added data Menus to database")
}
