package domain

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	gorm.Model
	ID        ID `gorm:"primaryKey"`
	Date      time.Time
	Period    int
  Reason string
	TagID     ID
	Tag       Tag
	User      User
	UserEmail Email
}

func ConstructAttendance(
	id ID,
	date time.Time,
	period int,
  reason string,
	tagID ID,
	tag Tag,
	userEmail Email,
	user User,
) Attendance {
	return Attendance{
		ID:        id,
		Date:      date,
		Period:    period,
		TagID:     tagID,
		Tag:       tag,
		UserEmail: userEmail,
		User:      user,
	}
}
