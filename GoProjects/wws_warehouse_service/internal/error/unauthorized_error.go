package error

import "fmt"

type UnauthorizedError struct {
	Msg    string
	UserID interface{}
}

func (e UnauthorizedError) Code() Code {
	return CodeUnauthorized
}

func (e UnauthorizedError) Error() string {
	return fmt.Sprintf("Msg: %s \t User: %v unauthorized", e.Msg, e.UserID)
}
