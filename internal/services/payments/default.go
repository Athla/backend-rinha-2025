package payments

import (
	"backend/internal/models"
	"backend/internal/repository"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type defaultPaymentProcessor struct {
	processorUrl string
	db           repository.SQLRepository
}

func (p *defaultPaymentProcessor) Process(req *models.PaymentRequest) error {
	if err := uuid.Validate(req.CorrelationId); err != nil {
		return err
	}
	// Validar Value
	// TODO

	t, err := time.Parse(time.RFC3339, req.RequestedAt)
	if err != nil {
		return err
	}
	// send to the payment processor properly
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/payments", p.processorUrl), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	cli := &http.Client{}
	resp, err := cli.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Check if operation was correct
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Let's hop into the fallback")
	}

	tNull := sql.NullTime{Time: t, Valid: true}

	_, err = p.db.CreatePaymentRecord(context.TODO(), repository.CreatePaymentRecordParams{
		CorrelationID: req.CorrelationId,
		Amount:        req.Amount,
		Timestamp:     tNull,
	})

	if err != nil {
		return err
	}

	return nil
}

func newDefaultPaymentProcessor(serviceAddr string, db repository.SQLRepository) (*defaultPaymentProcessor, error) {
	return nil, nil
}
