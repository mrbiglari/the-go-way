package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

func setupSubscription(topicID string, subID string, client *pubsub.Client) *pubsub.Subscription {
	topic := client.Topic(topicID)
	sub := client.Subscription(subID)

	if exists, err := sub.Exists(ctx); err != nil {
		log.Fatal(err)
	} else if !exists {
		sub, err = client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
			Topic: topic,
		})
		if err != nil {
			log.Fatal("Failed to create subscription:", err)
		}
	}

	return sub
}

func _pubsubReceive() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	ctx := context.Background()

	projectID := "project-id"
	topicID := "topic-id"
	subID := "subscription-id"

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	sub := setupSubscription(topicID, subID, client)

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("Received:", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatal(err)
	}
}
