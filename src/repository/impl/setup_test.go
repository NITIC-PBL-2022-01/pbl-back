package impl_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func InsertTestData(db *gorm.DB) {
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		panic(err)
	}

	user := domain.ConstructUser(email, "TEST DATA", false)

	db.Create(&user)

  emailTag, err := domain.ConstructEmail("tag@example.com")
  if err != nil {
    panic(err)
  }

	tagUser := domain.ConstructUser(emailTag, "TAG TARO", true)

  db.Create(&tagUser)

	tag := domain.ConstructTag(domain.ID("0"), "hoge", "#fff", []domain.User{user}, []domain.User{}, domain.None)
	db.Create(&tag)

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
			time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local),
			time.Date(2022, 2, 28, 0, 0, 0, 0, time.Local),
		),
	)

	db.Create(&event)
}

func TestMain(m *testing.M) {
	db = impl.SetupDB()
	impl.SetupRepository(db)
  InsertTestData(db)
	os.Exit(m.Run())
}
