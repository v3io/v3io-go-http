package v3io

import (
	"fmt"

	"github.com/nuclio/nuclio/pkg/errors"
)

type ErrorWithStatusCode struct {
	error
	statusCode int
	message    string
}

func NewErrorWithStatusCode(statusCode int, format string, args ...interface{}) ErrorWithStatusCode {
	return ErrorWithStatusCode{
		error:      errors.New(fmt.Sprintf(format, args...)),
		statusCode: statusCode,
	}
}

func (e *ErrorWithStatusCode) StatusCode() int {
	return e.statusCode
}
