package impl_test

import (
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"os"
	"testing"

	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
  db = impl.SetupDB()
  impl.SetupRepository(db)
  os.Exit(m.Run())
}
