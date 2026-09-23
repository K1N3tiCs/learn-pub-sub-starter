package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

const brokerURI = "amqp://guest:guest@localhost:5672/"

func main() {
	conn, err := amqp.Dial(brokerURI)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ connection: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connection was successful to RabbitMQ managment console")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	fmt.Println("RabbitMQ connection closed.")
}
