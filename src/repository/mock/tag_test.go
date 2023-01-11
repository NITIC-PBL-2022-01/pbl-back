package mock_test

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"nitic-pbl-2022-01/pbl-back/src/repository/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	id, err := domain.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	user, err := mock.User.GetByEmail(email)
	if err != nil {
		t.Fatal(err)
	}

	tag := domain.ConstructTag(id, "hoge", "#fff", []domain.User{user}, []domain.User{}, domain.None)
	created, err := mock.Tag.CreateTag(tag)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, created.ID, id)
	assert.Equal(t, created.Name, "hoge")
	assert.Equal(t, created.Admin[0].Email, email)
	assert.Equal(t, len(created.Member), 0)
	assert.Equal(t, created.Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, created.Color, "#fff")
}

func TestGetByID(t *testing.T) {
	fetched, err := mock.Tag.GetByID(domain.ID("0"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fetched.ID, domain.ID("0"))
	assert.Equal(t, fetched.Name, "hoge")
	assert.Equal(t, fetched.Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(fetched.Member), 0)
	assert.Equal(t, fetched.Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, fetched.Color, "#fff")
}

func TestEditTag(t *testing.T) {
	fetched, err := mock.Tag.GetByID(domain.ID("0"))
	if err != nil {
		t.Fatal(err)
	}

	fetched.Name = "fuga"
	fetched.Color = "#000"
	fetched.Type = domain.Class

	edited, err := mock.Tag.EditTag(fetched)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, edited.ID, domain.ID("0"))
	assert.Equal(t, edited.Name, "fuga")
	assert.Equal(t, edited.Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(edited.Member), 0)
	assert.Equal(t, edited.Type.Parse(), domain.TagType(domain.Class).Parse())
	assert.Equal(t, edited.Color, "#000")

	edited.Name = "hoge"
	edited.Color = "#fff"
	edited.Type = domain.None

	_, err = mock.Tag.EditTag(edited)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEditTagMember(t *testing.T) {
	user := domain.User{
		Email:     domain.Email("foo@example.com"),
		Name:      "山田 一郎",
		IsStudent: true,
	}

	tag, err := mock.Tag.AddMemberToTag(domain.ID("0"), []domain.User{user})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tag.ID, domain.ID("0"))
	assert.Equal(t, tag.Name, "hoge")
	assert.Equal(t, tag.Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(tag.Member), 1)
	assert.Equal(t, tag.Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, tag.Color, "#fff")

	restored, err := mock.Tag.DeleteMemberFromTag(domain.ID("0"), []domain.User{user})
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, restored.ID, domain.ID("0"))
	assert.Equal(t, restored.Name, "hoge")
	assert.Equal(t, restored.Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(restored.Member), 0)
	assert.Equal(t, restored.Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, restored.Color, "#fff")
}

func TestGetByUserEmail(t *testing.T) {
	fetched, err := mock.Tag.GetByUserEmail(domain.Email("test-data@example.com"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fetched[0].ID, domain.ID("0"))
	assert.Equal(t, fetched[0].Name, "hoge")
	assert.Equal(t, fetched[0].Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(fetched[0].Member), 0)
	assert.Equal(t, fetched[0].Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, fetched[0].Color, "#fff")
}

func TestGetModeratingTags(t *testing.T) {
	fetched, err := mock.Tag.GetModeratingTags(domain.Email("test-data@example.com"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fetched[0].ID, domain.ID("0"))
	assert.Equal(t, fetched[0].Name, "hoge")
	assert.Equal(t, fetched[0].Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(fetched[0].Member), 0)
	assert.Equal(t, fetched[0].Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, fetched[0].Color, "#fff")
}

func TestDeleteTag(t *testing.T) {
	id, err := domain.GenerateID()
	if err != nil {
		t.Fatal(err)
	}

	email, err := domain.ConstructEmail("test-data@example.com")
	if err != nil {
		t.Fatal(err)
	}

	user, err := mock.User.GetByEmail(email)
	if err != nil {
		t.Fatal(err)
	}

	tag := domain.ConstructTag(id, "hoge", "#fff", []domain.User{user}, []domain.User{}, domain.None)
	_, err = mock.Tag.CreateTag(tag)
	if err != nil {
		t.Fatal(err)
	}

	deleted, err := mock.Tag.DeleteTag(id)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, deleted.ID, id)
	assert.Equal(t, deleted.Name, "hoge")
	assert.Equal(t, deleted.Admin[0].Email, domain.Email("test-data@example.com"))
	assert.Equal(t, len(deleted.Member), 0)
	assert.Equal(t, deleted.Type.Parse(), domain.TagType(domain.None).Parse())
	assert.Equal(t, deleted.Color, "#fff")
}
