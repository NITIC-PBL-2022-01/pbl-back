package impl_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var attendanceID domain.ID

func TestCreateAttendance(t *testing.T) {
	var err error
	attendanceID, err = domain.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)

	tag := domain.Tag{ID: "0"}
	db.First(&tag)

	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	user, err := impl.User.GetByEmail(email)
	if err != nil {
		t.Fatal(err)
	}

	attendance := domain.ConstructAttendance(attendanceID, date, 1, tag.ID, tag, email, user)
	created, err := impl.Attendance.CreateAttendance(attendance)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, attendanceID, created.ID)
	assert.Equal(t, domain.ID("0"), created.Tag.ID)
	assert.Equal(t, 1, created.Period)
	assert.Equal(t, date, created.Date)
	assert.Equal(t, email, created.User.Email)
}

func TestFetchAttendanceByYear(t *testing.T) {
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	attendances, err := impl.Attendance.FetchAttendanceByYear(2022, email)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(attendances))
	assert.Equal(t, domain.ID("0"), attendances[0].Tag.ID)
}

func TestFetchModeratingAttendance(t *testing.T) {
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	attendances, err := impl.Attendance.FetchModeratingAttendance(2022, email, impl.Tag)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, len(attendances))
	assert.Equal(t, domain.ID("0"), attendances[0].Tag.ID)
}
