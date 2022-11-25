package domain

import "time"

type RepeatUnit int

const (
  Day = iota
  Week
  Month
  Year
)

func (unit RepeatUnit) Parse() string {
  switch (unit) {
    case Day:
      return "DAY"
    case Week:
      return "WEEK"
    case Month:
      return "MONTH"
    case Year:
      return "YEAR"
    default:
      return "UNREACHABLE"
  }
}

type Repeat struct {
  Unit RepeatUnit
  Value int
  Since time.Time
  Until time.Time
}

func ConstructRepeat(unit RepeatUnit, value int, since time.Time, until time.Time) Repeat {
  return Repeat {
    Unit: unit,
    Value: value,
    Since: since,
    Until: until,
  }
}
