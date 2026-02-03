# AWS-SQS-Golang

# Golang SQS POC (.env credentials)

## Setup create a .env file in root path
```
AWS_ACCESS_KEY_ID=*************
AWS_SECRET_ACCESS_KEY=*********
AWS_REGION=<your-aws-region>
SQS_QUEUE_URL=<your-sqs-queue-url>
```

## Run
```bash
go run cmd/app/main.go
```

## Folder Structure
```sqs-poc/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── sqsconsumer/
│   │   └── consumer.go
│   └── sqsproducer/
│       └── producer.go
├── .env
├── .gitignore
├── go.mod
└── README.md
```
## AWS IAM User inline policy
```
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": [
				"sqs:SendMessage",
				"sqs:ReceiveMessage",
				"sqs:DeleteMessage",
				"sqs:GetQueueAttributes"
			],
			"Resource": "your-sqs-arn" //arn:aws:sqs:<aws-region>:123456789012:<sqs-name>
		}
	]
}
```
