package error

import "fmt"

type InvalidArgumentError struct {
	Msg   string
	Field string
}

func (e InvalidArgumentError) Code() Code {
	return CodeBadRequest
}

func (e InvalidArgumentError) Error() string {
	return fmt.Sprintf("Field: %s, \tMsg: %s", e.Field, e.Msg)
}
