package impl

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"time"

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
	event.ID = id
	if err := r.db.First(&event).Error; err != nil {
		return domain.Event{}, err
	}
	dist := domain.Event{}
	dist.ID = id
	if err := r.db.Delete(&dist).Error; err != nil {
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
	if err := r.db.Preload("Admin").Preload("Member").Find(&all).Error; err != nil {
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

	nextMonth := func(month int) int {
		if month == 12 {
			return 1
		}
		return month + 1
	}

	nextYear := func(year int, month int) int {
		if month == 12 {
			return year
		}
		return year + 1
	}

	parseDate := func(year int, month int) time.Time {
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}

	events := []domain.Event{}
	if err := r.db.Joins("Repeat").Where("tag_id IN ? AND date BETWEEN ? AND ?", tags, parseDate(year, month), parseDate(nextYear(year, month), nextMonth(month))).Find(&events).Error; err != nil {
		return []domain.Event{}, err
	}

	repeatEvents := []domain.Event{}

	//  repeats.until > '?-?-01 00:00:00' AND AND repeats.since < '?-?-01 00:00:00'
	// , year, month, nextYear(year, month), nextMonth(month)
	if err := r.db.Joins("Repeat").Joins("JOIN repeats ON events.repeat_id = repeats.id").
		Where("tag_id IN ?", tags).
		Find(&repeatEvents).Error; err != nil {
		return []domain.Event{}, err
	}

	for _, e := range repeatEvents {
		duration := e.Repeat.Until.Sub(e.Repeat.Since)
		days := int(duration.Abs().Hours() / 24)
		monthBegin := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		nextMonthBegin := time.Date(nextYear(year, month), time.Month(nextMonth(month)), 1, 0, 0, 0, 0, time.UTC)
		isInclude := func(date time.Time) bool {
			return date.After(monthBegin) && date.Before(nextMonthBegin)
		}

		addDate := func(date time.Time, amount int, unit domain.RepeatUnit) time.Time {
			switch unit {
			case domain.Day:
				return date.AddDate(0, 0, amount)
			case domain.Week:
				return date.AddDate(0, 0, amount*7)
			case domain.Month:
				return date.AddDate(0, amount, 0)
			case domain.Year:
				return date.AddDate(amount, 0, 0)
			default:
				panic("unreachable")
			}
		}

		var count int
		switch e.Repeat.Unit {
		case domain.Day:
			count = days
		case domain.Week:
			count = days / 7
		case domain.Month:
			count = days / 30
		case domain.Year:
			count = days / 365
		default:
			count = days
		}
		cache := e
		for i := 0; i <= count; i++ {
			if isInclude(addDate(e.Date, i, e.Repeat.Unit)) {
				cache.Date = addDate(e.Date, i, e.Repeat.Unit)
				events = append(events, cache)
			}
		}
	}

	return events, nil
}
