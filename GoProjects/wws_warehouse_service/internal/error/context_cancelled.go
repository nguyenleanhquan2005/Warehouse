package error

type ContextCancelled struct {
	Msg string
}

func (e *ContextCancelled) Code() Code {
	return CodeContextCancelled
}

func (e *ContextCancelled) Error() string {
	return e.Msg
}
