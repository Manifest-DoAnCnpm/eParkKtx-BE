package controllers

import (
	"eParkKtx/dto/request"
	"eParkKtx/services"
	"log"
	"net/http"
	"eParkKtx/dto/response"

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

func (sc *StudentController) GetStudentByName(c *gin.Context) {
    var req request.GetStudentByNameRequest

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

    // Gọi service để lấy thông tin sinh viên
    student, err := sc.StudentService.GetStudentByName(req)
    if err != nil {
        log.Printf("[StudentController] Failed to get student: %v", err)
        
        if err.Error() == "student not found" {
            c.JSON(http.StatusNotFound, gin.H{
                "success": false,
                "error":   "Student not found",
            })
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{
                "success": false,
                "error":   "Failed to get student information",
            })
        }
        return
    }

    // Chuyển đổi sang response DTO
    studentResponse := response.NewStudentResponse(student)

    // Trả về response thành công
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    studentResponse,
    })
}

// RegisterVehicle xử lý yêu cầu đăng ký xe cho sinh viên
// @Summary Đăng ký xe cho sinh viên
// @Description Đăng ký thông tin xe mới cho sinh viên
// @Tags students
// @Accept json
// @Produce json
// @Param vehicle body request.RegisterVehicleRequest true "Thông tin đăng ký xe"
// @Success 200 {object} map[string]interface{} "Đăng ký thành công"
// @Failure 400 {object} map[string]interface{} "Dữ liệu không hợp lệ"
// @Failure 404 {object} map[string]interface{} "Không tìm thấy sinh viên"
// @Failure 500 {object} map[string]interface{} "Lỗi server"
// @Router /students/vehicles [post]
func (sc *StudentController) RegisterVehicle(c *gin.Context) {
	var req request.RegisterVehicleRequest

	// Validate và bind dữ liệu từ request
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("[StudentController] Invalid vehicle registration data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	log.Printf("[StudentController] Registering vehicle for student: %s", req.StudentID)

	// Gọi service để đăng ký xe
	err := sc.StudentService.RegisterVehicle(req)
	if err != nil {
		log.Printf("[StudentController] Error registering vehicle: %v", err)
		statusCode := http.StatusInternalServerError
		errMsg := "Failed to register vehicle"
		
		if err.Error() == "student not found" {
			statusCode = http.StatusNotFound
			errMsg = "Student not found"
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   errMsg,
			"details": err.Error(),
		})
		return
	}

	log.Printf("[StudentController] Successfully registered vehicle for student: %s", req.StudentID)

	// Trả về thông báo thành công
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Đăng ký xe thành công",
		"data":    req,
	})
}
