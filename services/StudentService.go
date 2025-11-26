package services

import (
	"eParkKtx/dto/request"
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"fmt"
	"log"
	"time"
)

/*
	WARNING: CHANGE PARAMATERS OF FUNCTION TO DTO LATER

*/

/*
	Student service:
	- inherited from userservice
*/

type StudentService struct {
	UserService *UserService // composition userservice into subclass service
	StudentRepo *repositories.StudentRepo
}

// ------------------------Constructor---------------------------------------
func NewStudentService(UserService *UserService, StudentRepo *repositories.StudentRepo) *StudentService {
	return &StudentService{UserService: UserService, StudentRepo: StudentRepo}
}

//------------------------Method---------------------------------------

// parseDateOfBirth chuyển đổi chuỗi ngày tháng sang time.Time
func parseDateOfBirth(dobString string) time.Time {
	// Định dạng: "2006-01-02"
	dob, err := time.Parse("2006-01-02", dobString)
	if err != nil {
		log.Printf("[StudentService] Error parsing date of birth: %v", err)
		return time.Now() // Trả về thời gian hiện tại nếu có lỗi
	}
	return dob
}

func (s *StudentService) CreateStudent(req request.CreateStudentRequest) error {
	log.Printf("[StudentService] Received create student request: %+v", req)

	// Tạo user trước
	hashpass, err := s.UserService.HashPassword(req.UserRequest.Password)
	if err != nil {
		return err
	}
	user := &entities.User{
		UserID:      req.UserRequest.CCCD, // Sử dụng CCCD làm UserID
		Name:        req.UserRequest.Name,
		Password:    hashpass,
		PhoneNumber: req.UserRequest.PhoneNumber,
		DoB:         parseDateOfBirth(req.UserRequest.DoB),
		Gender:      req.UserRequest.Gender,
	}
	log.Printf("[StudentService] Creating user with ID: %s", user.UserID)

	// Tạo user
	// if err := s.UserService.CreateUser(user); err != nil {
	//     log.Printf("[StudentService] Failed to create user: %v", err)
	//     return fmt.Errorf("failed to create user: %w", err)
	// }
	// log.Printf("[StudentService] Successfully created user: %s", user.UserID)

	// Tạo student
	student := &entities.Student{
		UserID: user.UserID,
		School: req.School,
		Room:   req.Room,
		User:   *user,
	}
	log.Printf("[StudentService] Creating student with user ID: %s", student.UserID)

	// Lưu student vào database
	if err := s.StudentRepo.CreateNewStudent(student); err != nil {
		log.Printf("[StudentService] Failed to create student: %v", err)
		return fmt.Errorf("failed to create student: %w", err)
	}

	log.Printf("[StudentService] Successfully created student with user ID: %s", student.UserID)
	return nil
}

func (Sservice *StudentService) GetStudentByName(req request.GetStudentByNameRequest) (*entities.Student, error) {

	ExistedUser, err := Sservice.UserService.GetUserByName(req.Name)

	if err != nil {
		return nil, err
	}

	ExistedStu, err := Sservice.StudentRepo.GetByStudentID(ExistedUser.UserID)

	if err != nil {
		return nil, err
	}

	return ExistedStu, nil
}

// RegisterVehicle registers a new vehicle for a student
func (s *StudentService) RegisterVehicle(req request.RegisterVehicleRequest) error {
	log.Printf("[StudentService] Registering vehicle: %+v", req)

	// Check if student exists
	_, err := s.StudentRepo.GetByStudentID(req.StudentID)
	if err != nil {
		log.Printf("[StudentService] Student not found: %v", err)
		return fmt.Errorf("student not found")
	}

	// Check if vehicle with this number plate already exists
	var existingVehicle entities.Vehicle
	if err := s.StudentRepo.UserRepo.DB.Where("number_plate = ?", req.NumberPlate).First(&existingVehicle).Error; err == nil {
		log.Printf("[StudentService] Vehicle with number plate %s already exists", req.NumberPlate)
		return fmt.Errorf("vehicle with this number plate already exists")
	}

	// Create new vehicle
	vehicle := &entities.Vehicle{
		NumberPlate:      req.NumberPlate,
		VehicleType:      req.VehicleType,
		Color:            req.Color,
		RegisterDate:     time.Now(),
		StudentID:        req.StudentID,
		ParkManagementID: req.ParkManagementID,
	}

	// Save to database
	if err := s.StudentRepo.UserRepo.DB.Create(vehicle).Error; err != nil {
		log.Printf("[StudentService] Failed to register vehicle: %v", err)
		return fmt.Errorf("failed to register vehicle: %w", err)
	}

	log.Printf("[StudentService] Successfully registered vehicle %s for student %s",
		req.NumberPlate, req.StudentID)
	return nil
}

// func (Userv *UserService) GetStudentByID(ID string, Password string)  (*entities.User,error){

// 	ExistedUser, err := Userv.UserRepo.GetByID(ID);
// 	if  err != nil {
// 		return nil,err;
// 	}

// 	IsMatched := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(Password));

// 	if IsMatched != nil{
// 		return nil, IsMatched;
// 	}

// 	return ExistedUser, nil;
// }
