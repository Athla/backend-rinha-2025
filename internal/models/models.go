package models

type Money struct {
	Value          int
	ConversionRate int
}

// 1000 -> 1000/100 -> 10

// Check if conversionn rate is working properly in this context
func NewMoneyEntry(amount int, conversionRate int) *Money {
	return &Money{
		Value:          amount,
		ConversionRate: conversionRate,
	}
}

func (m *Money) Parse() string {
	return ""
}

type PaymentRequest struct {
	CorrelationId string `json:"correlationId"`
	Amount        int16  `json:"amount"`
	RequestedAt   string `json:"requestedAt"`
}
