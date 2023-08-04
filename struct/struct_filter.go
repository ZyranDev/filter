package _struct

import (
	"go.zyran.dev/filter/entry"
	"go.zyran.dev/filter/errors"
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
