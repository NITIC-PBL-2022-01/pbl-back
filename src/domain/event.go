package domain

import "time"

type EventBase struct {
  ID ID
  Date time.Time
  Duration time.Duration
  Title string
  Detail string
  Tag Tag
  Color string
  AuthorName string
  Location string
  Repeat Repeat
}

type Event struct {
  EventBase
}

type ToDo struct {
  EventBase
  IsToDo bool
  Done bool
}

func ConstructEvent(
  id ID,
  date time.Time,
  duration time.Duration,
  title string,
  detail string,
  tag Tag,
  color string,
  authorName string,
  location string,
  repeat Repeat,
) Event {
  return Event {
    EventBase {
      ID: id,
      Date: date,
      Duration: duration,
      Title: title,
      Detail: detail,
      Tag: tag,
      Color: color,
      AuthorName: authorName,
      Location: location,
      Repeat: repeat,
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
  authorName string,
  location string,
  repeat Repeat,
  isToDo bool,
  done bool,
) ToDo {
  return ToDo {
    EventBase: EventBase {
      ID: id,
      Date: date,
      Duration: duration,
      Title: title,
      Detail: detail,
      Tag: tag,
      Color: color,
      AuthorName: authorName,
      Location: location,
      Repeat: repeat,
    },
    IsToDo: isToDo,
    Done: done,
  }
}
