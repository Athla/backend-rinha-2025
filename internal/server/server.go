package server

import (
	"backend/config"
	"backend/internal/services/payments"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	cfg            *config.Config
	port           int
	paymentService *payments.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	cfg := config.New()
	svc := payments.NewPaymentService(cfg)

	NewServer := &Server{
		cfg:            cfg,
		port:           port,
		paymentService: svc,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
