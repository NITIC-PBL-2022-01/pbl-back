package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type AttendanceRepository interface {
	CreateAttendance(attendance domain.Attendance) (domain.Attendance, error)
	FetchAttendanceByYear(year int, email domain.Email) ([]domain.Attendance, error)
	FetchModeratingAttendance(year int, email domain.Email, tagRepo TagRepository) ([]domain.Attendance, error)
}
