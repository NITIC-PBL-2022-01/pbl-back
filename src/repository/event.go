package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type EventRepository interface {
  CreateEvent(event domain.Event) (domain.Event, error)
  UpdateEvent(event domain.Event) (domain.Event, error)
  DeleteEvent(event domain.Event) (domain.Event, error)
  FetchMonthlyEvent(email string, year int, month int) ([]domain.Event, error)
}
