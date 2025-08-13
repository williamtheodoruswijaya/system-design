package messaging

import (
	"06-ddd-implementation/internal/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
)

// step 1: buat struct Producer dengan generic type T yang merupakan model.Event -> cara pakai-nya Producer[UserEvent]
/*
	Keuntungannya adalah type safety (keamanan tipe).
	Kita tidak akan bisa secara tidak sengaja mengirim OrderCreatedEvent
	menggunakan producer yang dirancang khusus untuk UserRegisteredEvent.

	Ibaratnya seperti ini:
	- Jika kita punya `Producer[UserEvent]`, kita hanya bisa mengirim `UserEvent`.
	- Jika kita punya `Producer[OrderEvent]`, kita hanya bisa mengirim `OrderEvent`.
	Jadi, kita tidak akan pernah salah mengirim event yang salah ke topic yang salah.
*/
type Producer[T model.Event] struct {
	Producer *kgo.Client
	Topic    string
}

// step 2: buat function buat dapet topic dari producer (getter ini optional, tapi bisa berguna untuk kedepannya, jadi ywd taro aja)
func (producer *Producer[T]) GetTopic() string {
	return producer.Topic
}

// step 3: buat function untuk publish event ke topic yang sudah ditentukan
func (producer *Producer[T]) Publish(event T) error {
	// step 4: convert event ke json
	value, err := json.Marshal(event)
	if err != nil {
		fmt.Printf("Error marshalling event: %v\n", err)
		return err
	}

	// step 5: buat record untuk dikirim ke topic
	record := &kgo.Record{
		Topic: producer.Topic,
		Key:   []byte(event.GetId()), // Gunakan ID dari event sebagai key
		Value: value,
	}

	// step 6: publish record ke topic
	ctx := context.Background()

	// Gunakan ProduceSync untuk memastikan bahwa message sudah berhasil dikirim sebelum melanjutkan
	// Ini akan mengembalikan error jika terjadi kesalahan saat mengirim message
	err = producer.Producer.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		fmt.Printf("Error producing record: %v\n", err)
		return err
	}

	// step 7: return nil jika tidak ada error
	return nil
}
