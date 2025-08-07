package utils

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

// CreateTopic creates a new Kafka topic with the specified name. By default, by creating a producer, the topic will be created automatically if it does not exist.
func CreateTopic(topic string) {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err)
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err)
	}
	defer controllerConn.Close()

	topicConfig := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfig...)
	if err != nil {
		panic(err)
	}
}

// WriteMessage writes a message to the specified Kafka topic using the provided writer.
func WriteMessage(writer *kafka.Writer, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	return writer.WriteMessages(context.Background(), msg)
}

// ReadMessage reads a message from the Kafka topic using the provided reader and prints it to the console.
func ReadMessage(reader *kafka.Reader) (kafka.Message, error) {
	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		return kafka.Message{}, err
	} else {
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
	return msg, nil
}