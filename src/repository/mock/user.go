package mock

import (
	"errors"
	"nitic-pbl-2022-01/pbl-back/src/domain"
)

type UserRepository struct {
  users []domain.User
}

func (repo *UserRepository) GetByEmail(email domain.Email) (domain.User, error) {
  for _, v := range repo.users {
    if (v.Email == email) {
      return v, nil
    }
  }

  return domain.User{}, errors.New("User not found")
}

func (repo *UserRepository) Create(newobj domain.User) (domain.User, error) {
  for _, v := range repo.users {
    if (v.Email == newobj.Email) {
      return domain.User{}, errors.New("User already exist")
    }
  }

  repo.users = append(repo.users, newobj)

  return newobj, nil
}
