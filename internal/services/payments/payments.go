package payments

import (
	"backend/config"
	"backend/internal/models"
	"context"
)

type PaymentProcessor interface {
	Process(*models.PaymentRequest) error
}

type Service struct {
	mainService     PaymentProcessor
	fallbackService PaymentProcessor
}

func NewPaymentService(cfg *config.Config) *Service {
	mainService, err := newDefaultPaymentProcessor()
	if err != nil {
		panic(err)
	}

	fallbackService, err := newFallbackPaymentProcessor()
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
