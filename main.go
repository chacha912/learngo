package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", handleHome)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
