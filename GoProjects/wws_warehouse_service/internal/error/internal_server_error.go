package error

type InternalServerError struct {
	Msg string
}

func (e InternalServerError) Code() Code {
	return CodeInternalServer
}

func (e InternalServerError) Error() string {
	return e.Msg
}
