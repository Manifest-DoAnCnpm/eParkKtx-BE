package repositories


import (
    "eParkKtx/entities"

)


type ParkManagementRepo struct{
	UserRepo *UserRepo; // composition (nhúng userrepo)
}

// constructor
func NewParkManagementRepo(userRepo *UserRepo) *ParkManagementRepo {
    return &ParkManagementRepo{UserRepo: userRepo}
}

// create student
func (repo *StudentRepo) CreateNewParkManagement(ParkManagement *entities.ParkManagement) error {
    // tạo user trước
    if err := repo.UserRepo.CreateNewUser(&ParkManagement.User); err != nil {
        return err
    }

    return repo.UserRepo.DB.Create(ParkManagement).Error

}




