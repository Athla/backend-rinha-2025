-- name: GetPaymentByCorrelationId :one
SELECT CORRELATION_ID, AMOUNT, TIMESTAMP FROM PAYMENTS
WHERE CORRELATION_ID = ? LIMIT 1;

-- name: GetAllPayments :many
SELECT CORRELATION_ID, AMOUNT, TIMESTAMP FROM PAYMENTS
ORDER BY TIMESTAMP DESC;

-- name: GetPaymentsByInterval :many
SELECT CORRELATION_ID, AMOUNT, TIMESTAMP FROM PAYMENTS
WHERE TIMESTAMP BETWEEN ? AND ?
ORDER BY TIMESTAMP DESC;

-- name: SummarizePaymentsByInterval :many
SELECT SUM(AMOUNT), PROCESSOR, COUNT(*) FROM PAYMENTS
WHERE TIMESTAMP BETWEEN ? AND ?
GROUP BY PROCESSOR;

-- name: CreatePaymentRecord :one
INSERT INTO PAYMENTS (
  CORRELATION_ID, AMOUNT, TIMESTAMP, PROCESSOR
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- {
--     "default" : {
--         "totalRequests": 43236,
--         "totalAmount": 415542345.98
--     },
--     "fallback" : {
--         "totalRequests": 423545,
--         "totalAmount": 329347.34
--     }
-- }
