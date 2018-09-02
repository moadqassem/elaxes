package api

import (
	"errors"
	"github.com/labstack/echo"
)

type server struct {
	host   string
	port   int
	router *echo.Echo
}

// NewServer returns a new echo router with the sent host and port.
func NewServer(host string, port int) (*server, error) {
	if port == 0 {
		return nil, errors.New("port cannot be empty")
	}

	return &server{
		host:   host,
		port:   port,
		router: echo.New(),
	}, nil
}