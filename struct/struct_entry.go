package _struct

import (
	"filter/entry"
	"filter/errors"
	"reflect"
	"strings"
)

func FromStruct(value interface{}) (entries []entry.Entry, err error) {
	return FromStructWithTag(value, "scopes")
}

func FromStructWithTag(value interface{}, tag string) (entries []entry.Entry, err error) {
	t := reflect.TypeOf(value)
	if t.Kind() != reflect.Struct {
		return nil, errors.IsNotStruct
	}
	field := t.NumField()
	for i := 0; i < field; i++ {
		f := t.Field(i)
		raw := f.Tag.Get(tag)
		if raw == "" {
			continue
		}
		var scopes []string
		for _, scope := range strings.Split(raw, ",") {
			scope = strings.TrimSpace(scope)
			if scope != "" {
				scopes = append(scopes, scope)
			}
		}
		entries = append(entries, entry.Entry{
			Name:   f.Name,
			Scopes: scopes,
		})
	}
	return
}
