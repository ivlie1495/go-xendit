package models

import "time"

type InvoiceBody struct {
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}

type Invoice struct {
	ID         string  `json:"id"`
	ExternalID string  `json:"external_id"`
	Status     string  `json:"status"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	ExpiredAt  string  `json:"expired_at"`
}

type InvoiceCreated struct {
	ID                     string    `json:"id"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 int       `json:"amount"`
	PaidAmount             int       `json:"paid_amount"`
	BankCode               string    `json:"bank_code"`
	PaidAt                 time.Time `json:"paid_at"`
	PayerEmail             string    `json:"payer_email"`
	Description            string    `json:"description"`
	AdjustedReceivedAmount int       `json:"adjusted_received_amount"`
	FeesPaidAmount         int       `json:"fees_paid_amount"`
	Updated                time.Time `json:"updated"`
	Created                time.Time `json:"created"`
	Currency               string    `json:"currency"`
	PaymentChannel         string    `json:"payment_channel"`
	PaymentDestination     string    `json:"payment_destination"`
}

var invoices = []Invoice{}

func GetInvoices() []Invoice {
	return invoices
}

func (i Invoice) AddInvoice(invoice Invoice) {
	invoices = append(invoices, invoice)
}
