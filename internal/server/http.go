package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) runHttpServer() {

	s.echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})
}
