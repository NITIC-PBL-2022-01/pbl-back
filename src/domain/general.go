package domain

import "github.com/google/uuid"

type ID string

func GenerateID() (ID, error) {
  id, err := uuid.NewRandom()
  if err != nil {
    return ID(""), err
  }

  return ID(id.String()), nil
}
