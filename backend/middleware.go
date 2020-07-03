package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var (
	// ErrUnauthorized is returned when user is not authorized
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
)

// NewWrapContextMiddleware wraps the default echo.Context with a custom one
func NewWrapContextMiddleware(ctx func(echo.Context) echo.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(ctx(c))
		}
	}
}

// JwtMiddlewareOptions define options for jwt middleware
type JwtMiddlewareOptions struct {
	Enabled       bool
	KeyFunc       jwt.Keyfunc
	SigningMethod jwt.SigningMethod
}

// NewJwtMiddleware create new jwt token auth middleware
func NewJwtMiddleware(options JwtMiddlewareOptions) echo.MiddlewareFunc {
	jwt := &Jwt{
		KeyFunc:       options.KeyFunc,
		SigningMethod: options.SigningMethod,
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if options.Enabled {
				parsedToken, err := jwt.ParseTokenFromRequest(c.Request())

				if err != nil {
					return ErrUnauthorized
				}

				if jwt.VerifySigningMethod(parsedToken) != nil {
					return ErrUnauthorized
				}

				if !parsedToken.Valid {
					return ErrUnauthorized
				}

				// add jwt token to the request context
				cc := c.(*Context)
				cc.Token = parsedToken
			}

			return next(c)
		}
	}
}
