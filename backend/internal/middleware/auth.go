package middleware

import (
	"net/http"

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
			token, err := f.VerifyAuthToken(c.Request().Context(), c.Request().Header.Get("Authorization"))
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token verification failed")
			}
			cc.Token = token
			return next(c)
		}
	}
}
