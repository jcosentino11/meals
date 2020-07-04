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

	e := echo.New()

	// TODO restrict to personal domain
	e.Use(middleware.CORS())

	e.Use(NewWrapContextMiddleware(func(ctx echo.Context) echo.Context {
		return &Context{
			Db:      db,
			Context: ctx,
		}
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtMiddleware := NewJwtMiddleware(
		JwtMiddlewareOptions{
			Enabled:       true,
			SigningMethod: jwt.SigningMethodRS256,
			KeyFunc: NewAuth0KeyGetter(Auth0KeyGetterOptions{
				ExpectedAudience: "http://localhost:8080",                                    // TODO make configurable
				ExpectedIssuer:   "https://meals-staging.us.auth0.com/",                      // TODO make configurable
				JwksEndpoint:     "https://meals-staging.us.auth0.com/.well-known/jwks.json", // TODO make configurable
			}).GetValidationKey,
		})

	g := e.Group("/api")
	g.Use(jwtMiddleware)
	g.GET("/", RouteHelloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}
