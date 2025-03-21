package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func echoServer() {
	_echo := echo.New()

	_echo.GET("/", helloEcho)

	_echo.Use(middleware.Logger())  // logs each incoming HTTP request with method, path, status code, and latency
	_echo.Use(middleware.Recover()) // recovers from panics within HTTP handlers and prevents the server from crashing. If a panic occurs, it logs the error and returns a 500 Internal Server Error.

	_error := _echo.Start(":8080")

	if _error != nil && !errors.Is(_error, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", _error)
	}
}

func helloEcho(c echo.Context) error {
	return c.String(http.StatusOK, "hello world, my name is echo!")
}
