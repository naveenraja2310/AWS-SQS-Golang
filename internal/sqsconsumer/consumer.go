package sqsconsumer

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Consumer struct {
	Client   *sqs.Client
	QueueURL string
}

func New(cfg aws.Config, queueURL string) *Consumer {
	return &Consumer{
		Client:   sqs.NewFromConfig(cfg),
		QueueURL: queueURL,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	log.Println("SQS consumer started...")

	for {
		out, err := c.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &c.QueueURL,
			MaxNumberOfMessages: 1,
			WaitTimeSeconds:     20,
			VisibilityTimeout:   60,
		})

		if err != nil {
			log.Println("receive error:", err)
			continue
		}

		for _, msg := range out.Messages {
			log.Println("message received:", *msg.Body)

			_, err := c.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      &c.QueueURL,
				ReceiptHandle: msg.ReceiptHandle,
			})

			if err != nil {
				log.Println("delete failed:", err)
			} else {
				log.Println("message deleted")
			}
		}
	}
}
