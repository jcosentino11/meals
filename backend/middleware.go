package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

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
	Enabled          bool
	SigningMethod    jwt.SigningMethod
	ExpectedAudience string
	ExpectedIssuer   string
	JwksEndpoint     string
}

// NewJwtMiddleware create new jwt token auth middleware
func NewJwtMiddleware(options JwtMiddlewareOptions) echo.MiddlewareFunc {
	jwt := &Jwt{
		SigningMethod:    options.SigningMethod,
		ExpectedAudience: options.ExpectedAudience,
		ExpectedIssuer:   options.ExpectedIssuer,
		JwksEndpoint:     options.JwksEndpoint,
	}
	jwt.Initialize()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if options.Enabled {
				parsedToken, err := jwt.ParseTokenFromRequest(c.Request())
				if err != nil {
					log.Printf("auth error: %s", err)
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
