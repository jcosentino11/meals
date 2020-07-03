package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RouteHelloWorld is a test. Kindly move along...
func RouteHelloWorld(c echo.Context) error {
	cc := c.(*Context)
	numDatabases, err := cc.Db.GetNumDatabases()
	if err != nil {
		return err
	}
	return cc.String(http.StatusOK, fmt.Sprintf("Hello, World! Found %d databases.", numDatabases))
}
