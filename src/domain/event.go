package domain

import (
	"time"

	"gorm.io/gorm"
)

type EventBase struct {
	gorm.Model
	ID       ID `gorm:"primaryKey"`
	Date     time.Time
	Duration time.Duration
	Title    string
	Detail   string
	Tag      Tag
	TagID    ID
	Color    string
	Author   User
	AuthorID ID
	Location string
	Repeat   Repeat
	RepeatID ID
}

type Event struct {
	EventBase
}

type ToDo struct {
	EventBase
	IsToDo bool
	Done   bool
}

func ConstructEvent(
	id ID,
	date time.Time,
	duration time.Duration,
	title string,
	detail string,
	tag Tag,
	color string,
	author User,
	location string,
	repeat Repeat,
) Event {
	return Event{
		EventBase{
			ID:       id,
			Date:     date,
			Duration: duration,
			Title:    title,
			Detail:   detail,
			Tag:      tag,
			Color:    color,
			Author:   author,
			Location: location,
			Repeat:   repeat,
		},
	}
}

func ConstructTodo(
	id ID,
	date time.Time,
	duration time.Duration,
	title string,
	detail string,
	tag Tag,
	color string,
	author User,
	location string,
	repeat Repeat,
	isToDo bool,
	done bool,
) ToDo {
	return ToDo{
		EventBase: EventBase{
			ID:       id,
			Date:     date,
			Duration: duration,
			Title:    title,
			Detail:   detail,
			Tag:      tag,
			Color:    color,
			Author:   author,
			Location: location,
			Repeat:   repeat,
		},
		IsToDo: isToDo,
		Done:   done,
	}
}
