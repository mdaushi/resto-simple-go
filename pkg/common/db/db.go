package db

import (
	"log"

	"github.com/mdaushi/resto-simple-go/pkg/common/models"
	"github.com/mdaushi/resto-simple-go/pkg/common/seeds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// migrate models
	db.AutoMigrate(&models.User{}, &models.Menu{}, &models.Order{}, &models.Receipt{})

	// Apply Seed
	// comment jika diperlukan
	seeds.ApplyUsersSeed(db)
	seeds.ApplyMenusSeed(db)

	return db
}
