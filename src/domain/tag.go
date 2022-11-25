package domain

type TagType int

const (
  HR = iota
  Class
  None
)

func (t TagType) Parse() string {
  switch (t) {
    case HR:
      return "HR"
    case Class:
      return "Class"
    case None:
      return "None"
    default:
      return "UNREACHEABLE"
  }
}

type Tag struct {
  ID ID
  Name string
  Color string
  Admin []User
  Member []User
  Type TagType
}

func ConstructTag(id ID, name string, color string, admin []User, member []User, t TagType) Tag {
  return Tag {
    ID: id,
    Name: name,
    Color: color,
    Admin: admin,
    Member: member,
    Type: t,
  }
}
