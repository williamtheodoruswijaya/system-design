package messaging

import (
	"kafka-confluent-pzn/model"

	"github.com/twmb/franz-go/pkg/kgo"
)

type UserProducer struct {
	Producer *Producer[*model.UserEvent] 	// Gunakan generic type untuk UserEvent
}

func NewUserProducer(producer *kgo.Client) *UserProducer {
	return &UserProducer{
		Producer: &Producer[*model.UserEvent]{
			Producer: producer,
			Topic:    "users",
		},
	}
}