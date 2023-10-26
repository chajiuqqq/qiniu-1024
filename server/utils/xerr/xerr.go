package xerr

import "fmt"

// ServerError always the same
var ServerError = New(500, "ServerError",
	"server error, please report to us or try again later.")

// XError
type XError struct {
	// ErrorKey is PascalCase
	ErrorKey   string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

// New create an XError
func New(code int, key string, msg string) *XError {
	return &XError{
		StatusCode: code,
		ErrorKey:   key,
		Message:    msg,
	}
}

// Newf create an XError use format
func Newf(code int, key string, format string, a ...interface{}) *XError {
	return New(code, key, fmt.Sprintf(format, a...))
}

// Error makes it compatible with `error` interface.
func (e *XError) Error() string {
	return e.Message
}

// IsXError "err" the instance of Error,and has <key>?
func IsXError(err error, key string) bool {
	src, ok := err.(*XError)
	if !ok {
		return false
	}
	if src.ErrorKey == key {
		return true
	}
	return false
}
