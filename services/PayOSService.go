package services

import (
	"eParkKtx/dto/request"
	"eParkKtx/dto/response"
	"fmt"
	"time"

	payos "github.com/payOSHQ/payos-lib-golang"
)

type PayOSService struct {
	clientID    string
	apiKey      string
	checksumKey string
}

func NewPayOSService(clientID, apiKey, checksumKey string) *PayOSService {
	return &PayOSService{
		clientID:    clientID,
		apiKey:      apiKey,
		checksumKey: checksumKey,
	}
}

// CreatePaymentLink creates a payment link using PayOS
func (s *PayOSService) CreatePaymentLink(req request.CreatePaymentLinkRequest) (*response.PaymentLinkResponse, error) {
	// Initialize PayOS with credentials
	payos.Key(s.clientID, s.apiKey, s.checksumKey)

	// If no order code is provided, use current timestamp
	orderCode := req.OrderCode
	if orderCode == 0 {
		orderCode = time.Now().Unix()
	}

	// Prepare payment request
	paymentReq := payos.CheckoutRequestType{
		OrderCode:   orderCode,
		Amount:      int(req.Amount),
		Description: req.Description,
		CancelUrl:   "http://localhost:8080/api/payment/cancel",
		ReturnUrl:   "http://localhost:8080/api/payment/success",
	}

	// Create payment link
	paymentData, err := payos.CreatePaymentLink(paymentReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment link: %v", err)
	}

	// Prepare response
	resp := &response.PaymentLinkResponse{
		Success:     true,
		PayUrl:      paymentData.CheckoutUrl,
		OrderCode:   paymentData.OrderCode,
		Amount:      paymentData.Amount,
		Description: paymentData.Description,
	}

	return resp, nil
}
