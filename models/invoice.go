package models

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

var invoices = []Invoice{}

func GetInvoices() []Invoice {
	return invoices
}

func (i Invoice) AddInvoice(invoice Invoice) {
	invoices = append(invoices, invoice)
}
