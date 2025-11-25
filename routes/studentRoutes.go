package routes

import (
	"eParkKtx/controllers"

	"github.com/gin-gonic/gin"
)

// SetupStudentRoutes thiết lập các routes cho student
func SetupStudentRoutes(router *gin.Engine, studentController *controllers.StudentController) {

	api := router.Group("/api")
	{
		students := api.Group("/students")
		{
			students.POST("", studentController.CreateStudent)
			students.POST("/search", studentController.GetStudentByName)

			// Có thể thêm các endpoint khác ở đây

			// students.GET("", studentController.GetStudents)
			// students.GET("/:id", studentController.GetStudentByID)
			// students.PUT("/:id", studentController.UpdateStudent)
			// students.DELETE("/:id", studentController.DeleteStudent)
		}
	}
}
