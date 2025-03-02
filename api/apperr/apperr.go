package apperr

import (
	"fmt"
	"net/http"
)

type Error struct {
	err             error
	id              string
	code            int
	message         string
	messageJapanese string
}

func (e Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %s", e.message, e.err)
	}
	return e.message
}

func ErrDuplicateUserEmail(err error) Error {
	return Error{
		err:             err,
		id:              "duplicate-user-email",
		code:            http.StatusConflict,
		message:         "the mail address is already in use",
		messageJapanese: "そのメールアドレスは既に使用されています",
	}
}
