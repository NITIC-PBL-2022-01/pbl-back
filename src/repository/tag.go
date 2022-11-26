package repository

import "nitic-pbl-2022-01/pbl-back/src/domain"

type TagRepository interface {
  CreateTag(tag domain.Tag) (domain.Tag, error)
  EditTag(tag domain.Tag) (domain.Tag, error)
  AddMemberToTag(id domain.ID, users []domain.User) (domain.Tag, error)
  GetByID(id domain.ID) (domain.Tag, error)
  GetByUserID(id domain.ID) ([]domain.Tag, error)
  GetModeratingTags(id domain.ID) ([]domain.Tag, error)
  DeleteTag(id domain.ID) (domain.Tag, error)
}
