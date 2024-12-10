package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

type Notification struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
	Message string `json:"message"`
}

func SendNotification(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var notif Notification
	if err := json.NewDecoder(r.Body).Decode(&notif); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	if notif.UserID == "" || notif.EventID == "" || notif.Message == "" {
		http.Error(w, "Missing required fields: user_id, event_id, message", http.StatusBadRequest)
		return
	}

	topicID := os.Getenv("PUBSUB_TOPIC")
	if topicID == "" {
		http.Error(w, "PUBSUB_TOPIC environment variable not set", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		log.Printf("Failed to create Pub/Sub client: %v", err)
		http.Error(w, "Failed to create Pub/Sub client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	topic := client.Topic(topicID)
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(fmt.Sprintf("User %s for Event %s: %s", notif.UserID, notif.EventID, notif.Message)),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		http.Error(w, "Failed to send notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"status":       "Notification sent",
		"message_id":   id,
		"user_id":      notif.UserID,
		"event_id":     notif.EventID,
		"notification": notif.Message,
	}
	json.NewEncoder(w).Encode(response)
}
