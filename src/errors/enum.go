package errors

import "fmt"

type ErrorType uint16

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorConfigDirectoryNotFound
)

var errorTypeNames = []string{
	"Unknown",
	"ConfigDirectoryNotFound",
}

type NextLaunchError struct {
	Code     ErrorType
	Message  string
	CodeName string
	Fatal    bool
}

func (e *NextLaunchError) Error() string {
	return fmt.Sprintf("%d - %s: %s", e.Code, e.CodeName, e.Message)
}

func (e *NextLaunchError) FatalError() {
	if e.Fatal {
		panic(e)
	}
}

func NewError(code ErrorType, error error, fatal bool) *NextLaunchError {
	return &NextLaunchError{
		Code:     code,
		Message:  error.Error(),
		CodeName: errorTypeNames[code],
		Fatal:    fatal,
	}
}

func NewErrorf(code ErrorType, fatal bool, format string, a ...interface{}) *NextLaunchError {
	return &NextLaunchError{
		Code:     code,
		Message:  fmt.Sprintf(format, a...),
		CodeName: errorTypeNames[code],
		Fatal:    fatal,
	}
}
