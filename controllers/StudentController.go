package controllers

import (
	request "eParkKtx/dto/Request"
	"eParkKtx/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentService services.StudentService;

}

//constructor
func NewStudentController(StudentService services.StudentService) *StudentController{
	return &StudentController{
		StudentService: StudentService,
	}
}

// API: POST /students
/*
	- bindjson: tự động xửa lí lỗi luôn != shouldbindjson mình tự handle lỗi
    
*/
func (sc *StudentController) CreateStudent(c *gin.Context) {
    var req request.CreateStudentRequest

    // Nhận dữ liệu JSON từ client
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Gọi service xử lý nghiệp vụ
    if err := sc.StudentService.CreateStudent(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Trả response cho client
    c.JSON(http.StatusOK, gin.H{
        "message": "Student created successfully",
    })
}