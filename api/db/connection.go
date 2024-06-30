package db

import (
	"fmt"

	"anime/config"
	"anime/db/seeders"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	userSeeder := seeders.NewUserSeeder(db)
	userSeeder.SetUsers()

	return db
}
