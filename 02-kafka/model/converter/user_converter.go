package converter

import (
	"kafka-confluent-pzn/entity"
	"kafka-confluent-pzn/model"
)

func UserToEvent(user *entity.User) *model.UserEvent {
	return &model.UserEvent{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}