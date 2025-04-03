package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func getConfig() aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("key", "secret", ""),
		),
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint("http://localhost:4566"),
	)
	return cfg
}

func _sqs() {
	cfg := getConfig()
	client := sqs.NewFromConfig(cfg)
	const queueURL = "http://localhost:4566/000000000000/liveMetricsV3"

	for {
		output, err := client.ReceiveMessage(context.Background(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL), // url of the sqs queue to receive messages from
			MaxNumberOfMessages: 1,                    // maximum number of messages to retrieve (up to 10)
			VisibilityTimeout:   30,                   // time in seconds the message stays invisible to others after being received
			WaitTimeSeconds:     20,                   // enable long polling to wait up to 20 seconds for a message
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		var wg sync.WaitGroup
		for _, msg := range output.Messages {
			wg.Add(1)
			go func(m types.Message) {
				defer wg.Done()
				fmt.Println(*m.Body) // ...message processing logic
				client.DeleteMessage(context.Background(), &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: m.ReceiptHandle,
				})
			}(msg)
		}
		wg.Wait()
	}
}
