package response

import (
	"nitic-pbl-2022-01/pbl-back/src/domain"
	"strconv"
)

type Repeat struct {
  Unit string `json:"unit"`
  Value int `json:"value"`
  Since string `json:"since"`
  Until string `json:"until"`
}

func ConvertRepeat(repeat domain.Repeat) Repeat {
  return Repeat {
    Unit: repeat.Unit.Parse(),
    Value: repeat.Value,
    Since: strconv.Itoa(int(repeat.Since.Unix())),
    Until: strconv.Itoa((repeat.Until.Unix())),
  }
}
