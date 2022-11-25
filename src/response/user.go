package response

type user struct {
  Email string `json:"email"`
  Name string `json:"name"`
  IsStudent bool `json:"is_student"`
}

// TODO: receive domain object for User
func ParseToUser() user {
  return user {
    Email: "",
    Name: "",
    IsStudent: false,
  }
}
