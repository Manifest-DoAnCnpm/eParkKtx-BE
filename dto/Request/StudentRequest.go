package request

type CreateStudentRequest struct {
	UserRequest CreateUserRequest;
	School        string `json:"school" validate:"required"`
	Room    string `json:"room" validate:"required"`

}

// constructor
func CreateStudentRequestInitialize(
	name string,
	password string,
	phoneNumber string,
	dob string,
	gender string,
	school string,
	room string,
) CreateStudentRequest {
	return CreateStudentRequest{
		UserRequest: CreateUserRequestInitialize(name, password,phoneNumber,dob,gender),
		School: school,
		Room:   room,
	}
}




