package server

import (
	"backend/internal/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.GET("/ping", s.PingHandler)
	e.GET("/", s.HelloWorldHandler)

	e.POST("/payments", s.PaymentHandler)
	e.GET("/payments-summary", s.PaymentsSummaryHandler)

	return e
}

func (s *Server) PingHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message:   "pong",
		Timestamp: time.Now(),
	})
}

func (s *Server) PaymentHandler(c echo.Context) error {
	// recebendo uma struct/payload -> models.Payment

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message:   "Everything working as intenteded in payments endpoint",
		Timestamp: time.Now(),
	})
}

func (s *Server) PaymentsSummaryHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message:   "Everything working as intenteded in the endpoint payments summary",
		Timestamp: time.Now(),
	})

}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
