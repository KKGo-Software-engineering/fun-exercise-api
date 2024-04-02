package errortype

import "net/http"

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Type() string {
	return "NotFoundError"
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e NotFoundError) Status() int {
	return http.StatusNotFound
}
