package request

// RegisterVehicleRequest represents the request body for registering a vehicle
type RegisterVehicleRequest struct {
	StudentID        string `json:"student_id" binding:"required"`
	NumberPlate      string `json:"number_plate" binding:"required"`
	VehicleType      string `json:"vehicle_type" binding:"required"`
	Color            string `json:"color" binding:"required"`
	ParkManagementID string `json:"park_management_id" binding:"required"`
}
