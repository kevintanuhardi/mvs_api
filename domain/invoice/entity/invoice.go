package entity

import "time"

type Invoice struct {
	ID *int `json:"id,omitempty"`
	InvoiceNo int`json:"invoice_no"`
	InvoiceDate time.Time `json:"invoice_date"`
	PaymentPeriodInDays int `json:"payment_period_in_days"`
	ShippingContactName string`json:"shipping_contact_name"`
	ShippingContactPhone string`json:"shipping_contact_phone"`
	ShippingAddress string`json:"shipping_address"`
	ShippingCity string`json:"shipping_city"`
	ShippingPostalCode string`json:"shipping_postal_code"`
	BillingContactName string`json:"billing_contact_name"`
	BillingContactPhone string`json:"billing_contact_phone"`
	BillingAddress string`json:"billing_address"`
	BillingCity string`json:"billing_city"`
	BillingPostalCode string`json:"billing_postal_code"`
	Status string`json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Invoice) TableName() string {
	return "invoice"
}