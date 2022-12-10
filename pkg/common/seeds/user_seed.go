package seeds

import (
	"fmt"

	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"gorm.io/gorm"
)

func usersSeed() []models.User {
	firdaus := models.User{
		Nama:   "firdaus",
		Alamat: "jl. sabutung",
		Telp:   "342904902",
	}

	daus := models.User{
		Nama:   "daus",
		Alamat: "jl. sabutung",
		Telp:   "3423242",
	}

	users := []models.User{firdaus, daus}

	return users
}

func ApplyUsersSeed(db *gorm.DB) {
	users := usersSeed()
	for _, user := range users {
		db.Create(&user)
	}

	fmt.Println("Successfully added data Users to database")
}
