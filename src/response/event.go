package response

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"strconv"
)

type Event struct {
  ID string `json:"id"`
  Date string `json:"date"`
  Duration string `json:"duration"`
  Title string `json:"title"`
  Detail string `json:"detail"`
  Tag Tag `json:"tag"`
  Color string `json:"color"`
  AuthorName string `json:"author_name"`
  Location string `json:"location"`
  Repeat Repeat `json:"repeat"`
  IsToDo bool `json:"is_todo"`
  Done bool `json:"done"`
}

func ConvertEvent(event domain.Event) Event {
  return Event {
    ID: string(event.ID),
    Date: strconv.Itoa(int(event.Date.Unix())),
    Duration: strconv.Itoa(int(event.Duration.Seconds())),
    Title: event.Title,
    Detail: event.Detail,
    Tag: ConvertTag(event.Tag),
    Color: event.Color,
    AuthorName: event.AuthorName,
    Location: event.Location,
    Repeat: ConvertRepeat(event.Repeat),
    IsToDo: false,
    Done: false,
  }
}

func ConvertTodo(event domain.ToDo) Event {
  return Event {
    ID: string(event.ID),
    Date: strconv.Itoa(int(event.Date.Unix())),
    Duration: strconv.Itoa(int(event.Duration.Seconds())),
    Title: event.Title,
    Detail: event.Detail,
    Tag: ConvertTag(event.Tag),
    Color: event.Color,
    AuthorName: event.AuthorName,
    Location: event.Location,
    Repeat: ConvertRepeat(event.Repeat),
    IsToDo: event.IsToDo,
    Done: event.IsToDo,
  }
}
