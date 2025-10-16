package repositories


import (
    "eParkKtx/entities"
   
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
    // tạo user trước
    if err := repo.UserRepo.CreateNewUser(&student.User); err != nil {
        return err
    }

    return repo.UserRepo.DB.Create(student).Error

}


