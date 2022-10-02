package utils

import (
	"Improve/src/errors"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(ctx *gin.Context, data interface{}) {
	if data != nil {
		ctx.JSON(http.StatusOK, data)
	} else {
		ctx.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}

// HandleError handles error of HTTP request.
func HandleError(c *gin.Context, err error) {
	if err != nil {
		appErr, ok := err.(errors.AppError)
		if ok {
			newContext := SetErrorCode(c.Request.Context(), appErr.Meta.Code)
			c.Request = c.Request.WithContext(newContext)
			c.JSON(appErr.ErrorCode.Status(), appErr)
		} else {
			c.JSON(errors.InternalServerError.Status(), errors.New(errors.InternalServerError))
		}
	} else {
		c.JSON(errors.ErrNoResponse.Status(), errors.New(errors.ErrNoResponse))
	}
}

type ctxKey string

var userIDCtxKey ctxKey = "user_id"
var errorCodeCtxKey ctxKey = "error_code"

func SetErrorCode(parent context.Context, errorCode int) context.Context {
	return setCtxValue(parent, errorCodeCtxKey, errorCode)
}
func setCtxValue(parent context.Context, key ctxKey, val interface{}) context.Context {
	return context.WithValue(parent, key, val)
}
