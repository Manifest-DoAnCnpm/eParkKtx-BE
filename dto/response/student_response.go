package response

import (
	"time"

	"eParkKtx/entities"
)

type StudentResponse struct {
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	DoB         time.Time `json:"dob"`
	Gender      string    `json:"gender"`
	PhoneNumber string    `json:"phone_number"`
	School      string    `json:"school"`
	Room        string    `json:"room"`
}

func NewStudentResponse(student *entities.Student) *StudentResponse {
	if student == nil {
		return nil
	}
	return &StudentResponse{
		UserID:      student.UserID,
		Name:        student.User.Name,
		DoB:         student.User.DoB,
		Gender:      student.User.Gender,
		PhoneNumber: student.User.PhoneNumber,
		School:      student.School,
		Room:        student.Room,
	}
}
