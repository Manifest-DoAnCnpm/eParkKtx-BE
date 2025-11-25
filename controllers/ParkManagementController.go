package controllers

import (
	"eParkKtx/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ParkManagementController struct {
	ParkManagementService *services.ParkManagementService
}

// Constructor
func NewParkManagementController(parkManagementService *services.ParkManagementService) *ParkManagementController {
	return &ParkManagementController{
		ParkManagementService: parkManagementService,
	}
}

// GetAllVehiclesWithStudents retrieves all vehicles with their associated student information
// @Summary Lấy danh sách tất cả xe đã đăng ký kèm thông tin sinh viên
// @Description Trả về danh sách tất cả xe đã đăng ký cùng với thông tin sinh viên tương ứng
// @Tags park-management
// @Produce json
// @Success 200 {object} map[string]interface{} "Thành công"
// @Failure 500 {object} map[string]interface{} "Lỗi server"
// @Router /api/park-management/vehicles [get]
func (pc *ParkManagementController) GetAllVehiclesWithStudents(c *gin.Context) {
	log.Println("[ParkManagementController] Getting all vehicles with student information")

	// Call service to get all vehicles with student info
	vehicles, err := pc.ParkManagementService.GetAllVehiclesWithStudents()
	if err != nil {
		log.Printf("[ParkManagementController] Error getting vehicles with students: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to get vehicles with student information",
			"details": err.Error(),
		})
		return
	}

	// Create response DTO
	type VehicleWithStudentResponse struct {
		NumberPlate  string `json:"number_plate"`
		VehicleType  string `json:"vehicle_type"`
		Color        string `json:"color"`
		RegisterDate string `json:"register_date"`
		Student      struct {
			UserID      string `json:"user_id"`
			Name        string `json:"name"`
			PhoneNumber string `json:"phone_number"`
			School      string `json:"school"`
			Room        string `json:"room"`
		} `json:"student"`
	}

	// Map entities to response DTO
	response := make([]VehicleWithStudentResponse, len(vehicles))
	for i, v := range vehicles {
		response[i] = VehicleWithStudentResponse{
			NumberPlate:  v.NumberPlate,
			VehicleType:  v.VehicleType,
			Color:        v.Color,
			RegisterDate: v.RegisterDate.Format("2006-01-02 15:04:05"),
		}

		// Add student info if available
		if v.Student.UserID != "" {
			response[i].Student = struct {
				UserID      string `json:"user_id"`
				Name        string `json:"name"`
				PhoneNumber string `json:"phone_number"`
				School      string `json:"school"`
				Room        string `json:"room"`
			}{
				UserID:      v.Student.UserID,
				Name:        v.Student.User.Name,
				PhoneNumber: v.Student.User.PhoneNumber,
				School:      v.Student.School,
				Room:        v.Student.Room,
			}
		}
	}

	log.Printf("[ParkManagementController] Successfully retrieved %d vehicles with student information", len(vehicles))

	// Return the list of vehicles with student info
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}
