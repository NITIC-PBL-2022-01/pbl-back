package domain

import "gorm.io/gorm"

type User struct {
  gorm.Model
  ID ID `gorm:"primaryKey"`
  Email Email
  Name string
  IsStudent bool
}

func ConstructUser(email Email, name string, isStudent bool) User {
  return User {
    Email: email,
    Name: name,
    IsStudent: isStudent,
  }
}
