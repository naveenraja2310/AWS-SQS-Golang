package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/spf13/viper"
)

type Config struct {
	AwsConfig aws.Config
	QueueURL  string
}

func Load() *Config {
	ctx := context.Background()

	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error reading .env file: %v", err)
	}

	accessKey := v.GetString("AWS_ACCESS_KEY_ID")
	secretKey := v.GetString("AWS_SECRET_ACCESS_KEY")
	region := v.GetString("AWS_REGION")
	queueURL := v.GetString("SQS_QUEUE_URL")

	if accessKey == "" || secretKey == "" || region == "" {
		log.Fatal("AWS credentials or region missing in .env")
	}

	awsCfg, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithRegion(region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKey,
				secretKey,
				"",
			),
		),
	)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	return &Config{
		AwsConfig: awsCfg,
		QueueURL:  queueURL,
	}
}
