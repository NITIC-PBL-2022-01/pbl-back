package domain

type User struct {
  Email string
  Name string
  IsStudent bool
}

func ConstructUser(email string, name string, isStudent bool) User {
  return User {
    Email: email,
    Name: name,
    IsStudent: isStudent,
  }
}
