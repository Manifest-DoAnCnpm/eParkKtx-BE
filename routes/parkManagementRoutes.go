package routes

import (
	"eParkKtx/controllers"

	"github.com/gin-gonic/gin"
)

// SetupParkManagementRoutes sets up the routes for park management
func SetupParkManagementRoutes(r *gin.Engine, parkManagementController *controllers.ParkManagementController) {
	// Group park management routes under /api/park-management
	parkManagementGroup := r.Group("/api/park-management")
	{
		// Get all vehicles with student information
		parkManagementGroup.GET("/vehicles", parkManagementController.GetAllVehiclesWithStudents)
	}
}
