package repositories

import (
	"eParkKtx/entities"
	//    "errors"
	"fmt"
)

type StudentRepo struct {
	UserRepo *UserRepo // composition (nhúng userrepo)
}

// constructor
func NewStudentRepo(userRepo *UserRepo) *StudentRepo {
	return &StudentRepo{UserRepo: userRepo}
}

// create student
func (repo *StudentRepo) CreateNewStudent(student *entities.Student) error {
	// Tạo user trước
	if err := repo.UserRepo.CreateNewUser(&student.User); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Gán lại user_id từ user vừa tạo
	student.UserID = student.User.UserID

	// Tạo student
	if err := repo.UserRepo.DB.Create(student).Error; err != nil {
		return fmt.Errorf("failed to create student: %w", err)
	}

	return nil
}

func (r *StudentRepo) GetByStudentID(userID string) (*entities.Student, error) {
	var student entities.Student
	if err := r.UserRepo.DB.Preload("User").Where("user_id = ?", userID).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
