package routes

import (
	"final_project_cloud_2024/handlers"
	"final_project_cloud_2024/notification"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

// SetupRoutes initializes the router with routes and handlers
func SetupRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Health check
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Golang Event App!"))
	}).Methods("GET")

	// User routes
	r.HandleFunc("/users", handlers.GetAllUsers(db)).Methods("GET")
	r.HandleFunc("/users", handlers.RegisterUser(db)).Methods("POST") // Updated for user registration

	// Notification route
	r.HandleFunc("/send-notification", notification.SendNotification).Methods("POST")

	// Event routes
	r.HandleFunc("/events", handlers.GetAllEvents(db)).Methods("GET")
	r.HandleFunc("/events", handlers.CreateEvent(db)).Methods("POST") // Updated for event creation

	// Ticket purchasing route
	r.HandleFunc("/tickets/purchase", handlers.PurchaseTicket(db)).Methods("POST")

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	return r
}
