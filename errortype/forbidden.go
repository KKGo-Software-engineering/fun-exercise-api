package errortype

import "net/http"

type ForbiddenError struct {
	Message string
}

func (e ForbiddenError) Type() string {
	return "ForbiddenError"
}

func (e ForbiddenError) Error() string {
	return e.Message
}

func (e ForbiddenError) Status() int {
	return http.StatusForbidden
}
