package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type TagRepository interface {
  CreateTag(tag domain.Tag) (domain.Tag, error)
  EditTag(tag domain.Tag) (domain.Tag, error)
  AddMemberToTag(id domain.ID, users []domain.User) (domain.Tag, error)
  DeleteMemberFromTag(id domain.ID, users []domain.User) (domain.Tag, error)
  GetByID(id domain.ID) (domain.Tag, error)
  GetByUserEmail(email domain.Email) ([]domain.Tag, error)
  GetModeratingTags(email domain.Email) ([]domain.Tag, error)
  DeleteTag(id domain.ID) (domain.Tag, error)
}
