package messaging

import (
	"06-ddd-implementation/internal/model/event"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

type UserConsumer interface {
	Consume(record *kgo.Record) error
}

type UserConsumerImpl struct {
}

func NewUserConsumer() UserConsumer {
	return &UserConsumerImpl{}
}

func (userConsumer *UserConsumerImpl) Consume(record *kgo.Record) error {
	// step 1: convert record value to model user event (biar bentuknya sesuai)
	userEvent := event.UserEvent{}
	err := json.Unmarshal(record.Value, &userEvent)
	if err != nil {
		fmt.Printf("Error unmarshalling record value: %v\n", err)
		return err
	}

	// step 2: print user event (untuk tutorial/debug purpose saja) bisa loncat ke step 3 kalau app-nya udah ke defined
	fmt.Printf("Received user event: UserID=%s, Email=%s, Fullname=%d, CreatedAt=%d\n",
		userEvent.UserID, userEvent.Email, userEvent.Fullname, userEvent.CreatedAt)

	// step 3: process user event (disini bisa ditambahkan logic untuk menyimpan ke database atau melakukan proses lainnya)
	// misalnya simpan ke database atau melakukan proses lainnya
	// ...

	// step 4: return nil jika tidak ada error
	return nil
}
