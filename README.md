# AWS-SQS-Golang

# Golang SQS POC (.env credentials)

## Setup
1. Create IAM user with SQS permissions
2. Add credentials to `.env`
3. Do NOT commit `.env`

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
│       └── producer.go   👈 NEW
├── .env
├── .gitignore
├── go.mod
└── README.md
```
