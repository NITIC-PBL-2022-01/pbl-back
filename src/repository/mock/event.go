package mock

import (
	"errors"
	"nitic-pbl-2022-01/pbl-back/src/domain"
)

type EventRepository struct {
  events []domain.Event
}

func (repo *EventRepository) CreateEvent(event domain.Event) (domain.Event, error) {
  for _, v := range repo.events {
    if v.ID == event.ID {
      return domain.Event{}, errors.New("Event ID is duplicated")
    }
  }

  repo.events = append(repo.events, event)

  return event, nil
}

func (repo *EventRepository) UpdateEvent(event domain.Event) (domain.Event, error) {
  for i, v := range repo.events {
    if v.ID == event.ID {
      repo.events[i].Date = event.Date
      repo.events[i].Duration = event.Duration
      repo.events[i].Title = event.Title
      repo.events[i].Detail = event.Detail
      repo.events[i].Tag = event.Tag
      repo.events[i].Color = event.Color
      repo.events[i].Author = event.Author
      repo.events[i].Location = event.Location
      repo.events[i].Repeat = event.Repeat

      return repo.events[i], nil
    }
  }

  return domain.Event{}, errors.New("Event ID is not found")
}

func (repo *EventRepository) DeleteEvent(id domain.ID) (domain.Event, error) {
  for i, e := range repo.events {
    if e.ID == id {
      repo.events = append(repo.events[:i], repo.events[i+1:]...)

      return e, nil
    }
  }

  return domain.Event{}, errors.New("Event ID is not found")
}

func inTheTag(email domain.Email, tag domain.Tag) bool {
  for _, a := range tag.Admin {
    if a.Email == email {
      return true
    }
  }
  for _, m := range tag.Member {
    if m.Email == email {
      return true
    }
  }

  return false
}

func (repo *EventRepository) FetchMonthlyEvent(email domain.Email, year int, month int) ([]domain.Event, error) {
  events := []domain.Event{}

  for _, e := range repo.events {
    if e.Date.Year() == year && int(e.Date.Month()) == month && inTheTag(email, e.Tag) {
      events = append(events, e)
    }
  }

  return events, nil
}
