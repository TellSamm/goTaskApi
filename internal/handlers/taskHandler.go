package handlers

import (
	"context"
	"github.com/google/uuid"
	"taskServer/internal/models"
	"taskServer/internal/taskService"
	"taskServer/internal/web/tasks"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewTaskHandler(service taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) PostTasks(_ context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskReq := req.Body

	newTask := models.Task{
		ID:     uuid.New().String(),
		Title:  taskReq.Title,
		Status: taskReq.Status,
	}

	if err := h.service.CreateTask(&newTask); err != nil {
		return nil, err
	}

	return tasks.PostTasks201JSONResponse{
		Id:     newTask.ID,
		Title:  newTask.Title,
		Status: newTask.Status,
	}, nil
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasks200JSONResponse
	for _, t := range allTasks {
		response = append(response, tasks.Task{
			Id:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}

	return response, nil
}

func (h *TaskHandler) GetTasksId(_ context.Context, req tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := h.service.GetTaskByID(req.Id)
	if err != nil {
		return tasks.GetTasksId404Response{}, nil
	}

	return tasks.GetTasksId200JSONResponse{
		Id:     task.ID,
		Title:  task.Title,
		Status: task.Status,
	}, nil
}

func (h *TaskHandler) PatchTasksId(_ context.Context, req tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	task, err := h.service.GetTaskByID(req.Id)
	if err != nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	if req.Body.Title != nil {
		task.Title = *req.Body.Title
	}
	if req.Body.Status != nil {
		task.Status = *req.Body.Status
	}

	if err := h.service.UpdateTask(task); err != nil {
		return nil, err
	}

	return tasks.PatchTasksId200JSONResponse{
		Id:     task.ID,
		Title:  task.Title,
		Status: task.Status,
	}, nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.service.DeleteTaskByID(req.Id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, nil
	}
	return tasks.DeleteTasksId204Response{}, nil
}
