package services

import (
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"fmt"
	"log"
)

type ParkManagementService struct {
	ParkManagementRepo *repositories.ParkManagementRepo
}

// Constructor
func NewParkManagementService(parkManagementRepo *repositories.ParkManagementRepo) *ParkManagementService {
	return &ParkManagementService{
		ParkManagementRepo: parkManagementRepo,
	}
}

// GetAllVehiclesWithStudents retrieves all vehicles with their associated student information
func (s *ParkManagementService) GetAllVehiclesWithStudents() ([]entities.Vehicle, error) {
	log.Println("[ParkManagementService] Getting all vehicles with student information")

	var vehicles []entities.Vehicle

	// Preload the Student and nested User information for each Vehicle
	err := s.ParkManagementRepo.UserRepo.DB.
		Preload("Student").
		Preload("Student.User").
		Find(&vehicles).Error

	if err != nil {
		log.Printf("[ParkManagementService] Error getting vehicles with students: %v", err)
		return nil, fmt.Errorf("failed to get vehicles with students: %w", err)
	}

	log.Printf("[ParkManagementService] Successfully retrieved %d vehicles with student information", len(vehicles))
	return vehicles, nil
}
