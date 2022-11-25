package response

import "nitic-pbl-2022-01/pbl-back/src/domain"

type User struct {
  Email string `json:"email"`
  Name string `json:"name"`
  IsStudent bool `json:"is_student"`
}

func ConvertUser(user domain.User) User {
  return User {
    Email: string(user.Email),
    Name: user.Name,
    IsStudent: user.IsStudent,
  }
}
