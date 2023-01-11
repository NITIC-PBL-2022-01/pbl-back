package mock_test

import (
	"nitic-pbl-2022-01/pbl-back/src/repository/mock"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	mock.ConstructTestData()
	os.Exit(m.Run())
}
