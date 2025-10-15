package repositories

import (
	"eParkKtx/config"
	"eParkKtx/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	NewUserRepository()
	CreateNewUser(user *entities.User) error
	GetByName(Name string) (*entities.User, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepo {
	return &UserRepo{config.DB} // create user repository (constructor)
}

func (repo *UserRepo) CreateNewUser(user *entities.User) error {
	return repo.DB.Create(user).Error // tạo user trong db
}

func (r *UserRepo) GetByID(ID string) (*entities.User, error) {
	var user entities.User

	if err := r.DB.First(&user, "UserID = ?", ID).Error; err != nil { // if <khai báo biến>; <condition>{}
		return nil, err
	}
	return &user, nil // trả về con trỏ tới user vừa tạo và nil -> ko lỗi
}

func (r *UserRepo) GetByName(Name string) (*entities.User, error) {
	var user entities.User

	if err := r.DB.First(&user, "Name = ?", Name).Error; err != nil { // if <khai báo biến>; <condition>{}
		return nil, err
	}
	return &user, nil // trả về con trỏ tới user vừa tạo và nil -> ko lỗi
}

// Get all users
func (r *UserRepo) GetAll() ([]entities.User, error) {
	var users []entities.User
	if err := r.DB.Find(&users).Error; err != nil { // nil là null
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) Update(newU *entities.User) error {

	var user entities.User
	if err := repo.DB.First(&user, "UserID = ?", newU.UserID).Error; err != nil {
		return err
	}

	repo.DB.Save(newU)
	return nil
}
