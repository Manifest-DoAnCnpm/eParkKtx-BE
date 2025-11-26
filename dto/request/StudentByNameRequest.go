package request

type GetStudentByNameRequest struct {
	Name string `json:"name" validate:"required"`

}

// constructor
func GetStudentByNameRequestInitialize(
	name string,

) GetStudentByNameRequest {
	return GetStudentByNameRequest{
		Name: name,
	}
}




