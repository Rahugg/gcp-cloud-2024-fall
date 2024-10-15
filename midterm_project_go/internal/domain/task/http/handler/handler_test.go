package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"midterm_cloud_project_2024/internal/domain/task/entity"
	"midterm_cloud_project_2024/internal/domain/task/http/handler"
	_ "midterm_cloud_project_2024/internal/domain/task/service"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) GetTasks() ([]entity.Task, error) {
	args := m.Called()
	return args.Get(0).([]entity.Task), args.Error(1)
}

func (m *MockTaskService) AddTask(task entity.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func TestTaskHandler_HandleTasks_Get(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("GetTasks").Return([]entity.Task{{ID: strconv.Itoa(1), Description: "Test Task"}}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()

	taskHandler := handler.NewHandler(mockService)
	taskHandler.HandleTasks(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestTaskHandler_HandleTasks_Post(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("AddTask", mock.Anything).Return(nil)

	task := entity.Task{Description: "Test Task"}
	taskJson, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJson))
	rr := httptest.NewRecorder()

	taskHandler := handler.NewHandler(mockService)
	taskHandler.HandleTasks(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockService.AssertExpectations(t)
}

func TestTaskHandler_HandleTasks_Get_Error(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("GetTasks").Return(nil, errors.New("error"))

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	rr := httptest.NewRecorder()

	taskHandler := handler.NewHandler(mockService)
	taskHandler.HandleTasks(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}

func TestTaskHandler_HandleTasks_Post_Error(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("AddTask", mock.Anything).Return(errors.New("error"))

	task := entity.Task{Description: "Test Task"}
	taskJson, _ := json.Marshal(task)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(taskJson))
	rr := httptest.NewRecorder()

	taskHandler := handler.NewHandler(mockService)
	taskHandler.HandleTasks(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockService.AssertExpectations(t)
}
