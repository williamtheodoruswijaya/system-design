package event

import "time"

type UserEvent struct {
	UserID    int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Getter for UserID attributes implementing the contract
func (u *UserEvent) GetId() string {
	return string(u.UserID)
}
