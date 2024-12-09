package routes

import (
	"final_project_cloud_2024/handlers"
	"github.com/gorilla/mux"

	"gorm.io/gorm"
	"net/http"
)

// SetupRoutes initializes the router with routes and handlers
func SetupRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Health check route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Golang Event App!"))
	}).Methods("GET")

	// User routes
	r.HandleFunc("/users", handlers.GetAllUsers(db)).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")

	// Event routes
	r.HandleFunc("/events", handlers.GetAllEvents(db)).Methods("GET")
	r.HandleFunc("/events", handlers.CreateEvent(db)).Methods("POST")

	return r
}
