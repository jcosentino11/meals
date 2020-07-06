package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloWorldResponse is a test. Kindly move along...
type HelloWorldResponse struct {
	Message string `json:"message"`
}

// RouteHelloWorld is a test. Kindly move along...
func RouteHelloWorld(c echo.Context) error {
	cc := c.(*Context)
	numDatabases, err := cc.Db.GetNumDatabases()
	if err != nil {
		return err
	}

	resp := &HelloWorldResponse{
		Message: fmt.Sprintf("Hello, World! Found %d databases.", numDatabases),
	}

	return cc.JSON(http.StatusOK, resp)
}
