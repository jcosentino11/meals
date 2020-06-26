package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/jcosentino11/meals/internal/auth"
	"github.com/jcosentino11/meals/internal/context"
	"github.com/labstack/echo/v4"
)

// FirebaseVerifyToken middleware validates JWT token provided
// in the authorization header and adds it to the context.
func FirebaseVerifyToken(f *auth.FirebaseAuth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*context.MealsContext)

			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
			}

			// trim Bearer value
			authToken := strings.Replace(authHeader, "Bearer ", "", 1)

			token, err := f.VerifyAuthToken(c.Request().Context(), authToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token verification failed")
			}

			cc.Token = token
			return next(c)
		}
	}
}
