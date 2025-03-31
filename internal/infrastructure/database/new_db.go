package database

import (
	"EmailNGo/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=nxm123m16 dbname=wblm-tracking port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to database")
	}

	//migrations, if u use all the structs in the same line as above gorm will create the foreign keys too
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
