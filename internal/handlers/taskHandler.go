package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskServer/internal/models"
	"taskServer/internal/taskService"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var task models.Task

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неправильный JSON"})
	}

	task.ID = uuid.New().String()

	if err := h.service.CreateTask(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось создать задачу"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось получить список задач"})
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	id := c.Param("id")

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")

	var req struct {
		Title  *string `json:"title"`
		Status *string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "неверный JSON"})
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Status != nil {
		task.Status = *req.Status
	}
	if err := h.service.UpdateTask(task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "не удалось обновить задачу"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")

	err := h.service.DeleteTaskByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "задача не найдена"})
	}
	return c.NoContent(http.StatusNoContent)
}
