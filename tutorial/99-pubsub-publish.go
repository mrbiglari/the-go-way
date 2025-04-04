package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func setupTopic(topicID string, client *pubsub.Client) *pubsub.Topic {
	topic := client.Topic(topicID)

	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		if _, err := client.CreateTopic(ctx, topicID); err != nil {
			log.Fatal("Failed to create topic:", err)
		}
	}
	return topic
}

func _pubsubPublish() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx := context.Background()

	projectID := "project-id"
	topicID := "topic-id"

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	topic := setupTopic(topicID, client)

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("Test Message"),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Published message ID:", id)
}
