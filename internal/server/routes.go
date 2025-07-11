package server

import (
	"backend/internal/models"
	"github.com/charmbracelet/log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	// e.Use(middleware.Logger())
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
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (s *Server) PaymentHandler(c echo.Context) error {
	var paymentReq models.PaymentRequest
	if err := c.Bind(&paymentReq); err != nil {
		log.Errorf("Unable to bind req due: %v", err)
		return c.JSON(http.StatusBadRequest, models.FailureResponse{
			Message:   "Invalid request body. Unable to bind.",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}

	// Do something later

	go s.paymentService.ProcessPayment(c.Request().Context(), &paymentReq)

	return c.JSON(http.StatusOK, models.SuccessResponse{
		Message:   "Payment received!",
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func (s *Server) PaymentsSummaryHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, models.SummaryResponse{
		Default: models.ReqSummary{
			TotalRequest: 10000,
			TotalAmount:  102391231.0,
		},
		Fallback: models.ReqSummary{
			TotalRequest: 10000,
			TotalAmount:  102391231.0,
		},
	})

}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
