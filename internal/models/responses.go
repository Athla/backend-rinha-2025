package models

type SuccessResponse struct {
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Timestamp string `json:"timestamp"`
}

type FailureResponse struct {
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Timestamp string `json:"timestamp"`
}

type ReqSummary struct {
	TotalRequest int     `json:"totalRequest"`
	TotalAmount  float32 `json:"totalAmount"`
}

type SummaryResponse struct {
	Default  ReqSummary `json:"default"`
	Fallback ReqSummary `json:"fallback"`
}
