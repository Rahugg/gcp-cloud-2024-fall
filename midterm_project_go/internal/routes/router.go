package routes

import (
	"encoding/json"
	"midterm_cloud_project_2024/internal/domain/task/http/handler"
	"midterm_cloud_project_2024/internal/domain/task/repository"
	"midterm_cloud_project_2024/internal/domain/task/router"
	"midterm_cloud_project_2024/internal/domain/task/service"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		msg := Message{Message: "Hello to task service"}
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			http.Error(w, "Unable to encode message", http.StatusInternalServerError)

			return
		}
	})

	repo := repository.NewRepository()
	svc := service.NewService(repo)
	taskHandler := handler.NewHandler(svc)

	v1 := http.NewServeMux()
	router.NewTaskRoutes(v1, taskHandler)

	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	return mux
}
