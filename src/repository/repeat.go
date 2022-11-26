package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type RepeatRepository interface {
  CreateRepeat(repeat domain.Repeat) (domain.Repeat, error)
  EditRepeat(repeat domain.Repeat) (domain.Repeat, error)
  DeleteRepeat(repeat domain.Repeat) (domain.Repeat, error)
}
