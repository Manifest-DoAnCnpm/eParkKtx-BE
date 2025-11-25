package response

// PaymentLinkResponse represents the response for a payment link creation
type PaymentLinkResponse struct {
	Success     bool   `json:"success"`
	PayUrl      string `json:"pay_url"`
	OrderCode   int64  `json:"order_code"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Message     string `json:"message,omitempty"`
}
