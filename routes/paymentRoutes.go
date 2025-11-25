package routes

import (
	"eParkKtx/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(router *gin.Engine, paymentController *controllers.PaymentController) {
	paymentGroup := router.Group("/api/payment")
	{
		paymentGroup.POST("/create", paymentController.CreatePaymentLink)
		paymentGroup.GET("/success", paymentController.HandlePaymentSuccess)
		paymentGroup.GET("/cancel", paymentController.HandlePaymentCancel)
	}
}
