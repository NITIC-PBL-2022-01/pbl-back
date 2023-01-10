package impl

import (
	"fmt"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"os"
	"time"

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

  db.AutoMigrate(&domain.Event{}, &domain.ToDo{}, &domain.User{}, &domain.Tag{})

  return db
}

func SetupRepository(db *gorm.DB) {
  Event.db = db
}

func InsertTestData(db *gorm.DB) {
  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    panic(err)
  }

  user := domain.ConstructUser(email, "TEST DATA", false);

  db.Create(&user);

  tag := domain.ConstructTag(domain.ID("0"), "hoge", "#fff", []domain.User{user}, []domain.User{}, domain.None)
  db.Create(&tag);

  date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
  duration := time.Minute * 60

  event := domain.ConstructEvent(
    "0",
    date,
    duration,
    "Title",
    "Detail",
    tag,
    "#000",
    user,
    "1号館",
    domain.ConstructRepeat(
      domain.Week,
      1,
      time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
      time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
    ),
  )

  db.Create(&event);
}
