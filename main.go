package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	payos "github.com/payOSHQ/payos-lib-golang"
)

// === Thông tin sandbox PayOS (tạm) ===
const (
	PAYOS_CLIENT_ID    = "07622372-99f6-4a21-9376-52932d63d091"
	PAYOS_API_KEY      = "ac6155d0-0d3e-46c8-8c42-49975f7956d8"
	PAYOS_CHECKSUM_KEY = "fdd5d4c9d46d2b984e115a285b08b97243a919ebbbfae812bf2593206d2e324c"
)

// Tạo instance PayOS client
var client *Client

type Client struct {
	ClientID    string
	APIKey      string
	ChecksumKey string
}

func NewClient(clientID, apiKey, checksum string) *Client {
	return &Client{
		ClientID:    clientID,
		APIKey:      apiKey,
		ChecksumKey: checksum,
	}
}

// CreatePaymentLink is a minimal stub that returns a mocked response.
// Replace with real HTTP calls to PayOS API if needed.
func (c *Client) CreatePaymentLink(body payos.CheckoutRequestType) (interface{}, error) {
	resp := map[string]interface{}{
		"checkout_url": fmt.Sprintf("https://payos.example/checkout/%v", body.OrderCode),
		"order_code":   body.OrderCode,
		"amount":       body.Amount,
		"description":  body.Description,
	}
	return resp, nil
}

type CreateLinkRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

// API: POST /create-payment
func createPaymentLink(w http.ResponseWriter, r *http.Request) {
	var req CreateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Body gửi PayOS
	body := payos.CheckoutRequestType{
		OrderCode:   time.Now().Unix(),
		Amount:      int(req.Amount),
		Description: req.Description,
		CancelUrl:   "http://localhost:8080/payment/cancel",
		ReturnUrl:   "http://localhost:8080/payment/success",
	}

	// Gọi API PayOS
	res, err := client.CreatePaymentLink(body)
	if err != nil {
		http.Error(w, "Failed to create payment link", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func handleSuccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Thanh toán thành công ✅",
	})
}

func handleCancel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "cancel",
		"message": "Đã hủy thanh toán ❌",
	})
}

func main() {
	// ✅ Khởi tạo PayOS client đúng cách
	client = NewClient(PAYOS_CLIENT_ID, PAYOS_API_KEY, PAYOS_CHECKSUM_KEY)

	http.HandleFunc("/create-payment", createPaymentLink)
	http.HandleFunc("/payment/success", handleSuccess)
	http.HandleFunc("/payment/cancel", handleCancel)

	fmt.Println("🚀 Server chạy tại http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
