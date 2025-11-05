package presenter

import "net/http"

type PersistentCode string

const (
	CodeConflict PersistentCode = "CONFLICT"
)

func (pc PersistentCode) String() string {
	return string(pc)
}

func (pc PersistentCode) Status() int {
	switch pc {
	case CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
