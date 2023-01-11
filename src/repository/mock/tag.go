package mock

import (
	"errors"
	"nitic-pbl-2022-01/pbl-back/src/domain"
)

type TagRepository struct {
	tags []domain.Tag
}

func (repo *TagRepository) CreateTag(tag domain.Tag) (domain.Tag, error) {
	if tag.Admin[0].IsStudent && (tag.Type == domain.Class || tag.Type == domain.HR) {
		return domain.Tag{}, errors.New("Student could not create Class and HR")
	}

	for _, v := range repo.tags {
		if v.ID == tag.ID {
			return domain.Tag{}, errors.New("ID is duplicated")
		}
	}

	repo.tags = append(repo.tags, tag)

	return tag, nil
}

func (repo *TagRepository) EditTag(tag domain.Tag) (domain.Tag, error) {
	if tag.Admin[0].IsStudent && (tag.Type == domain.Class || tag.Type == domain.HR) {
		return domain.Tag{}, errors.New("Student could not create Class and HR")
	}
	for i, v := range repo.tags {
		if v.ID == tag.ID {
			repo.tags[i].Name = tag.Name
			repo.tags[i].Color = tag.Color
			repo.tags[i].Type = tag.Type

			return tag, nil
		}
	}

	return domain.Tag{}, errors.New("Tag ID is not found")
}

func (repo *TagRepository) AddMemberToTag(id domain.ID, users []domain.User) (domain.Tag, error) {
	for i, v := range repo.tags {
		if v.ID == id {
			repo.tags[i].Member = append(repo.tags[i].Member, users...)

			return repo.tags[i], nil
		}
	}

	return domain.Tag{}, errors.New("Tag ID is not found")
}

func (repo *TagRepository) DeleteMemberFromTag(id domain.ID, users []domain.User) (domain.Tag, error) {
	for i, v := range repo.tags {
		if v.ID == id {
			for _, u := range users {
				for n, m := range v.Member {
					if u.Email == m.Email {
						repo.tags[i].Member = append(repo.tags[i].Member[:n], repo.tags[i].Member[n+1:]...)
					}
				}
			}

			return repo.tags[i], nil
		}
	}

	return domain.Tag{}, errors.New("Tag ID is not found")
}

func (repo *TagRepository) GetByID(id domain.ID) (domain.Tag, error) {
	for _, v := range repo.tags {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.Tag{}, errors.New("Tag ID is not found")
}

func (repo *TagRepository) GetByUserEmail(email domain.Email) ([]domain.Tag, error) {
	tags := []domain.Tag{}
	for _, v := range repo.tags {
		for _, a := range v.Admin {
			if a.Email == email {
				tags = append(tags, v)
			}
		}
		for _, m := range v.Member {
			if m.Email == email {
				tags = append(tags, v)
			}
		}
	}

	return tags, nil
}

func (repo *TagRepository) GetModeratingTags(email domain.Email) ([]domain.Tag, error) {
	tags := []domain.Tag{}
	for _, v := range repo.tags {
		for _, a := range v.Admin {
			if a.Email == email {
				tags = append(tags, v)
			}
		}
	}

	return tags, nil
}

func (repo *TagRepository) DeleteTag(id domain.ID) (domain.Tag, error) {
	for i, v := range repo.tags {
		if v.ID == id {
			repo.tags = append(repo.tags[:i], repo.tags[i+1:]...)
			return v, nil
		}
	}

	return domain.Tag{}, nil
}
