package response

import "time"

type CreateUserResponse struct {
	UserID    int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ValidateUserResponse struct {
	Token string `json:"token"`
}
