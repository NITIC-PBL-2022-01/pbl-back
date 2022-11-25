package domain

type User struct {
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
