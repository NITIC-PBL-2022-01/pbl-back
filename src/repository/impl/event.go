package impl

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"

	"gorm.io/gorm"
)

type EventRepository struct {
  db *gorm.DB
}

func (r *EventRepository) CreateEvent(event domain.Event) (domain.Event, error) {
  result := r.db.Create(&event)
  if result.Error != nil {
    return domain.Event{}, result.Error
  }

  return event, nil
}

func (r *EventRepository) UpdateEvent(event domain.Event) (domain.Event, error) {
  if err := r.db.Model(&event).Updates(event).Error; err != nil {
    return domain.Event{}, err
  }

  return event, nil
}

func (r *EventRepository) DeleteEvent(id domain.ID) (domain.Event, error) {
  event := domain.Event{}
  if err := r.db.Delete(&event, id).Error; err != nil {
    return domain.Event{}, err
  }

  return event, nil
}

func (r *EventRepository) FetchMonthlyEvent(email domain.Email, year int, month int) ([]domain.Event, error) {
  user := domain.User{}

  if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
    return []domain.Event{}, err
  }

  all := []domain.Tag{}
  if err := r.db.Find(&all).Error; err != nil {
    return []domain.Event{}, err
  }

  tags := []domain.ID{}

  for _, t := range all {
    for _, a := range t.Admin {
      if a.Email == email {
        tags = append(tags, t.ID)
      }
    }

    for _, m := range t.Member {
      if m.Email == email {
        tags = append(tags, t.ID)
      }
    }
  }

  events := []domain.Event{}
  if err := r.db.Where("tag_id IN ?", tags).Error; err != nil {
    return []domain.Event{}, err
  }

  return events, nil
}
