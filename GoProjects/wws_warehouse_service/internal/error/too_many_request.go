package error

type TooManyRequest struct {
	Msg string
}

func (e TooManyRequest) Code() Code {
	return CodeTooManyRequest
}

func (e TooManyRequest) Error() string {
	return e.Msg
}
