package user

type ErrNotFound struct {
	Message string
}

func (e ErrNotFound) Error() string {
	return e.Message
}
