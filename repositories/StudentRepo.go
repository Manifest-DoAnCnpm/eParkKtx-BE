package repositories


import (
    "eParkKtx/entities"
   "errors"
)


type StudentRepo struct{
	UserRepo *UserRepo; // composition (nhúng userrepo)
}

// constructor
func NewStudentRepo(userRepo *UserRepo) *StudentRepo {
    return &StudentRepo{UserRepo: userRepo}
}

// create student
func (repo *StudentRepo) CreateNewStudent(student *entities.Student) error {

    // kiểm tra user có tồn tại chưa
	_, err := repo.UserRepo.GetByID(student.User.UserID)
	if err == nil {
		return errors.New("user already exists")
	}

    // tạo user trước
    if err := repo.UserRepo.CreateNewUser(&student.User); err != nil {
        return err
    }

    return repo.UserRepo.DB.Create(student).Error

}


