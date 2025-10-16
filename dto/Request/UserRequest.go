package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	PhoneNumber    string `json:"phone,omitempty" validate:"required,min=10,max=10"`
	DoB 	 string `json:"dob" validate:"required"`
	Gender 	 string `json:"gender" `
}




// ------------check sauuu

// type UpdateUserRequest struct {
// 	Name   *string `json:"name,omitempty" validate:"omitempty"`
// 	Email  *string `json:"email,omitempty" validate:"omitempty,email"`
// 	Phone  *string `json:"phone,omitempty" validate:"omitempty"`
// 	Role   *string `json:"role,omitempty" validate:"omitempty,oneof=user admin driver"`
// 	Status *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive banned"`
// }


// type LoginRequest struct {
// 	Email    string `json:"email" validate:"required,email"`
// 	Password string `json:"password" validate:"required"`
// }

