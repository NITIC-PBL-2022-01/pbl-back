package api

import (
	"fmt"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository"
	"nitic-pbl-2022-01/pbl-back/src/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	// FIXME: get from token
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	type reqBody struct {
		Title    string
		Detail   string
		TagID    string
		Color    string
		Duration int
		IsToDo   bool
		Done     bool
		Location string
		IsRepeat bool
		Repeat   struct {
			Unit  string
			Value int
			Since string
			Until string
		}
	}

	var body reqBody
	if err := c.BindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	id, err := domain.GenerateID()
	if err != nil {
		handleError(c, err)
		return
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		handleError(c, err)
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		handleError(c, err)
		return
	}

	date, err := strconv.Atoi(c.Param("date"))
	if err != nil {
		handleError(c, err)
		return
	}

	eventDate := time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.UTC)
	duration, err := time.ParseDuration(fmt.Sprintf("%dm", body.Duration))
	if err != nil {
		handleError(c, err)
		return
	}

	tag, err := repository.Tag.GetByID(domain.ID(body.TagID))
	if err != nil {
		handleError(c, err)
		return
	}

	author, err := repository.User.GetByEmail(email)
	if err != nil {
		handleError(c, err)
		return
	}

	var repeat domain.Repeat
	if body.IsRepeat {
		since, err := strEpochToTime(body.Repeat.Since)
		if err != nil {
			handleError(c, err)
			return
		}

		until, err := strEpochToTime(body.Repeat.Until)
		if err != nil {
			handleError(c, err)
			return
		}
		repeat = domain.ConstructRepeat(repeat.Unit, repeat.Value, since, until)
	}

	if body.IsToDo {
		// TODO: TODOのrepository作ったら書く
		domain.ConstructTodo(
			id,
			eventDate,
			duration,
			body.Title,
			body.Detail,
			tag,
			body.Color,
			author,
			body.Location,
			repeat,
			body.IsToDo,
			body.Done,
		)

		c.JSON(501, map[string]string{"message": "Not implemented"})
	} else {
		event := domain.ConstructEvent(
			id,
			eventDate,
			duration,
			body.Title,
			body.Detail,
			tag,
			body.Color,
			author,
			body.Location,
			repeat,
		)

		created, err := repository.Event.CreateEvent(event)
		if err != nil {
			handleError(c, err)
			return
		}

		c.JSON(201, response.ConvertEvent(created))
	}
}

func GetEventWithMonth(c *gin.Context) {
	// FIXME: get from token
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		handleError(c, err)
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		handleError(c, err)
		return
	}

	events, err := repository.Event.FetchMonthlyEvent(email, year, month)
	if err != nil {
		handleError(c, err)
		return
	}

	// TODO: add ToDo

	c.JSON(200, fold(events, response.ConvertEvent))
}

func UpdateEvent(c *gin.Context) {
	type reqBody struct {
		Title    string
		Detail   string
		Date     string
		TagID    string
		Color    string
		Duration int
		IsToDo   bool
		Done     bool
		Location string
		IsRepeat bool
		Repeat   struct {
			Unit  string
			Value int
			Since string
			Until string
		}
	}

	var body reqBody
	if err := c.BindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	id := c.Param("id")

	var date time.Time
	if body.Date != "" {
		var err error
		date, err = strEpochToTime(body.Date)
		if err != nil {
			handleError(c, err)
		}
	}

	var duration time.Duration
	if body.Duration != 0 {
		var err error
		duration, err = time.ParseDuration(fmt.Sprintf("%dm", body.Duration))
		if err != nil {
			handleError(c, err)
			return
		}
	}

	var tag domain.Tag
	if body.TagID != "" {
		var err error
		tag, err = repository.Tag.GetByID(domain.ID(body.TagID))
		if err != nil {
			handleError(c, err)
			return
		}
	}

	var repeat domain.Repeat
	if body.IsRepeat && body.Repeat.Since != "" && body.Repeat.Unit != "" && body.Repeat.Value != 0 &&
		body.Repeat.Until != "" {
		since, err := strEpochToTime(body.Repeat.Since)
		if err != nil {
			handleError(c, err)
			return
		}

		until, err := strEpochToTime(body.Repeat.Until)
		if err != nil {
			handleError(c, err)
			return
		}
		repeat = domain.ConstructRepeat(repeat.Unit, repeat.Value, since, until)
	}

	if body.IsToDo {
		// TODO: repository needed
		c.JSON(501, map[string]string{"message": "not implemented"})
	} else {
		newObj := domain.Event{
			EventBase: domain.EventBase{
				ID:       domain.ID(id),
				Date:     date,
				Duration: duration,
				Title:    body.Title,
				Detail:   body.Detail,
				Tag:      tag,
				Color:    body.Color,
				Location: body.Location,
				Repeat:   repeat,
			},
		}
		edited, err := repository.Event.UpdateEvent(newObj)
		if err != nil {
			handleError(c, err)
			return
		}

		c.JSON(200, response.ConvertEvent(edited))
	}
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := repository.Event.DeleteEvent(domain.ID(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(200, response.ConvertEvent(event))
}
