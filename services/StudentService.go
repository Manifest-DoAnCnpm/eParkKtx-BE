package services

import (
	request "eParkKtx/dto/Request"
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"log"
	"time"

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
	// Tạo user trước
	user := &entities.User{
		UserID:      uuid.New().String(),
		Name:        req.UserRequest.Name,
		Password:    req.UserRequest.Password, // Mật khẩu sẽ được hash trong UserService
		PhoneNumber: req.UserRequest.PhoneNumber,
		DoB:         parseDateOfBirth(req.UserRequest.DoB),
		Gender:      req.UserRequest.Gender,
	}

	// Tạo user
	if err := s.UserService.CreateUser(user); err != nil {
		log.Printf("[StudentService] Failed to create user: %v", err)
		return err
	}

	// Tạo student
	student := &entities.Student{
		UserID: user.UserID, // Giả sử UserID được tạo tự động bởi database
		School: req.School,
		Room:   req.Room,
	}

	// Lưu student vào database
	if err := s.StudentRepo.CreateNewStudent(student); err != nil {
		log.Printf("[StudentService] Failed to create student: %v", err)
		return err
	}

	return nil
}

// func (Sservice *StudentService) GetStudentByName(Student *entities.Student) (*entities.Student,error){

// 	ExistedStudent , err := Sservice.UserService.GetUserByName(&Student.User);

// 	if err != nil{
// 		return err;
// 	}

// }

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
