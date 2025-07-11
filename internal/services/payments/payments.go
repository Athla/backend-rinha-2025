package payments

import (
	"backend/config"
	"backend/internal/models"
	"backend/internal/repository"
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type PaymentProcessor interface {
	Process(*models.PaymentRequest) error
}

type Service struct {
	mainService     PaymentProcessor
	fallbackService PaymentProcessor
}

func NewPaymentService(cfg *config.Config) *Service {
	// Initalize db connection
	dbconn, err := sqlx.Open("sqlite3", cfg.SqlConnString)
	dbtx := repository.New(dbconn)

	mainService, err := newDefaultPaymentProcessor(cfg.DefaultServiceAddr, dbtx)
	if err != nil {
		panic(err)
	}

	fallbackService, err := newFallbackPaymentProcessor(cfg.FallbackServiceAddr, dbtx)
	if err != nil {
		panic(err)
	}

	return &Service{
		mainService:     mainService,
		fallbackService: fallbackService,
	}
}

// TODO
// func fallbackContextCancelFunc() {}

func (s *Service) ProcessPayment(ctx context.Context, paymentRequest *models.PaymentRequest) error {
	err := s.mainService.Process(paymentRequest)
	if err != nil {
		err := s.fallbackService.Process(paymentRequest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) RetrievePayments(ctx context.Context, from, to string) {

}
