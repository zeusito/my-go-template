package errors

import "fmt"

type ErrorCode string

const (
	ErrorCodeInternalError      ErrorCode = "internal_error"
	ErrorCodePreconditionFailed ErrorCode = "precondition_failed"
	ErrorCodeForbidden          ErrorCode = "forbidden"
	ErrorCodeNotFound           ErrorCode = "not_found"
	ErrorCodeBadRequest         ErrorCode = "bad_request"
)

type CustomError struct {
	Code    ErrorCode
	Message string
}

func New(code ErrorCode, message string) CustomError {
	return CustomError{
		Code:    code,
		Message: message,
	}
}

func Newf(code ErrorCode, format string, args ...any) CustomError {
	return New(code, fmt.Sprintf(format, args...))
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) GetCode() ErrorCode {
	return e.Code
}
