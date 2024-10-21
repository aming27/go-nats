package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *server) runHttpServer() {
	s.echo.GET("/health", func(c echo.Context) error {
		log.Printf("Entron aqui")
		return c.String(http.StatusOK, "Ok")
	})
}
