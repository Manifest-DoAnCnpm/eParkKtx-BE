package repositories

import (
	"eParkKtx/config"
	"eParkKtx/entities"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	NewUserRepository()
	CreateNewUser(user *entities.User) error
	GetByName(Name string) (*entities.User, error)
	Delete(ID string) error // Thêm hàm delete
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepo {
	return &UserRepo{config.DB} // create user repository (constructor)
}

func (repo *UserRepo) CreateNewUser(user *entities.User) error {
	    // Kiểm tra user đã tồn tại chưa (sửa lại câu query)
    var existingUser entities.User
    if err := repo.DB.Where("name = ?", user.Name).First(&existingUser).Error; err == nil {
        return errors.New("user with this name already exists")
    }
	return repo.DB.Create(user).Error // tạo user trong db
}

func (r *UserRepo) GetByID(id string) (*entities.User, error) {
    var user entities.User
    // Sửa lại câu query để dùng đúng tên cột
    if err := r.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}


func (r *UserRepo) GetByName(Name string) (*entities.User, error) {
	var user entities.User;

	if err := r.DB.First(&user, "name = ?", Name).Error; err != nil { // if <khai báo biến>; <condition>{}
		return nil, err;
	}
	return &user, nil; // trả về con trỏ tới user vừa tạo và nil -> ko lỗi
}

// Get all users
func (r *UserRepo) GetAll() ([]entities.User, error) {
	var users []entities.User;
	if err := r.DB.Find(&users).Error; err != nil { // nil là null
		return nil, err;
	}
	return users, nil;
}

func (repo *UserRepo) Update(newU *entities.User) error {

	if newU.UserID == "" {
		return errors.New("userID cannot be empty");
	}
	var user entities.User
	if err := repo.DB.First(&user, "user_id = ?", newU.UserID).Error; err != nil {
		return errors.New("undefined user");
	}

	repo.DB.Save(newU);
	return nil
}

func (repo *UserRepo) Delete(ID string) error {
	// Tìm và xóa user dựa trên ID
    result := repo.DB.Delete(&entities.User{}, "user_id = ?", ID)
    
    // Nếu có lỗi trong quá trình xóa (trừ trường hợp không tìm thấy)
    if result.Error != nil {
        return result.Error
    }
    
    // Tùy chọn: Kiểm tra xem có bản ghi nào bị xóa hay không (nếu cần xác nhận user tồn tại)
    if result.RowsAffected == 0 {
        return errors.New("user not found")
    }
	
    return nil
}

