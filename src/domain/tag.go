package domain

import "gorm.io/gorm"

type TagType int

const (
	HR = iota
	Class
	None
)

func (t TagType) Parse() string {
	switch t {
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

func TagTypeFromString(str string) TagType {
  switch str {
  case "HR":
    return HR
  case "Class":
    return Class
  case "None":
    return None
  default:
    return None
  }
}

type Tag struct {
	gorm.Model
	ID     ID `gorm:"primaryKey"`
	Name   string
	Color  string
	Admin  []User `gorm:"many2many:user_admins;"`
	Member []User `gorm:"many2many:user_members;"`
	Type   TagType
}

func ConstructTag(id ID, name string, color string, admin []User, member []User, t TagType) Tag {
	return Tag{
		ID:     id,
		Name:   name,
		Color:  color,
		Admin:  admin,
		Member: member,
		Type:   t,
	}
}
