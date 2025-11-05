package error

import "fmt"

type NotFoundError struct {
	Msg    string
	Entity string
	ID     interface{}
}

func (e NotFoundError) Code() Code {
	return CodeNotFound
}

func (e NotFoundError) Error() string {
	if e.ID != nil {
		return fmt.Sprintf("Msg: %s, Entity: %s, ID: %v not found", e.Msg, e.Entity, e.ID)
	}
	return fmt.Sprintf("Msg: %s, Entity: %s, not found", e.Msg, e.Entity)
}
