package main

import (
	"final_project_cloud_2024/db"
	"final_project_cloud_2024/routes"
	"log"
	"net/http"

	_ "github.com/swaggo/files" // Required for serving Swagger UI
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Connect to the database and run migrations
	dbConn := db.ConnectAndMigrate()

	// Setup routes
	r := routes.SetupRoutes(dbConn)

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/docs/openapi.yaml"), // Path to your OpenAPI spec
	))

	// Serve the OpenAPI spec file
	fs := http.FileServer(http.Dir("./docs"))
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	log.Println("Server starting on http://localhost:8100")
	log.Fatal(http.ListenAndServe(":8100", r))
}
