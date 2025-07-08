package repository

import "backend/internal/models"

type SQL interface {
	ProcessPayment(entry models.PaymentRequest) error
	FetchPayment(correlationId string) error
}
