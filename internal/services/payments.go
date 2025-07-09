package services

import (
	"backend/config"
	"backend/internal/models"
	"database/sql"
)

type PaymentProcessor interface {
	Process(models.PaymentRequest) error
}

type defaultPaymentProcessor struct {
	db *sql.DB
}

func (p *defaultPaymentProcessor) Process(models.PaymentRequest) error {

	return nil
}

func newDefaultPaymentProcessor() (*defaultPaymentProcessor, error) {

	return nil, nil
}

type fallbackPaymentProcessor struct {
	db *sql.DB
}

func newFallbackPaymentProcessor() (*fallbackPaymentProcessor, error) {

	return nil, nil
}

func (p *fallbackPaymentProcessor) Process(models.PaymentRequest) error {

	return nil
}

type PaymentService struct {
	mainService     PaymentProcessor
	fallbackService PaymentProcessor
}

func NewPaymentService(cfg *config.Config) *PaymentService {
	mainService, err := newDefaultPaymentProcessor()
	if err != nil {
		panic(err)
	}
	fallbackService, err := newFallbackPaymentProcessor()
	if err != nil {
		panic(err)
	}

	return &PaymentService{
		mainService:     mainService,
		fallbackService: fallbackService,
	}
}
