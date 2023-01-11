package api

import (
	"errors"
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository"
	"nitic-pbl-2022-01/pbl-back/src/response"

	"github.com/gin-gonic/gin"
)

func CreateTagStudent(c *gin.Context) {
  type reqBody struct {
    Name string
    Color string
    AdminEmail []string
    MemberEmail []string
  }

  var body reqBody
  if err := c.BindJSON(&body); err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  admin := []domain.User{}
  member := []domain.User{}

  for _, a := range body.AdminEmail {
    email, err := domain.ConstructEmail(a)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }
    u, err := repository.User.GetByEmail(email)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }

    admin = append(admin, u)
  }

  for _, m := range body.MemberEmail {
    email, err := domain.ConstructEmail(m)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }
    u, err := repository.User.GetByEmail(email)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }

    member = append(member, u)
  }

  id, err := domain.GenerateID()
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  tag := domain.ConstructTag(id, body.Name, body.Color, admin, member, domain.None)

  created, err := repository.Tag.CreateTag(tag)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  resJson := response.ConvertTag(created)

  c.JSON(201, resJson)
}

func CreateTagTeacher(c *gin.Context) {
  type reqBody struct {
    Name string
    Color string
    AdminEmail []string
    MemberEmail []string
    Type string
  }

  var body reqBody
  if err := c.BindJSON(&body); err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  admin := []domain.User{}
  member := []domain.User{}

  for _, a := range body.AdminEmail {
    email, err := domain.ConstructEmail(a)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }
    u, err := repository.User.GetByEmail(email)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }

    admin = append(admin, u)
  }

  for _, m := range body.MemberEmail {
    email, err := domain.ConstructEmail(m)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }
    u, err := repository.User.GetByEmail(email)
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }

    member = append(member, u)
  }

  id, err := domain.GenerateID()
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  tag := domain.ConstructTag(id, body.Name, body.Color, admin, member, domain.TagTypeFromString(body.Type))

  created, err := repository.Tag.CreateTag(tag)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  resJson := response.ConvertTag(created)

  c.JSON(201, resJson)
}

func GetTags(c *gin.Context) {
  // FIXME: get from token
  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  tags, err := repository.Tag.GetByUserEmail(email)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  res := []response.Tag{}

  for _, t := range tags {
    res = append(res, response.ConvertTag(t))
  }

  c.JSON(200, res)
}

func UpdateTag(c *gin.Context) {
  type reqBody struct {
    Name string
    Color string
  }

  var body reqBody
  if err := c.BindJSON(&body); err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  id := c.Param("id")

  tag, err := repository.Tag.GetByID(domain.ID(id))
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  tag.Name = body.Name
  tag.Color = body.Color

  edited, err := repository.Tag.EditTag(tag)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  c.JSON(200, response.ConvertTag(edited))
}

func AddMemberToTag(c *gin.Context) {
  type reqBody struct {
    Email string
    Role string
  }

  var body reqBody
  if err := c.BindJSON(&body); err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  id := c.Param("id")

  email, err := domain.ConstructEmail(body.Email)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  user, err := repository.User.GetByEmail(email)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  var tag domain.Tag

  switch body.Role {
  case "ADMIN":
    c.Error(errors.New("Not implemented")).SetType(gin.ErrorTypePublic)
    break
  case "MEMBER":
    tag, err = repository.Tag.AddMemberToTag(domain.ID(id), []domain.User{user})
    if err != nil {
      c.Error(err).SetType(gin.ErrorTypePublic)
      return
    }
  default:
    c.Error(errors.New("Role does not exist")).SetType(gin.ErrorTypePublic)
    break
  }

  c.JSON(200, response.ConvertTag(tag))
}

func LeaveFromTag(c *gin.Context) {
  // FIXME: get from token
  email, err := domain.ConstructEmail("test-data@example.com")
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  user, err := repository.User.GetByEmail(email)
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  id := c.Param("id")

  tag, err := repository.Tag.DeleteMemberFromTag(domain.ID(id), []domain.User{user})
  if err != nil {
    c.Error(err).SetType(gin.ErrorTypePublic)
    return
  }

  c.JSON(200, response.ConvertTag(tag))
}
