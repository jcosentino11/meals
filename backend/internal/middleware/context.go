package middleware

import (
	"github.com/jcosentino11/meals/internal/auth"
	"github.com/jcosentino11/meals/internal/mongo"
	"github.com/jcosentino11/meals/internal/context"
	"github.com/labstack/echo/v4"
)

// WrapContext takes echo.Context and converts it to a MealsContext
func WrapContext(db *mongo.BasicClient, auth *auth.FirebaseAuth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.MealsContext{
				Db:      db,
				Auth:    auth,
				Context: c,
			}
			return next(cc)
		}
	}
}
