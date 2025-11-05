package error

import "fmt"

type ForbiddenError struct {
	Msg        string
	UserID     interface{}
	ResourceID interface{}
}

func (e ForbiddenError) Code() Code {
	return CodeForbidden
}

func (e ForbiddenError) Error() string {
	return fmt.Sprintf("User: %v forbidden to access resource: %v, \tMsg: %s", e.UserID, e.ResourceID, e.Msg)
}
