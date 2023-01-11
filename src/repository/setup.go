package repository

import (
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"nitic-pbl-2022-01/pbl-back/src/repository/mock"
)

var (
  Event EventRepository
  User UserRepository
  Tag TagRepository
)

func SetupRepository(isProd bool) {
  if isProd {
    Event = &impl.Event
    User = &impl.User
    Tag = &impl.Tag
  } else {
    Event = &mock.Event
    User = &mock.User
    Tag = &mock.Tag
  }
}
