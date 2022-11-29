package mock_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
  email, err := domain.ConstructEmail("hoge@example.com")
  if err != nil {
    t.Fatal(err)
  }
  user := domain.User{
    Email: email,
    Name: "山田 太郎",
    IsStudent: true,
  }
  created, err := mock.User.Create(user)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, user.Email, created.Email)
  assert.Equal(t, user.Name, created.Name)
  assert.Equal(t, user.IsStudent, created.IsStudent)

  _, err = mock.User.Create(user)
  if err == nil {
    t.Fatal("error must occured here")
  }
}

func TestGetByEmail(t *testing.T) {
  email, err := domain.ConstructEmail("fuga@example.com")
  if err != nil {
    t.Fatal(err)
  }
  user := domain.User{
    Email: email,
    Name: "山田 太郎2",
    IsStudent: true,
  }
  _, err = mock.User.Create(user)
  if err != nil {
    t.Fatal(err)
  }

  found, err := mock.User.GetByEmail(email)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, user.Email, found.Email)
  assert.Equal(t, user.Name, found.Name)
  assert.Equal(t, user.IsStudent, found.IsStudent)

  invalid, err := domain.ConstructEmail("not.found@example.com")
  if err != nil {
    t.Fatal(err)
  }

  _, err = mock.User.GetByEmail(invalid)
  if err == nil {
    t.Fatal("error must occured here")
  }
}
