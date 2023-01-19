package api

import (
	"log"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository"
	"nitic-pbl-2022-01/pbl-back/src/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAttendance(c *gin.Context) {
	email, err := getEmail(c)
	if err != nil {
		log.Println(err)
		return
	}

	type reqBody struct {
    Reason string
		Date    string
		Period  int
		ClassID string
	}

	var body reqBody
	if err := c.BindJSON(&body); err != nil {
		handleError(c, err)
		return
	}

	tag, err := repository.Tag.GetByID(domain.ID(body.ClassID))
	if err != nil {
		handleError(c, err)
		return
	}

	user, err := repository.User.GetByEmail(email)
	if err != nil {
		handleError(c, err)
		return
	}

	id, err := domain.GenerateID()
	if err != nil {
		handleError(c, err)
		return
	}

	d, err := strconv.Atoi(body.Date)
	if err != nil {
		handleError(c, err)
		return
	}

	attendance := domain.ConstructAttendance(id, time.Unix(int64(d), 0), body.Period, body.Reason, tag.ID, tag, user.Email, user)
	created, err := repository.Attendance.CreateAttendance(attendance)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(201, response.ConvertAttendance(created))
}

func FetchAttendance(c *gin.Context) {
	email, err := getEmail(c)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := repository.User.GetByEmail(email)
	if err != nil {
		handleError(c, err)
		return
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		handleError(c, err)
		return
	}

	if user.IsStudent {
		attendances, err := repository.Attendance.FetchAttendanceByYear(year, email)
		if err != nil {
			handleError(c, err)
			return
		}

		c.JSON(200, fold(attendances, response.ConvertAttendance))
	} else {
		attendances, err := repository.Attendance.FetchModeratingAttendance(year, email, repository.Tag)
		if err != nil {
			handleError(c, err)
			return
		}

		c.JSON(200, fold(attendances, response.ConvertAttendance))
	}
}
