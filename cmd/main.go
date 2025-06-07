package main

import (
	"github.com/labstack/echo/v4"
	"taskServer/internal/db"
	"taskServer/internal/handlers"
	"taskServer/internal/taskService"
)

func main() {

	database := db.InitDB()

	repo := taskService.NewTaskRepository(database)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	e := echo.New()

	e.GET("/tasks", handler.GetAllTasks)
	e.POST("/tasks", handler.CreateTask)
	e.GET("/tasks/:id", handler.GetTaskByID)
	e.PATCH("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)
	e.Start(":8080")
}
