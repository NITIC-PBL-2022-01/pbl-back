package impl_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
  email, err := domain.ConstructEmail("test@example.com")
  if err != nil {
    t.Fatal(err)
  }

  user := domain.ConstructUser(email, "THISIS TEST", false)
  created, err := impl.User.Create(user)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, email, created.Email)
  assert.Equal(t, "THISIS TEST", created.Name)
  assert.Equal(t, false, created.IsStudent)

  target := domain.User{ Email: email }
  if err := db.First(&target).Error; err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, email, target.Email)
  assert.Equal(t, "THISIS TEST", target.Name)
  assert.Equal(t, false, target.IsStudent)
}

func TestGetByEmail(t *testing.T) {
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		panic(err)
	}

  fetched, err := impl.User.GetByEmail(email)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, email, fetched.Email)
  assert.Equal(t, "TEST DATA", fetched.Name)
  assert.Equal(t, false, fetched.IsStudent)
}
