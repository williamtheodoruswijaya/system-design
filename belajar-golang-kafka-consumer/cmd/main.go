package main

import (
	"belajar-golang-kafka-consumer/config"
	"belajar-golang-kafka-consumer/utils"
	"fmt"
)

func main() {
	// Test the Kafka consumer and producer setup

	// step 0: Create a topic if it doesn't exist
	utils.CreateTopic("test-topic-belajar-kafka")

	// step 1: Create a new Kafka producer
	producer := config.NewKafkaProducer("test-topic-belajar-kafka")

	// step 2: Create a new Kafka consumer
	consumer := config.NewKafkaConsumer("test-topic-belajar-kafka", "test-group-0")

	// step 3: Write a message to the Kafka topic
	err := utils.WriteMessage(producer, []byte("key-1"), []byte("Hello, Kafka!"))
	if err != nil {
		panic(err)
	}

	// step 4: Read a message from the Kafka topic
	msg, err := utils.ReadMessage(consumer)
	if err != nil {
		panic(err)
	}

	// Print the message
	fmt.Printf("Received message: Key=%s, Value=%s\n", string(msg.Key), string(msg.Value))

	// Close the producer and consumer
	producer.Close()
	consumer.Close()
}