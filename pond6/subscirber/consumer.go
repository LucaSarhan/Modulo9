package main

import (
	"fmt"
	"os"
	"bytes"
	"net/http"

	godotenv "github.com/joho/godotenv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func PostData(message string) {
	const url string = "http://localhost:8000/sensors"

	req, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(message)))
	if err != nil {
		fmt.Println("Error making the request:", err)
	}

	if req.StatusCode == 200 {
		fmt.Println("Data sent successfully")
	} else {
		fmt.Printf("Sending data: %s\n", req.Status)
	}
	fmt.Println("")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
		"group.id":          os.Getenv("CLUSTER_ID"),
		"auto.offset.reset": "earliest",
		"security.protocol": "SASL_SSL",
		"sasl.mechanisms":   "PLAIN",
		"sasl.username":     os.Getenv("API_KEY"),
		"sasl.password":     os.Getenv("API_SECRET"),
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	fmt.Println("Consumer started")

	topic := "kafka"
	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if msg != nil && err == nil {
			msgValue := string(msg.Value)
			fmt.Printf("Received message: %s\n", msgValue)
			PostData(msgValue)
		} else if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}