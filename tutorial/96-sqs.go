package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type TestMessage struct {
	Message string `json:"Message"`
}

func extractMessage(msg types.Message) string {
	var m TestMessage
	err := json.Unmarshal([]byte(*msg.Body), &m)
	if err != nil {
		fmt.Print(err)
	}
	message := m.Message
	return message
}

func printMessage(msg types.Message) {
	fmt.Println(extractMessage(msg))
}

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

var cfg = getConfig()
var client = sqs.NewFromConfig(cfg)
var ctx = context.Background()

const queueURL = "http://localhost:4566/000000000000/queue1"

func _sqs() {
	readConcurrent() // change to read(), readConcurrent() or readPolled()
}

func read() {
	for {
		output, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL), // url of the sqs queue to receive messages from
			MaxNumberOfMessages: 5,                    // maximum number of messages to retrieve (up to 10)
			VisibilityTimeout:   30,                   // time in seconds the message stays invisible to others after being received
			WaitTimeSeconds:     20,                   // enable long polling to wait up to 20 seconds for a message
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, msg := range output.Messages {
			printMessage(msg) // ...message processing logic
			client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: msg.ReceiptHandle,
			})
		}
	}
}

func readConcurrent() {
	for {
		output, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL), // url of the sqs queue to receive messages from
			MaxNumberOfMessages: 5,                    // maximum number of messages to retrieve (up to 10)
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
				printMessage(msg) // ...message processing logic
				client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: m.ReceiptHandle,
				})
			}(msg)
		}
		wg.Wait()
	}
}

func readPolled() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		output, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL), // url of the sqs queue to receive messages from
			MaxNumberOfMessages: 5,                    // maximum number of messages to retrieve (up to 10)
			VisibilityTimeout:   20,                   // time in seconds the message stays invisible to others after being received
			WaitTimeSeconds:     20,                   // enable long polling to wait up to 20 seconds for a message
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, msg := range output.Messages {
			printMessage(msg) // ...message processing logic
			client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: msg.ReceiptHandle,
			})
		}
	}
}
