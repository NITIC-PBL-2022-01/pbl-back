package mock

import (
	"errors"
	"nitic-pbl-2022-01/pbl-back/src/domain"
)

type TagRepository struct {
  tags []domain.Tag
}

func (repo *TagRepository) CreateTag(tag domain.Tag) (domain.Tag, error) {
  return domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) EditTag(tag domain.Tag) (domain.Tag, error) {
  return domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) AddMemberToTag(id domain.ID, users []domain.User) (domain.Tag, error) {
  return domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) DeleteMemberFromTag(id domain.ID, users []domain.User) (domain.Tag, error) {
  return domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) GetByID(id domain.ID) (domain.Tag, error) {
  return domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) GetByUserEmail(email domain.Email) ([]domain.Tag, error) {
  return []domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) GetModeratingTags(id domain.ID) ([]domain.Tag, error) {
  return []domain.Tag{}, errors.New("Not implemented")
}

func (repo *TagRepository) DeleteTag(id domain.ID) (domain.Tag, error) {
  return domain.Tag{}, errors.New("Not implemented")
}
