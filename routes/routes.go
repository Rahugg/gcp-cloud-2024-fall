package routes

import (
	"final_project_cloud_2024/handlers"
	"final_project_cloud_2024/notification"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Golang Event App!"))
	}).Methods("GET")

	r.HandleFunc("/users", handlers.GetAllUsers(db)).Methods("GET")
	r.HandleFunc("/users", handlers.RegisterUser(db)).Methods("POST")

	r.HandleFunc("/send-notification", notification.SendNotification).Methods("POST")

	r.HandleFunc("/events", handlers.GetAllEvents(db)).Methods("GET")
	r.HandleFunc("/events", handlers.CreateEvent(db)).Methods("POST")

	r.HandleFunc("/tickets/purchase", handlers.PurchaseTicket(db)).Methods("POST")

	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs"))))

	return r
}
