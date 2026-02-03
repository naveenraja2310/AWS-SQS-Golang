package sqsproducer

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Producer struct {
	Client   *sqs.Client
	QueueURL string
}

func New(cfg aws.Config, queueURL string) *Producer {
	return &Producer{
		Client:   sqs.NewFromConfig(cfg),
		QueueURL: queueURL,
	}
}

func (p *Producer) Send(ctx context.Context, message string) error {
	out, err := p.Client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &p.QueueURL,
		MessageBody: &message,
	})
	if err != nil {
		return err
	}

	log.Println("message sent successfully, messageId:", *out.MessageId)
	return nil
}
