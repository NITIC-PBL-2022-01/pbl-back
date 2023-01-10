package mock_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/mock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
  id, err := domain.GenerateID()
  if err != nil {
    t.Fatal(err)
  }
  date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
  duration := time.Minute * 60

  tag, err := mock.Tag.GetByID(domain.ID("0"))
  if err != nil {
    t.Fatal(err)
  }

  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    t.Fatal(err)
  }

  user := domain.User{
    Email: email,
    Name: "山田 太郎",
    IsStudent: true,
  }

  event := domain.ConstructEvent(
    id,
    date,
    duration,
    "Title",
    "Detail",
    tag,
    "#000",
    user,
    "1号館",
    domain.ConstructRepeat(
      domain.Week,
      1,
      time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
      time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
    ),
  )

  created, err := mock.Event.CreateEvent(event)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, created.ID, event.ID)
  assert.Equal(t, created.Date, event.Date)
  assert.Equal(t, created.Duration, event.Duration)
  assert.Equal(t, created.Title, event.Title)
  assert.Equal(t, created.Detail, event.Detail)
  assert.Equal(t, created.Tag.ID, event.Tag.ID)
  assert.Equal(t, created.Color, event.Color)
  assert.Equal(t, created.Author.Email, event.Author.Email)
  assert.Equal(t, created.Location, event.Location)
  assert.Equal(t, created.Repeat.Value, event.Repeat.Value)
}

func TestUpdateEvent(t *testing.T) {
  date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
  duration := time.Minute * 60

  tag, err := mock.Tag.GetByID(domain.ID("0"))
  if err != nil {
    t.Fatal(err)
  }

  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    t.Fatal(err)
  }

  user := domain.User{
    Email: email,
    Name: "山田 太郎",
    IsStudent: true,
  }

  event := domain.ConstructEvent(
    domain.ID("0"),
    date,
    duration,
    "Title",
    "Detail",
    tag,
    "#000",
    user,
    "1号館",
    domain.ConstructRepeat(
      domain.Week,
      1,
      time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
      time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
    ),
  )

  edited, err := mock.Event.UpdateEvent(event)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, edited.ID, event.ID)
  assert.Equal(t, edited.Date, event.Date)
  assert.Equal(t, edited.Duration, event.Duration)
  assert.Equal(t, edited.Title, event.Title)
  assert.Equal(t, edited.Detail, event.Detail)
  assert.Equal(t, edited.Tag.ID, event.Tag.ID)
  assert.Equal(t, edited.Color, event.Color)
  assert.Equal(t, edited.Author.Email, event.Author.Email)
  assert.Equal(t, edited.Location, event.Location)
  assert.Equal(t, edited.Repeat.Value, event.Repeat.Value)
}

func TestDeleteEvent(t *testing.T) {
  id, err := domain.GenerateID()
  if err != nil {
    t.Fatal(err)
  }
  date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
  duration := time.Minute * 60

  tag, err := mock.Tag.GetByID(domain.ID("0"))
  if err != nil {
    t.Fatal(err)
  }

  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    t.Fatal(err)
  }

  user := domain.User{
    Email: email,
    Name: "山田 太郎",
    IsStudent: true,
  }

  event := domain.ConstructEvent(
    id,
    date,
    duration,
    "Title",
    "Detail",
    tag,
    "#000",
    user,
    "1号館",
    domain.ConstructRepeat(
      domain.Week,
      1,
      time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
      time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
    ),
  )

  created, err := mock.Event.CreateEvent(event)
  if err != nil {
    t.Fatal(err)
  }

  deleted, err := mock.Event.DeleteEvent(created.ID)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, deleted.ID, event.ID)
  assert.Equal(t, deleted.Date, event.Date)
  assert.Equal(t, deleted.Duration, event.Duration)
  assert.Equal(t, deleted.Title, event.Title)
  assert.Equal(t, deleted.Detail, event.Detail)
  assert.Equal(t, deleted.Tag.ID, event.Tag.ID)
  assert.Equal(t, deleted.Color, event.Color)
  assert.Equal(t, deleted.Author.Email, event.Author.Email)
  assert.Equal(t, deleted.Location, event.Location)
  assert.Equal(t, deleted.Repeat.Value, event.Repeat.Value)
}

func TestFetchMonthlyEvent(t *testing.T) {
  date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
  duration := time.Minute * 60

  tag, err := mock.Tag.GetByID(domain.ID("0"))
  if err != nil {
    t.Fatal(err)
  }

  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    t.Fatal(err)
  }

  user := domain.User{
    Email: email,
    Name: "山田 太郎",
    IsStudent: true,
  }

  event := domain.ConstructEvent(
    domain.ID("0"),
    date,
    duration,
    "Title",
    "Detail",
    tag,
    "#000",
    user,
    "1号館",
    domain.ConstructRepeat(
      domain.Week,
      1,
      time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
      time.Date(2022, 1, 31, 0, 0, 0, 0, time.Local),
    ),
  )

  events, err := mock.Event.FetchMonthlyEvent(email, 2022, 2)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, events[0].ID, event.ID)
  assert.Equal(t, events[0].Date, event.Date)
  assert.Equal(t, events[0].Duration, event.Duration)
  assert.Equal(t, events[0].Title, event.Title)
  assert.Equal(t, events[0].Detail, event.Detail)
  assert.Equal(t, events[0].Tag.ID, event.Tag.ID)
  assert.Equal(t, events[0].Color, event.Color)
  assert.Equal(t, events[0].Author.Email, event.Author.Email)
  assert.Equal(t, events[0].Location, event.Location)
  assert.Equal(t, events[0].Repeat.Value, event.Repeat.Value)
}
