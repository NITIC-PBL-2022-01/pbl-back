package domain

import (
	"net/mail"

	"github.com/google/uuid"
)

type ID string

func GenerateID() (ID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return ID(""), err
	}

	return ID(id.String()), nil
}

type Email string

func ConstructEmail(str string) (Email, error) {
	_, err := mail.ParseAddress(str)
	if err != nil {
		return Email(""), err
	}
	return Email(str), nil
}
