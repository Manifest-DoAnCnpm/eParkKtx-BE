package routes

import (
	"eParkKtx/controllers"
	"eParkKtx/middlewares"
	"eParkKtx/services"
	"github.com/gin-gonic/gin"
)

// SetupStudentRoutes thiết lập các routes cho student
func SetupStudentRoutes(router *gin.Engine, studentController *controllers.StudentController, userService *services.UserService) {

	authMiddleware := middlewares.NewAuthMiddleware(userService)
	api := router.Group("/api")
	{
		students := api.Group("/students")
		{
			students.POST("", studentController.CreateStudent)
			students.POST("/search", authMiddleware.AuthRequired(), studentController.GetStudentByName)
			students.POST("/vehicles", authMiddleware.AuthRequired(), studentController.RegisterVehicle)

			// Có thể thêm các endpoint khác ở đây

			// students.GET("", studentController.GetStudents)
			// students.GET("/:id", studentController.GetStudentByID)
			// students.PUT("/:id", studentController.UpdateStudent)
			// students.DELETE("/:id", studentController.DeleteStudent)
		}
	}
}
