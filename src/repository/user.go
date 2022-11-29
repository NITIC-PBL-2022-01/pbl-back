package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type UserRepository interface {
  GetByEmail(email domain.Email) (domain.User, error)
  Create(newobj domain.User) (domain.User, error)
}
