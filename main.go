package main

import (
	"final_project_cloud_2024/db"
	"final_project_cloud_2024/routes"
	"log"
	"net/http"
)

func main() {
	// Connect to the database and run migrations
	dbConn := db.ConnectAndMigrate()

	// Setup routes
	r := routes.SetupRoutes(dbConn)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
