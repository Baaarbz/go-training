package valueobject

import (
	"errors"
	"time"
)

var ErrDateBadFormat = errors.New("date: bad format")

type Date string

func NewDate(value time.Time) Date {
	return Date(value.Format(time.RFC3339))
}
func (date Date) Time() (time.Time, error) {
	return time.Parse(time.RFC3339, string(date))
}
