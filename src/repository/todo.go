package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type ToDoRepository interface {
	CreateToDo(event domain.ToDo) (domain.ToDo, error)
	UpdateToDo(event domain.ToDo) (domain.ToDo, error)
	DeleteToDo(id domain.ID) (domain.ToDo, error)
	FetchMonthlyToDo(email domain.Email, year int, month int) ([]domain.ToDo, error)
}
