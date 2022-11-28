package mock

import "nitic-pbl-2022-01/pbl-back/src/domain"

var (
  User = UserRepository{}
  Tag = TagRepository{}
)

func init() {
  User.users = []domain.User{}
  Tag.tags = []domain.Tag{}
}

func ConstructTestData() {
  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    panic(err)
  }
  User.users = append(User.users, domain.ConstructUser(
    email,
    "山田 太郎",
    false,
  ))
  Tag.tags = append(Tag.tags, domain.ConstructTag(
    domain.ID("0"),
    "hoge",
    "#fff",
    []domain.User{User.users[0]},
    []domain.User{},
    domain.None,
  ))
}
