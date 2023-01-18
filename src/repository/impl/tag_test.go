package impl_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/impl"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tagID domain.ID

func TestCreateTag(t *testing.T) {
	var err error
	tagID, err = domain.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	user := domain.User{Email: email}

	if err := db.First(&user).Error; err != nil {
		t.Fatal(err)
	}

	tag := domain.ConstructTag(tagID, "THIS IS TAG", "#000", []domain.User{user}, []domain.User{}, domain.Class)

	created, err := impl.Tag.CreateTag(tag)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, created.ID)
	assert.Equal(t, "THIS IS TAG", created.Name)
	assert.Equal(t, "#000", created.Color)
	assert.Equal(t, email, created.Admin[0].Email)
	assert.Equal(t, 0, len(created.Member))
	assert.Equal(t, domain.TagType(domain.Class), created.Type)

	fetched := domain.Tag{ID: tagID}

	if err := db.Preload("Admin").Preload("Member").First(&fetched).Error; err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, fetched.ID)
	assert.Equal(t, "THIS IS TAG", fetched.Name)
	assert.Equal(t, "#000", fetched.Color)
	assert.Equal(t, email, fetched.Admin[0].Email)
	assert.Equal(t, 0, len(fetched.Member))
	assert.Equal(t, domain.TagType(domain.Class), fetched.Type)
}

func TestEditTag(t *testing.T) {
	dist := domain.Tag{ID: tagID}
	if err := db.First(&dist).Error; err != nil {
		t.Fatal(err)
	}

	dist.Name = "Updated"

	edited, err := impl.Tag.EditTag(dist)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, edited.ID)
	assert.Equal(t, "Updated", edited.Name)
	assert.Equal(t, "#000", edited.Color)
	assert.Equal(t, 0, len(edited.Member))
	assert.Equal(t, domain.TagType(domain.Class), edited.Type)

	fetched := domain.Tag{ID: tagID}

	if err := db.Preload("Admin").Preload("Member").First(&fetched).Error; err != nil {
		t.Fatal(err)
	}

	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, fetched.ID)
	assert.Equal(t, "Updated", fetched.Name)
	assert.Equal(t, "#000", fetched.Color)
	assert.Equal(t, email, fetched.Admin[0].Email)
	assert.Equal(t, 0, len(fetched.Member))
	assert.Equal(t, domain.TagType(domain.Class), fetched.Type)
}

func TestAddMemberToTag(t *testing.T) {
	email, err := domain.ConstructEmail("tag@example.com")
	if err != nil {
		t.Fatal(err)
	}

	user := domain.User{Email: email}
	if err := db.First(&user).Error; err != nil {
		t.Fatal(err)
	}

	edited, err := impl.Tag.AddMemberToTag(tagID, []domain.User{user})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, edited.ID)
	assert.Equal(t, "Updated", edited.Name)
	assert.Equal(t, "#000", edited.Color)
	assert.Equal(t, 1, len(edited.Admin))
	assert.Equal(t, email, edited.Member[0].Email)
	assert.Equal(t, domain.TagType(domain.Class), edited.Type)

	fetched := domain.Tag{ID: tagID}

	if err := db.Preload("Admin").Preload("Member").First(&fetched).Error; err != nil {
		t.Fatal(err)
	}

	email2, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, fetched.ID)
	assert.Equal(t, "Updated", fetched.Name)
	assert.Equal(t, "#000", fetched.Color)
	assert.Equal(t, email2, fetched.Admin[0].Email)
	assert.Equal(t, email, fetched.Member[0].Email)
	assert.Equal(t, domain.TagType(domain.Class), fetched.Type)
}

func TestGetByUserEmail(t *testing.T) {
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	fetched, err := impl.Tag.GetByUserEmail(email)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(fetched))
	assert.NotEqual(t, fetched[0].ID, fetched[1].ID)

	email, err = domain.ConstructEmail("tag@example.com")
	if err != nil {
		t.Fatal(err)
	}

	fetched, err = impl.Tag.GetByUserEmail(email)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(fetched))
}

func TestGetModeratingTags(t *testing.T) {
	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	fetched, err := impl.Tag.GetModeratingTags(email)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(fetched))
	assert.NotEqual(t, fetched[0].ID, fetched[1].ID)

	email, err = domain.ConstructEmail("tag@example.com")
	if err != nil {
		t.Fatal(err)
	}

	fetched, err = impl.Tag.GetModeratingTags(email)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 0, len(fetched))
}

func TestDeleteMemberFromTag(t *testing.T) {
	email, err := domain.ConstructEmail("tag@example.com")
	if err != nil {
		t.Fatal(err)
	}

	user := domain.User{Email: email}
	if err := db.First(&user).Error; err != nil {
		t.Fatal(err)
	}

	edited, err := impl.Tag.DeleteMemberFromTag(tagID, []domain.User{user})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, edited.ID)
	assert.Equal(t, "Updated", edited.Name)
	assert.Equal(t, "#000", edited.Color)
	assert.Equal(t, 1, len(edited.Admin))
	assert.Equal(t, 0, len(edited.Member))
	assert.Equal(t, domain.TagType(domain.Class), edited.Type)

	fetched := domain.Tag{ID: tagID}

	if err := db.Preload("Admin").Preload("Member").First(&fetched).Error; err != nil {
		t.Fatal(err)
	}

	email2, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, fetched.ID)
	assert.Equal(t, "Updated", fetched.Name)
	assert.Equal(t, "#000", fetched.Color)
	assert.Equal(t, email2, fetched.Admin[0].Email)
	assert.Equal(t, 0, len(fetched.Member))
	assert.Equal(t, domain.TagType(domain.Class), fetched.Type)
}

func TestGetByID(t *testing.T) {
	fetched, err := impl.Tag.GetByID(tagID)
	if err != nil {
		t.Fatal(err)
	}

	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, fetched.ID)
	assert.Equal(t, "Updated", fetched.Name)
	assert.Equal(t, "#000", fetched.Color)
	assert.Equal(t, email, fetched.Admin[0].Email)
	assert.Equal(t, 0, len(fetched.Member))
	assert.Equal(t, domain.TagType(domain.Class), fetched.Type)
}

func TestDeleteTag(t *testing.T) {
	tag, err := impl.Tag.DeleteTag(tagID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tagID, tag.ID)

	dist := []domain.Tag{}
	db.Find(&dist)

	assert.Equal(t, 1, len(dist))
}
