package response

import "nitic-pbl-2022-01/pbl-back/src/domain"

type Tag struct {
  ID string
  Name string
  Color string
  Admin []User
  Member []User
  Type string
}

func ConvertTag(tag domain.Tag) Tag {
  return Tag {
    ID: string(tag.ID),
    Name: tag.Name,
    Color: tag.Color,
    Admin: convertArray(tag.Admin, ConvertUser),
    Member: convertArray(tag.Member, ConvertUser),
    Type: tag.Type.Parse(),
  }
}
