package tests

import (
	"encoding/json"
	"filter"
	"filter/struct"
	"testing"
)

func TestStructEntries(t *testing.T) {
	var user User
	entries, err := _struct.FromStruct(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("Entries: %v", entries)
}

func TestStructFilter(t *testing.T) {
	var user User

	entries, err := _struct.FromStruct(user)
	if err != nil {
		t.Error(err)
		return
	}

	userFilter := filter.NewFilter(
		entries,
		_struct.StructFilter(),
	)

	user = User{ID: 134, Name: "John Doe"}

	if err = userFilter.Filter(&user, []string{"user.name"}); err != nil {
		t.Fatal(err)
		return
	}

	bytes, _ := json.Marshal(user)
	t.Logf("User: %s", string(bytes))
}
