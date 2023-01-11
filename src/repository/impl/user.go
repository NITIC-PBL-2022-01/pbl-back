package impl

import (
	"errors"
	"nitic-pbl-2022-01/pbl-back/src/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
  db *gorm.DB
}

func (r *UserRepository) GetByEmail(email domain.Email) (domain.User, error) {
  return domain.User{}, errors.New("not implemented")
}

func (r *UserRepository) Create(newobj domain.User) (domain.User, error) {
  return domain.User{}, errors.New("not implemented")
}
