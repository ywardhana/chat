package http

type createChatRequest struct {
	Message       string `json:"message"`
	InvoiceNumber string `json:"invoice_number"`
}
