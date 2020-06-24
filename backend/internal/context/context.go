package context

import (
	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/jcosentino11/meals/internal/auth"
	"github.com/jcosentino11/meals/internal/mongo"
	"github.com/labstack/echo/v4"
)

// MealsContext is simple wrapper around echo.Context
// providing access to data layer, auth layer.
type MealsContext struct {
	Db    *mongo.BasicClient
	Auth  *auth.FirebaseAuth
	Token *firebaseAuth.Token
	echo.Context
}
