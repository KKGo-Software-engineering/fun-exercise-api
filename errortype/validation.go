package errortype

import "net/http"

type ValidationError struct {
	Message string
}

func (e ValidationError) Type() string {
	return "ValidationError"
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) Status() int {
	return http.StatusBadRequest
}
