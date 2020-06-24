package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jcosentino11/meals/internal/app"
	"github.com/jcosentino11/meals/internal/auth"
	"github.com/jcosentino11/meals/internal/context"
	mealsMiddleware "github.com/jcosentino11/meals/internal/middleware"
	"github.com/jcosentino11/meals/internal/mongo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	conf := app.NewConfigFromEnv()

	// establish db connection
	db, err := mongo.NewBasicClient(conf.MongoConfig)
	if err != nil {
		log.Fatalf("Unable to establish db connection: %s", err)
		os.Exit(1)
	}

	// create auth client
	var authClient *auth.FirebaseAuth
	if conf.AuthEnabled {
		authClient, err = auth.Firebase(nil)
		if err != nil {
			log.Fatalf("Unable to create auth client: %s", err)
			os.Exit(1)
		}
	}

	// setup web server
	e := echo.New()
	e.Use(mealsMiddleware.WrapContext(db, authClient))

	if conf.AuthEnabled {
		e.Use(mealsMiddleware.FirebaseVerifyToken(authClient))
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api")
	g.GET("/", hello)
	
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	cc := c.(*context.MealsContext)
	numDatabases, err := cc.Db.GetNumDatabases()
	if err != nil {
		return err
	}
	return cc.String(http.StatusOK, fmt.Sprintf("Hello, World! Found %d databases.", numDatabases))
}
