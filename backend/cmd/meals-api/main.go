package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	m "github.com/jcosentino11/meals/internal/mongo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type routeContext struct {
	db *m.BasicClient
	echo.Context
}

func main() {
	// establish db connection
	conf := m.NewConfigFromEnv()
	client, err := m.NewBasicClient(conf)
	if err != nil {
		log.Fatalf("Unable to establish db connection: %s", err)
		os.Exit(1)
	}

	// setup web server
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &routeContext{
				db:      client,
				Context: c,
			}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	cc := c.(*routeContext)
	numDatabases, err := cc.db.GetNumDatabases()
	if err != nil {
		return err
	}
	return cc.String(http.StatusOK, fmt.Sprintf("Hello, World! Found %d databases.", numDatabases))
}
