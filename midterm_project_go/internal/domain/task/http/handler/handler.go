package handler

import (
	"encoding/json"
	"midterm_cloud_project_2024/internal/domain/task/entity"
	"midterm_cloud_project_2024/internal/domain/task/service"
	"net/http"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewHandler(svc *service.TaskService) *TaskHandler {
	return &TaskHandler{service: svc}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getTasks(w, r)
	case http.MethodPost:
		h.addTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) getTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		http.Error(w, "Unable to get tasks", http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Unable to encode tasks", http.StatusInternalServerError)

		return
	}
}

func (h *TaskHandler) addTask(w http.ResponseWriter, r *http.Request) {
	var task entity.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid task data", http.StatusBadRequest)

		return
	}
	err := h.service.AddTask(task)

	if err != nil {
		http.Error(w, "Unable to add task", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
