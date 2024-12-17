package types

// input
type CreateUserAndStudent struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
	School string `json:"school" validate:"required"`
}

// response

type CreateUserAndStudentRes struct {
	UserId    string `json:"userId" validate:"required"`
	StudentId string `json:"studentId" validate:"required"`
}
