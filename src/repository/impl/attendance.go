package impl

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func (r *AttendanceRepository) CreateAttendance(attendance domain.Attendance) (domain.Attendance, error) {
	if err := r.db.Create(&attendance).Error; err != nil {
		return domain.Attendance{}, err
	}

	created := domain.Attendance{ID: attendance.ID}
	if err := r.db.Preload("User").Preload("Tag").Preload("Tag.Admin").Preload("Tag.Member").First(&created).Error; err != nil {
		return domain.Attendance{}, err
	}

	return created, nil
}

func (r *AttendanceRepository) FetchAttendanceByYear(year int, email domain.Email) ([]domain.Attendance, error) {
	parseDate := func(year int, month int) time.Time {
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}
	parseStartDate := func(year int) time.Time {
		return parseDate(year, 1)
	}
	parseEndDate := func(year int) time.Time {
		return parseDate(year+1, 1)
	}

	attendances := []domain.Attendance{}

	if err := r.db.Preload("User").Preload("Tag").Preload("Tag.Admin").Preload("Tag.Member").Where("user_email = ? AND date BETWEEN ? AND ?", email, parseStartDate(year), parseEndDate(year)).Find(&attendances).Error; err != nil {
		return []domain.Attendance{}, err
	}

	return attendances, nil
}

func (r *AttendanceRepository) FetchModeratingAttendance(
	year int,
	email domain.Email,
	tagRepo TagRepository,
) ([]domain.Attendance, error) {
	tags, err := tagRepo.GetModeratingTags(email)
	if err != nil {
		return []domain.Attendance{}, err
	}

	tagID := []domain.ID{}

	for _, t := range tags {
		tagID = append(tagID, t.ID)
	}

	parseDate := func(year int, month int) time.Time {
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}
	parseStartDate := func(year int) time.Time {
		return parseDate(year, 1)
	}
	parseEndDate := func(year int) time.Time {
		return parseDate(year+1, 1)
	}

	attendances := []domain.Attendance{}

	if err := r.db.Preload("User").Preload("Tag").Preload("Tag.Admin").Preload("Tag.Member").Where("tag_id IN ? AND date BETWEEN ? AND ?", tagID, parseStartDate(year), parseEndDate(year)).Find(&attendances).Error; err != nil {
		return attendances, err
	}

	return attendances, nil
}
