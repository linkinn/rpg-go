package helpers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type Error struct {
	err     error
	message string
	code    ErrorCode
}

type ErrorCode uint

func WrapErrorf(err error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		code:    code,
		err:     err,
		message: fmt.Sprintf(format, a...),
	}
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%s: %v", e.message, e.err)
	}

	return e.message
}

// ServerError will display error page for internal server error
func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = log.Output(2, trace)

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	http.ServeFile(w, r, "./ui/static/500.html")
}
