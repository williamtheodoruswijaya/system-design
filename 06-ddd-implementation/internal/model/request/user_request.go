package request

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Fullname string `json:"fullname" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ValidateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Fullname string `json:"fullname" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Profile  string `json:"profile"`
}
