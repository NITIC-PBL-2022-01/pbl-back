package impl_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
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

  tag := domain.Tag{ ID: "0" }

  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    t.Fatal(err)
  }

  user := domain.User{ Email: email }

  db.First(&tag);
  db.First(&user);

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

  created, err := impl.Event.CreateEvent(event)
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
