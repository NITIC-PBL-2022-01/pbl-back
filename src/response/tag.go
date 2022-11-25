package response

import "nitic-pbl-2022-01/pbl-back/src/domain"

type Tag struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Color string `json:"color"`
  Admin []User `json:"admin"`
  Member []User `json:"member"`
  Type string `json:"type"`
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
