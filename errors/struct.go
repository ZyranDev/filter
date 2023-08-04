package errors

import "errors"

var (
	IsNotStruct        = errors.New("value is not a struct")
	IsNotStructPointer = errors.New("value is not a struct pointer")
)
