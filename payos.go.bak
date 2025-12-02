package main

// import (
// 	"encoding/json"
// 	"fmt"
	
// 	"net/http"
// 	"time"

// 	payos "github.com/payOSHQ/payos-lib-golang"
// )

// // === Thông tin PayOS (sử dụng credentials thật của bạn) ===
// const (
// 	PAYOS_CLIENT_ID    = "07622372-99f6-4a21-9376-52932d63d091"
// 	PAYOS_API_KEY      = "ac6155d0-0d3e-46c8-8c42-49975f7956d8"
// 	PAYOS_CHECKSUM_KEY = "fdd5d4c9d46d2b984e115a285b08b97243a919ebbbfae812bf2593206d2e324c"
// )

// type CreateLinkRequest struct {
// 	Amount      int64  `json:"amount"`
// 	Description string `json:"description"`
// }

// // Enable CORS for FE local testing
// func enableCORS(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.Header().Set("Access-Control-Allow-Credentials", "true")
// }

// func withCORS(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		enableCORS(w, r)
// 		if r.Method == "OPTIONS" {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}
// 		handler(w, r)
// 	}
// }

// // API: POST /create-payment
// func createPaymentLink(w http.ResponseWriter, r *http.Request) {
// 	var req CreateLinkRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
// 		return
// 	}

// 	payos.Key(PAYOS_CLIENT_ID, PAYOS_API_KEY, PAYOS_CHECKSUM_KEY)
// 	body := payos.CheckoutRequestType{
// 		OrderCode:   time.Now().Unix(),
// 		Amount:      int(req.Amount),
// 		Description: req.Description,
// 		CancelUrl:   "http://localhost:8080/payment/cancel",
// 		ReturnUrl:   "http://localhost:8080/payment/success",
// 	}

// 	data, err := payos.CreatePaymentLink(body)
// 	if err != nil {
// 		fmt.Println("PayOS error:", err)
// 		http.Error(w, "Failed to create payment link", http.StatusInternalServerError)
// 		return
// 	}

// 	resp := map[string]interface{}{
// 		"payUrl":      data.CheckoutUrl,
// 		"order_code":  data.OrderCode,
// 		"amount":      data.Amount,
// 		"description": data.Description,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp)
// }

// func handleSuccess(w http.ResponseWriter, r *http.Request) {
// 	// orderId := r.URL.Query().Get("orderId")
// 	http.Redirect(w, r, "http://localhost:3000/payment-success", http.StatusSeeOther)
// }

// func handleCancel(w http.ResponseWriter, r *http.Request) {
// 	// orderId := r.URL.Query().Get("orderId")
// 	http.Redirect(w, r, "http://localhost:3000/payment-cancel", http.StatusSeeOther)
// }

// // func main() {
// // 	http.HandleFunc("/create-payment", withCORS(createPaymentLink))
// // 	http.HandleFunc("/payment/success", withCORS(handleSuccess))
// // 	http.HandleFunc("/payment/cancel", withCORS(handleCancel))

// // 	fmt.Println("🚀 Server chạy tại http://localhost:8080")
// // 	log.Fatal(http.ListenAndServe(":8080", nil))
// // }
