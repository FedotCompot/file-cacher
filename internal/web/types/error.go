package types

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(err error) Error {
	return Error{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}
