package impl

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
  db *gorm.DB
}

func (r *UserRepository) GetByEmail(email domain.Email) (domain.User, error) {
  found := domain.User{ Email: email }
  if err := r.db.First(&found).Error; err != nil {
    return domain.User{}, err
  }

  return found, nil
}

func (r *UserRepository) Create(newobj domain.User) (domain.User, error) {
  if err := r.db.Create(&newobj).Error; err != nil {
    return domain.User{}, err
  }

  return newobj, nil
}
