package errors

import (
	"fmt"
	"strconv"
)

// ErrorCode contains HTTP status, module and detail code.
type ErrorCode struct {
	status     int
	module     int
	detailCode int
}

func fmtErrorCode(status, module, code int) ErrorCode {
	return ErrorCode{
		status:     status,
		module:     module,
		detailCode: code,
	}
}

// Code returns the integer with format 4xxyyzz.
func (errCode ErrorCode) Code() int {
	errStr := fmt.Sprintf("%d%02d%02d", errCode.status, errCode.module, errCode.detailCode)
	code, _ := strconv.Atoi(errStr)
	return code
}

// Status returns HTTP status code.
func (errCode ErrorCode) Status() int {
	return errCode.status
}

// Module returns module error code.
func (errCode ErrorCode) Module() int {
	return errCode.module
}

// DetailCode returns detail error code.
func (errCode ErrorCode) DetailCode() int {
	return errCode.detailCode
}

// AppError describes application error.
type AppError struct {
	Meta          ErrorMeta `json:"meta"`
	OriginalError error     `json:"-"`
	ErrorCode     ErrorCode `json:"-"`
}

// ErrorMeta is the metadata of AppError.
type ErrorMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (appErr AppError) Error() string {
	if appErr.OriginalError != nil {
		return appErr.OriginalError.Error()
	}
	return appErr.Meta.Message
}

// New returns an AppError with args.
func New(errCode ErrorCode, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: GetErrorMessage(errCode, args...),
		},
		OriginalError: nil,
		ErrorCode:     errCode,
	}
}

// Newf returns an AppError with args and message.
func Newf(errCode ErrorCode, msg string, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: fmt.Sprintf(msg, args...),
		},
		OriginalError: nil,
		ErrorCode:     errCode,
	}
}

// Wrap returns an AppError with err, args.
func Wrap(errCode ErrorCode, err error, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: GetErrorMessage(errCode, args...),
		},
		OriginalError: err,
		ErrorCode:     errCode,
	}
}

// Wrapf returns an AppError with err, args and message.
func Wrapf(errCode ErrorCode, err error, msg string, args ...interface{}) error {
	return AppError{
		Meta: ErrorMeta{
			Code:    errCode.Code(),
			Message: fmt.Sprintf(msg, args...),
		},
		OriginalError: err,
		ErrorCode:     errCode,
	}
}
