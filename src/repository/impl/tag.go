package impl

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"

	"gorm.io/gorm"
)

type TagRepository struct {
  db *gorm.DB
}

func (r *TagRepository)	CreateTag(tag domain.Tag) (domain.Tag, error) {
  if err := r.db.Preload("Member").Preload("Admin").Create(&tag).Error; err != nil {
    return domain.Tag{}, err
  }
  return tag, nil
}

func (r *TagRepository)	EditTag(tag domain.Tag) (domain.Tag, error) {
	if err := r.db.Model(&tag).Updates(tag).Error; err != nil {
    return domain.Tag{}, err
  }

  return tag, nil
}

func (r *TagRepository)	AddMemberToTag(id domain.ID, users []domain.User) (domain.Tag, error) {
  if err := r.db.Model(&domain.Tag{ ID: id }).Association("Member").Append(&users); err != nil {
    return domain.Tag{}, err
  }

  tag := domain.Tag{ ID: id }
  if err := r.db.Preload("Member").Preload("Admin").First(&tag).Error; err != nil {
    return domain.Tag{}, err
  }

  return tag, nil
}

func (r *TagRepository)	DeleteMemberFromTag(id domain.ID, users []domain.User) (domain.Tag, error) {
  if err := r.db.Model(&domain.Tag{ ID: id }).Association("Member").Delete(users); err != nil {
    return domain.Tag{}, err
  }

  tag := domain.Tag{ ID: id }
  if err := r.db.Preload("Member").Preload("Admin").First(&tag).Error; err != nil {
    return domain.Tag{}, err
  }

  return tag, nil
}

func (r *TagRepository)	GetByID(id domain.ID) (domain.Tag, error) {
  tag := domain.Tag{ ID: id }
  if err := r.db.Preload("Member").Preload("Admin").First(&tag).Error; err != nil {
    return domain.Tag{}, err
  }

  return tag, nil
}

func (r *TagRepository)	GetByUserEmail(email domain.Email) ([]domain.Tag, error) {
  all := []domain.Tag{}
  if err := r.db.Preload("Member").Preload("Admin").Find(&all).Error; err != nil {
    return []domain.Tag{}, err
  }

  tags := []domain.Tag{}

  for _, t := range all {
		for _, a := range t.Admin {
			if a.Email == email {
				tags = append(tags, t)
			}
		}

		for _, m := range t.Member {
			if m.Email == email {
				tags = append(tags, t)
			}
		}
  }

  return tags, nil
}

func (r *TagRepository)	GetModeratingTags(email domain.Email) ([]domain.Tag, error) {
  all := []domain.Tag{}
  if err := r.db.Preload("Member").Preload("Admin").Find(&all).Error; err != nil {
    return []domain.Tag{}, err
  }

  tags := []domain.Tag{}

  for _, t := range all {
		for _, a := range t.Admin {
			if a.Email == email {
				tags = append(tags, t)
			}
		}
  }

  return tags, nil
}

func (r *TagRepository)	DeleteTag(id domain.ID) (domain.Tag, error) {
  dist, err := r.GetByID(id)
  if err != nil {
    return domain.Tag{}, err
  }

  if err := r.db.Delete(&dist).Error; err != nil {
    return domain.Tag{}, err
  }

  return dist, nil
}
