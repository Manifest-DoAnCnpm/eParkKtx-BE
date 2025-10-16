package services

import (
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt" // for encrypte password
)




/*
	WARNING: CHANG PARAMATERS OF FUNCTION TO DTO LATER
*/






/*
	User service:
	- encrypte/decrypte password
	- check username
	- phonenumber
*/

type UserService struct {
	UserRepo *repositories.UserRepo
}


//------------------------Constructor---------------------------------------
func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}



//------------------------Method---------------------------------------

func (Userv *UserService) CreateUser(User *entities.User) error {

	if User.Name == "" || User.Password == "" || len(User.PhoneNumber) < 10 {
		log.SetPrefix("[UserService]: ")
		return errors.New("user name or password cannot be empty")
	}

	// check name
	if CheckUser, err := Userv.UserRepo.GetByName(User.Name); CheckUser != nil {
		log.SetPrefix("[UserService]: ")
		return err;
	}

	// encrypte before saving to db
	User.Password, _ = HashPassword(User.Password)

	Userv.UserRepo.CreateNewUser(User)

	return nil
}

// Hash password
func HashPassword(password string) (string, error) {

	// default cost có độ phức tạp là 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// copare password string with password-hash in db
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func (Userv *UserService) GetUserByName(User *entities.User) (*entities.User,error){

	ExistedUser, err := Userv.UserRepo.GetByName(User.Name);
	if  err != nil {
		return nil,err;
	}

	IsMatched := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(User.Password));

	if IsMatched != nil{
		return nil, IsMatched;
	}

	return ExistedUser, nil;
}

func (Userv *UserService) GetUserByID(ID string, Password string)  (*entities.User,error){

	ExistedUser, err := Userv.UserRepo.GetByID(ID);
	if  err != nil {
		return nil,err;
	}

	IsMatched := bcrypt.CompareHashAndPassword([]byte(ExistedUser.Password), []byte(Password));

	if IsMatched != nil{
		return nil, IsMatched;
	}

	return ExistedUser, nil;
}



