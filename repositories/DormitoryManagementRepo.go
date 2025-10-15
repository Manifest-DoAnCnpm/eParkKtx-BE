package repositories


import (
    "eParkKtx/entities"

)


type DormitoryManagementRepo struct{
	UserRepo *UserRepo; // composition (nhúng userrepo)
}

// constructor
func NewDormitoryManagementRepo(userRepo *UserRepo) *DormitoryManagementRepo {
    return &DormitoryManagementRepo{UserRepo: userRepo}
}

// create student
func (repo *StudentRepo) CreateNewDormitoryManagement(DormitoryManagement *entities.DormitoryManagement) error {
    // tạo user trước
    if err := repo.UserRepo.CreateNewUser(&DormitoryManagement.User); err != nil {
        return err
    }

    return repo.UserRepo.DB.Create(DormitoryManagement).Error

}




