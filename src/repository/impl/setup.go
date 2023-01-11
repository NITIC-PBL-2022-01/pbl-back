package impl

import (
	"fmt"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Event = EventRepository{}
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.Event{}, &domain.ToDo{}, &domain.User{}, &domain.Tag{})
	if err != nil {
		panic(err)
	}

	return db
}

func SetupRepository(db *gorm.DB) {
	Event.db = db
}
