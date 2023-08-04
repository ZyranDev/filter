package tests

import (
	"encoding/json"
	"go.zyran.dev/filter"
	"go.zyran.dev/filter/entry"
	"testing"
)

func TestFilter(t *testing.T) {
	userFilter := filter.NewFilter(
		[]entry.Entry{
			{
				Name:   "id",
				Scopes: []string{"user.id", "user.details", "user.name"},
			},
			{
				Name:   "name",
				Scopes: []string{"user.name"},
			},
		},
		func(user any, entries []entry.Entry) error {
			usr, ok := user.(*User)
			if !ok {
				return filter.ErrInvalidType
			}
			for _, e := range entries {
				switch e.Name {
				case "id":
					usr.ID = 0
					continue
				case "name":
					usr.Name = ""
					continue
				}
			}
			return nil
		},
	)
	user := &User{ID: 134, Name: "John Doe"}

	err := userFilter.Filter(user, []string{"user.id"})
	if err != nil {
		t.Fatal(err)
		return
	}
	bytes, _ := json.Marshal(user)
	t.Logf("User: %s", string(bytes))
}
