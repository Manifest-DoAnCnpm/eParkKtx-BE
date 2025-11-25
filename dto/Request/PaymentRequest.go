package request

// CreatePaymentLinkRequest represents the request body for creating a payment link
type CreatePaymentLinkRequest struct {
	Amount      int64  `json:"amount" binding:"required"`
	Description string `json:"description" binding:"required"`
	OrderCode   int64  `json:"order_code,omitempty"`
}
