package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type EventRepository interface {
	CreateEvent(event domain.Event) (domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	DeleteEvent(id domain.ID) (domain.Event, error)
	FetchMonthlyEvent(email domain.Email, year int, month int) ([]domain.Event, error)
}
