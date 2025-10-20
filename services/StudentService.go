package services

import (
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"log"
)






/*
	WARNING: CHANG PARAMATERS OF FUNCTION TO DTO LATER
	
*/






/*
	Student service:
	- inherited from userservice
*/

type StudentService struct {
	UserService *UserService; // composition userservice into subclass service
	StudentRepo *repositories.StudentRepo;

}

//------------------------Constructor---------------------------------------
func NewStudentService(UserService *UserService, StudentRepo *repositories.StudentRepo) *StudentService {
	return &StudentService{UserService: UserService, StudentRepo: StudentRepo}
}

//------------------------Method---------------------------------------

func (Sservice *StudentService) CreateStudent(Student *entities.Student) error {

	Sservice.UserService.CreateUser(&Student.User);

	if err:= Sservice.StudentRepo.CreateNewStudent(Student); err != nil{
		log.SetPrefix("[StudentService]");
		return err;
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







