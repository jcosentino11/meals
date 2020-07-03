package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Context is simple wrapper around echo.Context
// providing access to data layer, auth layer.
type Context struct {
	// database client
	Db MongoClient
	// optional token provided with the request
	Token *jwt.Token
	echo.Context
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := NewConfigFromEnv()

	// establish db connection
	db := NewMongoClient(MongoConfig{
		MockDb:   conf.MockDb,
		Hosts:    conf.MongoHosts,
		Port:     conf.MongoPort,
		Username: conf.MongoUsername,
		Password: conf.MongoPassword,
	})
	if db.Initialize() != nil {
		log.Fatalf("Unable to establish db connection: %s", err)
	}

	// setup web server
	e := echo.New()

	// TODO restrict to personal domain
	e.Use(middleware.CORS())

	e.Use(NewWrapContextMiddleware(func(ctx echo.Context) echo.Context {
		return &Context{
			Db:      db,
			Context: ctx,
		}
	}),
	)

	e.Use(NewJwtMiddleware(
		JwtMiddlewareOptions{
			Enabled:       true,
			SigningMethod: jwt.SigningMethodHS256,
			KeyFunc: func(token *jwt.Token) (interface{}, error) {
				// TODO https://auth0.com/docs/quickstart/backend/golang
				return []byte(conf.AuthSecret), nil
			}}),
	)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api")
	g.GET("/", RouteHelloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}
