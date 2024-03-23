package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Define Kafka broker address
	kafkaBroker := "localhost:9093" // Update this with your Kafka broker address

	// Create a new Kafka writer
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    "your-topic2",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	// Trap SIGINT and SIGTERM to trigger a graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Produce messages
	go func() {
		for {
			select {
			case <-signals:
				return
			default:
				message := kafka.Message{
					Key:   nil,
					Value: []byte("your-message"),
				}
				err := writer.WriteMessages(context.Background(), message)
				if err != nil {
					log.Printf("Error sending message: %s\n", err)
				} else {
					log.Printf("Message sent successfully\n")
				}
			}
		}
	}()

	// Keep the producer running until interrupted
	<-signals
	log.Println("Shutting down")
}
