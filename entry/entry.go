package entry

import "regexp"

type Entry struct {
	Name   string
	Scopes []string
}

func (e *Entry) Valid(scopes []string) bool {
	if e.Scopes == nil || len(e.Scopes) <= 0 {
		return false
	}
	for _, scope := range scopes {
		for _, s := range e.Scopes {
			if s == scope {
				return true
			}
			pattern := "^" + scope + "$"
			if ok, _ := regexp.MatchString(pattern, s); ok {
				return true
			}
		}
	}
	return false
}
