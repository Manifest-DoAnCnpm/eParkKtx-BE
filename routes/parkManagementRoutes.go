package routes

import (
	"eParkKtx/controllers"
	"eParkKtx/services"

	"github.com/gin-gonic/gin"
)

// SetupParkManagementRoutes sets up the routes for park management
func SetupParkManagementRoutes(r *gin.Engine, parkManagementController *controllers.ParkManagementController, userService *services.UserService) {
	// Group park management routes under /api/park-management
	parkManagementGroup := r.Group("/api/park-management")
	{
		// Get all vehicles with student information
		parkManagementGroup.GET("/vehicles", parkManagementController.GetAllVehiclesWithStudents)
	}
}
