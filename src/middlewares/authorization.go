package middlewares

import (
	"Improve/src/errors"
	"Improve/src/logger"
	"Improve/src/token"
	"Improve/src/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeaderKey    = "authorization"
	authorizationHeaderBearer = "bearer"
	AuthorizationPayloadKey   = "authorization_payload"
)

func AuthMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			logger.Context(ctx).Errorf("[authMiddleware] authorization header is not provide")
			utils.HandleError(ctx, errors.New(errors.HeaderNotProvideError))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			logger.Context(ctx).Errorf("[authMiddleware] invalid authorization format")
			utils.HandleError(ctx, errors.New(errors.HeaderInvalidError))
			return
		}
		authorizationHeaderType := strings.ToLower(fields[0])
		if authorizationHeaderType != authorizationHeaderBearer {
			logger.Context(ctx).Errorf("[authMiddleware] unsupported authorization")
			utils.HandleError(ctx, errors.New(errors.UnauthorizedCodeError))
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			logger.Context(ctx).Errorf("[authMiddleware] verify token err %v: ", err)
			utils.HandleError(ctx, errors.New(errors.UnauthorizedCodeError))
			return
		}
		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
