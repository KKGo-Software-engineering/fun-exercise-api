package errortype

import "net/http"

type AuthenticationError struct {
	Message string
}

func (e AuthenticationError) Type() string {
	return "AuthenticationError"
}

func (e AuthenticationError) Error() string {
	return e.Message
}

func (e AuthenticationError) Status() int {
	return http.StatusUnauthorized
}
