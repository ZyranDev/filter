package filter

import (
	"errors"
	"filter/entry"
)

var (
	ErrInvalidType = errors.New("invalid type")
)

type Filter struct {
	entries []entry.Entry
	filter  func(any, []entry.Entry) error
}

func NewFilter(entries []entry.Entry, filter func(any, []entry.Entry) error) *Filter {
	return &Filter{
		entries,
		filter,
	}
}

func (f *Filter) Filter(data any, scopes []string) error {
	var decline []entry.Entry
	for _, e := range f.entries {
		if valid := e.Valid(scopes); !valid {
			decline = append(decline, e)
		}
	}
	return f.filter(data, decline)
}
