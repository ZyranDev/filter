package _struct

import (
	"filter/entry"
	"filter/errors"
	"reflect"
)

func StructFilter() func(any, []entry.Entry) error {
	return func(value any, entries []entry.Entry) error {
		v := reflect.ValueOf(value)
		if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
			return errors.IsNotStructPointer
		}
		for _, e := range entries {
			field := v.Elem().FieldByName(e.Name)
			if !field.IsValid() {
				continue
			}
			field.Set(reflect.Zero(field.Type()))
		}
		return nil
	}
}
