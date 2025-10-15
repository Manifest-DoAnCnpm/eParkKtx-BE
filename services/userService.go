package services

import (
	"eParkKtx/entities"
	"eParkKtx/repositories"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt" // for encrypte password
)

/*
	User service:
	- encrypte/decrypte password
	- check username
	- phonenumber
*/

type UserService struct {
	UserRepo *repositories.UserRepo
}

// constructor
func NewUserService(userRepo *repositories.UserRepo) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (Userv *UserService) CreateUser(User *entities.User) error {

	if User.Name == "" || User.Password == "" || len(User.PhoneNumber) < 10 {
		log.SetPrefix("[UserService]: ")
		return errors.New("user name or password cannot be empty")
	}

	// check name
	if CheckUser, _ := Userv.UserRepo.GetByName(User.Name); CheckUser != nil {
		log.SetPrefix("[UserService]: ")
		return errors.New("user name already exists")
	}

	// encrypte before saving to db
	User.Password, _ = Userv.HashPassword(User.Password)

	Userv.UserRepo.CreateNewUser(User)

	return nil
}

// Hash password
func (Userv *UserService) HashPassword(password string) (string, error) {

	// default cost có độ phức tạp là 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// copare password string with password-hash in db
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
