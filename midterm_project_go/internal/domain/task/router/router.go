package router

import (
	"midterm_cloud_project_2024/internal/domain/task/http/handler"
	"net/http"
)

func NewTaskRoutes(mux *http.ServeMux, taskHandler *handler.TaskHandler) {
	mux.HandleFunc("/task", taskHandler.HandleTasks)
}
