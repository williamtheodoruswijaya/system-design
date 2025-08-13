package messaging

import (
	"06-ddd-implementation/internal/model/event"

	"github.com/twmb/franz-go/pkg/kgo"
)

type UserProducer struct {
	Producer *Producer[*event.UserEvent] // Gunakan generic type untuk UserEvent
}

func NewUserProducer(producer *kgo.Client) *UserProducer {
	return &UserProducer{
		Producer: &Producer[*event.UserEvent]{
			Producer: producer,
			Topic:    "users",
		},
	}
}
