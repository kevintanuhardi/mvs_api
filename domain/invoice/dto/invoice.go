package dto

type GetInvoiceRequest struct {
	InvoiceDate string `json:"invoice_date"`
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
}