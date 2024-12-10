package handlers

import (
	"context"
	"encoding/json"
	"final_project_cloud_2024/pubsub"
	"log"
	"net/http"
	"time"

	"final_project_cloud_2024/models"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registration models.Registration
		if err := json.NewDecoder(r.Body).Decode(&registration); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		registration.RegistrationDate = time.Now()

		if err := db.Create(&registration).Error; err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}

		ctx := context.Background()
		message := pubsub.NotificationMessage{
			UserID:  registration.UserID,
			EventID: registration.EventID,
			Message: "You have successfully registered for the event!",
		}

		if err := pubsub.PublishNotification(ctx, "your-gcp-project-id", "notifications-topic", message); err != nil {
			log.Printf("Failed to publish notification: %v", err)
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(registration)
	}
}

func CreateEvent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		event.CreatedAt = time.Now()
		event.UpdatedAt = time.Now()

		if err := db.Create(&event).Error; err != nil {
			http.Error(w, "Failed to create event", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(event)
	}
}

func PurchaseTicket(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registration models.Registration
		if err := json.NewDecoder(r.Body).Decode(&registration); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		registration.RegistrationDate = time.Now()

		var event models.Event
		if err := db.First(&event, registration.EventID).Error; err != nil {
			http.Error(w, "Event not found", http.StatusNotFound)
			return
		}

		var user models.User
		if err := db.First(&user, registration.UserID).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if err := db.Create(&registration).Error; err != nil {
			http.Error(w, "Failed to register for event", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(registration)
	}
}
