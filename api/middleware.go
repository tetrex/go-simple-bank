package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tetrex/backend-masterclass-go/token"
	"github.com/tetrex/backend-masterclass-go/util"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authorizationHeader := ctx.Request().Header.Get(authorizationHeaderKey)

			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				return ctx.JSON(http.StatusUnauthorized, util.ErrorResponse{Error: err.Error()})
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				err := errors.New("invalid authorization header format")
				return ctx.JSON(http.StatusUnauthorized, util.ErrorResponse{Error: err.Error()})
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				return ctx.JSON(http.StatusUnauthorized, util.ErrorResponse{Error: err.Error()})
			}

			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
			}

			ctx.Set(authorizationPayloadKey, payload)
			return next(ctx)
		}
	}
}
