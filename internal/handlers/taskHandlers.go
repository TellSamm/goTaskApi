package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskServer/internal/models"
	"taskServer/internal/taskService"
	"taskServer/internal/userService"
	"taskServer/internal/web/tasks"
)

type TaskHandler struct {
	taskService taskService.TaskService
	userService userService.UserService
}

func NewTaskHandler(ts taskService.TaskService, us userService.UserService) *TaskHandler {
	return &TaskHandler{
		taskService: ts,
		userService: us,
	}
}

func (h *TaskHandler) PostTasks(_ context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskReq := req.Body

	userID, err := uuid.Parse(taskReq.UserId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid user_id")
	}

	newTask := models.Task{
		ID:     uuid.New().String(),
		Title:  taskReq.Title,
		Status: taskReq.Status,
		UserID: userID,
	}

	if err := h.taskService.CreateTask(&newTask); err != nil {
		return nil, err
	}

	return tasks.PostTasks201JSONResponse{
		Id:     newTask.ID,
		Title:  newTask.Title,
		Status: newTask.Status,
		UserId: newTask.UserID.String(),
	}, nil
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.taskService.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasks200JSONResponse
	for _, t := range allTasks {
		response = append(response, tasks.Task{
			Id:     t.ID,
			Title:  t.Title,
			Status: t.Status,
			UserId: t.UserID.String(),
		})
	}

	return response, nil
}

func (h *TaskHandler) GetTasksId(_ context.Context, req tasks.GetTasksIdRequestObject) (tasks.GetTasksIdResponseObject, error) {
	task, err := h.taskService.GetTaskByID(req.Id)
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
	task, err := h.taskService.GetTaskByID(req.Id)
	if err != nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	if req.Body.Title != nil {
		task.Title = *req.Body.Title
	}
	if req.Body.Status != nil {
		task.Status = *req.Body.Status
	}

	if err := h.taskService.UpdateTask(task); err != nil {
		return nil, err
	}

	return tasks.PatchTasksId200JSONResponse{
		Id:     task.ID,
		Title:  task.Title,
		Status: task.Status,
	}, nil
}

func (h *TaskHandler) GetTasksByUserId(ctx context.Context, req tasks.GetTasksByUserIdRequestObject) (tasks.GetTasksByUserIdResponseObject, error) {
	userID := req.Id

	taskList, err := h.userService.GetTasksForUser(userID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to get tasks for user")
	}

	var response []tasks.Task
	for _, t := range taskList {
		response = append(response, tasks.Task{
			Id:     t.ID,
			Title:  t.Title,
			Status: t.Status,
			UserId: t.UserID.String(),
		})
	}
	return tasks.GetTasksByUserId200JSONResponse(response), nil
}

func (h *TaskHandler) DeleteTasksId(_ context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.taskService.DeleteTaskByID(req.Id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, nil
	}
	return tasks.DeleteTasksId204Response{}, nil
}
