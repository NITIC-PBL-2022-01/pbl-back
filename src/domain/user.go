package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
  Email Email `gorm:"primaryKey"`
  Name string
  IsStudent bool
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

func ConstructUser(email Email, name string, isStudent bool) User {
  return User {
    Email: email,
    Name: name,
    IsStudent: isStudent,
  }
}
