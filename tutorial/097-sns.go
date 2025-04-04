package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

const topicArn = "arn:aws:sns:us-east-1:000000000000:topic1"

func _sns() {

	cfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("key", "secret", ""),
		),
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint("http://localhost:4566"),
	)
	client := sns.NewFromConfig(cfg)
	ctx := context.Background()

	out, err := client.Publish(ctx, &sns.PublishInput{
		Message:  aws.String("Test Message at " + time.Now().String()),
		TopicArn: aws.String(topicArn),
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*out.MessageId)
}
