package api

type NotFoundError struct {
	Message string
}

func (n NotFoundError) Error() string {
	return n.Message
}
