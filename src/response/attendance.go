package response

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"strconv"
)

type Attendance struct {
	ID        string `json:"id"`
	Date      string `json:"date"`
	Period    int    `json:"period"`
	TagID     string `json:"class_id"`
	UserEmail string `json:"user_email"`
}

func ConvertAttendance(attendance domain.Attendance) Attendance {
	return Attendance{
		ID:        string(attendance.ID),
		Date:      strconv.Itoa(int(attendance.Date.Unix())),
		Period:    attendance.Period,
		TagID:     string(attendance.TagID),
		UserEmail: string(attendance.UserEmail),
	}
}
