package utils

import (
	"06-ddd-implementation/internal/entity"
	"06-ddd-implementation/internal/model/event"
	"06-ddd-implementation/internal/model/response"
)

func ConvertUserResponse(user *entity.User) *response.CreateUserResponse {
	return &response.CreateUserResponse{
		UserID:    user.ID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ConvertUserEvent(user *response.CreateUserResponse) *event.UserEvent {
	return &event.UserEvent{
		UserID:    user.UserID,
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
