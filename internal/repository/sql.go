package repository

import (
	"context"
	"database/sql"
	"time"
)

type SQLRepository interface {
	CreatePaymentRecord(ctx context.Context, createParams CreatePaymentRecordParams) (Payment, error)
	GetPaymentByCorrelationId(ctx context.Context, correlationId string) (Payment, error)
	GetPaymentsByInterval(ctx context.Context, fromTimestamp string, toTimestamp string) ([]Payment, error)
}

type sqlRepository struct {
	*Queries
}

func NewSQLRepository(db DBTX) SQLRepository {
	return &sqlRepository{
		Queries: New(db),
	}
}

func (r *sqlRepository) CreatePaymentRecord(ctx context.Context, createParams CreatePaymentRecordParams) (Payment, error) {
	return r.Queries.CreatePaymentRecord(ctx, createParams)
}

func (r *sqlRepository) GetPaymentByCorrelationId(ctx context.Context, email string) (Payment, error) {
	return r.Queries.GetPaymentByCorrelationId(ctx, email)
}

// TODO - Proper time conversion and validation
func (r *sqlRepository) GetPaymentsByInterval(ctx context.Context, fromTimestamp string, toTimestamp string) ([]Payment, error) {
	return r.Queries.GetPaymentsByInterval(ctx, GetPaymentsByIntervalParams{
		FromTimestamp: sql.NullTime{Time: time.Now(), Valid: true},
		ToTimestamp:   sql.NullTime{Time: time.Now(), Valid: true},
	})
}
