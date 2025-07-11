package repository

import (
	"context"
)

type SQL interface {
	CreatePaymentRecord(ctx context.Context, createParams CreatePaymentRecordParams) (Payment, error)
	GetPaymentByCorrelationId(ctx context.Context, correlationId any) (GetPaymentByCorrelationIdRow, error)
	GetPaymentsByInterval(ctx context.Context, intervalParams GetPaymentsByIntervalParams) ([]GetPaymentsByIntervalRow, error)
}

type sqlRepository struct {
	*Queries
}

func NewSQLRepository(db DBTX) SQL {
	return &sqlRepository{
		Queries: New(db),
	}
}

func (r *sqlRepository) CreatePaymentRecord(ctx context.Context, createParams CreatePaymentRecordParams) (Payment, error) {
	return r.Queries.CreatePaymentRecord(ctx, createParams)
}

func (r *sqlRepository) GetPaymentByCorrelationId(ctx context.Context, correlationId any) (GetPaymentByCorrelationIdRow, error) {
	return r.Queries.GetPaymentByCorrelationId(ctx, correlationId)
}

// TODO - Proper time conversion and validation
func (r *sqlRepository) GetPaymentsByInterval(ctx context.Context, params GetPaymentsByIntervalParams) ([]GetPaymentsByIntervalRow, error) {
	return r.Queries.GetPaymentsByInterval(ctx, params)
}
