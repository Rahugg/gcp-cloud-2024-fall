package handlers

import (
	"encoding/json"
	"final_project_cloud_2024/models"
	"net/http"

	"gorm.io/gorm"
)

// GetAllEvents fetches all events from the database
func GetAllEvents(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var events []models.Event
		if err := db.Find(&events).Error; err != nil {
			http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(events)
	}
}

// CreateEvent adds a new event to the database
func CreateEvent(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var event models.Event
		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		if err := db.Create(&event).Error; err != nil {
			http.Error(w, "Failed to create event", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(event)
	}
}
