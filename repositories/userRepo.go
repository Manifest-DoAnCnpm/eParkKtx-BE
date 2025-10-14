package repositories


import (
    "eParkKtx/entities"
	"eParkKtx/config"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository() *UserRepository {
    return &UserRepository{config.DB} // create user repository 
}

func (repo *UserRepository) CreateNewUser(user *entities.User) error {
    return repo.DB.Create(user).Error // tạo user trong db
}


func (r *UserRepository) GetByName(Name string) (*entities.User, error) {
    var user entities.User

    if err := r.DB.First(&user, "Name = ?", Name).Error; err != nil { // if <khai báo biến>; <condition>{}
        return nil, err
    }
    return &user, nil // trả về con trỏ tới user vừa tạo và nil -> ko lỗi
}
