package main

import (
	"context"
	"log"
	"time"

	"sqs-poc/internal/config"
	"sqs-poc/internal/sqsconsumer"
	"sqs-poc/internal/sqsproducer"
)

func main() {
	log.Println("starting SQS POC using .env credentials...")

	cfg := config.Load()
	ctx := context.Background()

	producer := sqsproducer.New(cfg.AwsConfig, cfg.QueueURL)

	msg := "Hello from Golang SQS Producer 🚀 " + time.Now().Format(time.RFC3339)

	if err := producer.Send(ctx, msg); err != nil {
		log.Fatal("failed to send message:", err)
	}

	log.Println("done")

	consumer := sqsconsumer.New(cfg.AwsConfig, cfg.QueueURL)
	consumer.Start(ctx)
}
