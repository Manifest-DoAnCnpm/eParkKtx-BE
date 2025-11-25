package routes

import (
	"eParkKtx/controllers"

	"github.com/gin-gonic/gin"
)

// SetupStudentRoutes thiết lập các routes cho student
func AuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login-cccd", authController.LoginCCCD)
		}
	}
}
