package tests

type User struct {
	ID   int    `json:"id,omitempty" scopes:"user.id"`
	Name string `json:"name,omitempty" scopes:"user.name"`
}
