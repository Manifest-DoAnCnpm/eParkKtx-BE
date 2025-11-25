package services

import (
	"eParkKtx/dto/request"
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"log"
	"time"
	"fmt"
	"github.com/google/uuid"
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
    user := &entities.User{
        UserID:      uuid.New().String(),
        Name:        req.UserRequest.Name,
        Password:    req.UserRequest.Password,
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


func (Sservice *StudentService) GetStudentByName(req request.GetStudentByNameRequest) (*entities.Student,error){

	ExistedUser , err := Sservice.UserService.GetUserByName(req.Name);
	

	if err != nil{
		return nil, err;
	}

	ExistedStu , err := Sservice.StudentRepo.GetByStudentID(ExistedUser.UserID)

	if err != nil{
		return nil,err;
	}

	return ExistedStu, nil;
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
