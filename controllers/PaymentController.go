package controllers

import (
	"eParkKtx/dto/request"
	"eParkKtx/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	payOSService *services.PayOSService
}

func NewPaymentController(payOSService *services.PayOSService) *PaymentController {
	return &PaymentController{
		payOSService: payOSService,
	}
}

// CreatePaymentLink handles the creation of a payment link
// @Summary Tạo link thanh toán
// @Description Tạo link thanh toán thông qua PayOS
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body request.CreatePaymentLinkRequest true "Thông tin thanh toán"
// @Success 200 {object} response.PaymentLinkResponse "Tạo link thanh toán thành công"
// @Failure 400 {object} map[string]interface{} "Dữ liệu không hợp lệ"
// @Failure 500 {object} map[string]interface{} "Lỗi server"
// @Router /api/payment/create [post]
func (pc *PaymentController) CreatePaymentLink(c *gin.Context) {
	var req request.CreatePaymentLinkRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[PaymentController] Invalid request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Call service to create payment link
	paymentLink, err := pc.payOSService.CreatePaymentLink(req)
	if err != nil {
		log.Printf("[PaymentController] Failed to create payment link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create payment link",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, paymentLink)
}

// HandlePaymentSuccess handles successful payment callback
// @Summary Xử lý khi thanh toán thành công
// @Description Chuyển hướng khi thanh toán thành công
// @Tags payment
// @Success 303 "Redirect to success page"
// @Router /api/payment/success [get]
func (pc *PaymentController) HandlePaymentSuccess(c *gin.Context) {
	// In a real application, you might want to verify the payment here
	// before redirecting
	c.Redirect(http.StatusSeeOther, "http://localhost:3000/payment-success")
}

// HandlePaymentCancel handles cancelled payment callback
// @Summary Xử lý khi hủy thanh toán
// @Description Chuyển hướng khi người dùng hủy thanh toán
// @Tags payment
// @Success 303 "Redirect to cancel page"
// @Router /api/payment/cancel [get]
func (pc *PaymentController) HandlePaymentCancel(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "http://localhost:3000/payment-cancel")
}
