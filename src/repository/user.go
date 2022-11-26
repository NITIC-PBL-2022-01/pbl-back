package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type UserRepository interface {
  GetByID(email string) (domain.User, error)
  create(newobj domain.User) (domain.User, error)
}
