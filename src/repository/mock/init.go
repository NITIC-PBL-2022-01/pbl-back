package mock

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"time"
)

var (
  User = UserRepository{}
  Tag = TagRepository{}
  Event = EventRepository{}
)

func init() {
  User.users = []domain.User{}
  Tag.tags = []domain.Tag{}
  Event.events = []domain.Event{}
}

func ConstructTestData() {
  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    panic(err)
  }
  User.users = append(User.users, domain.ConstructUser(
    email,
    "山田 太郎",
    false,
  ))
  Tag.tags = append(Tag.tags, domain.ConstructTag(
    domain.ID("0"),
    "hoge",
    "#fff",
    []domain.User{User.users[0]},
    []domain.User{},
    domain.None,
  ))
  Event.events = append(Event.events, domain.ConstructEvent(
    domain.ID("0"),
    time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
    time.Minute * 90,
    "The event",
    "This event is created for test",
    Tag.tags[0],
    "#fff",
    User.users[0],
    "8号館",
    domain.ConstructRepeat(
      domain.Week,
      1,
      time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
      time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
    ),
  ))
}
