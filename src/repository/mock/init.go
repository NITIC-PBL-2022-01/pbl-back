package mock

import "nitic-pbl-2022-01/pbl-back/src/domain"

var (
  User = UserRepository{}
)

func init() {
  User.users = []domain.User{}
}
