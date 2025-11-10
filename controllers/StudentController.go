package controllers

import (
	"eParkKtx/dto/Request"
	"eParkKtx/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentService *services.StudentService
}

// constructor
func NewStudentController(StudentService *services.StudentService) *StudentController {
	return &StudentController{
		StudentService: StudentService,
	}
}

// CreateStudent xử lý yêu cầu tạo mới sinh viên
// @Summary Tạo mới sinh viên
// @Description Tạo tài khoản sinh viên mới với thông tin được cung cấp
// @Tags students
// @Accept json
// @Produce json
// @Param student body request.CreateStudentRequest true "Thông tin sinh viên"
// @Success 201 {object} map[string]interface{} "Tạo thành công"
// @Failure 400 {object} map[string]interface{} "Dữ liệu không hợp lệ"
// @Failure 500 {object} map[string]interface{} "Lỗi server"
// @Router /students [post]
func (sc *StudentController) CreateStudent(c *gin.Context) {
	var req request.CreateStudentRequest

	// Validate và bind dữ liệu từ request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[StudentController] Invalid request data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	log.Printf("Received request: %+v", req)

	
	// Gọi service để tạo sinh viên
	if err := sc.StudentService.CreateStudent(req); err != nil {
		log.Printf("[StudentController] Failed to create student: %v", err)
		
		// Phân loại lỗi để trả về status code phù hợp
		switch err.Error() {
		case "username already exists":
			c.JSON(http.StatusConflict, gin.H{
				"success": false,
				"error":   "Username already exists",
			})
		case "invalid user data":
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Invalid user data",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Failed to create student",
			})
		}
		return
	}

	// Trả về response thành công
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Student created successfully",
	})
}
