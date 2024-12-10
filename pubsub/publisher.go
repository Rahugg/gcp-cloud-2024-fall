package pubsub

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type NotificationMessage struct {
	UserID  uint   `json:"user_id"`
	EventID uint   `json:"event_id"`
	Message string `json:"message"`
}

// PublishNotification publishes a message to the Pub/Sub topic
func PublishNotification(ctx context.Context, projectID, topicID string, message NotificationMessage) error {
	client, err := pubsub.NewClient(ctx, projectID, option.WithCredentialsFile("path/to/service-account.json"))
	if err != nil {
		return err
	}
	defer client.Close()

	topic := client.Topic(topicID)
	defer topic.Stop()

	msgData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: msgData,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return err
	}

	log.Printf("Message published with ID: %s", id)
	return nil
}
