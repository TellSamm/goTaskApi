package main

import (
	"github.com/labstack/echo/v4"
	"taskServer/internal/db"
	"taskServer/internal/handlers"
	"taskServer/internal/taskService"
	"taskServer/internal/web/tasks"
)

func main() {

	database := db.InitDB()

	repo := taskService.NewTaskRepository(database)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	e := echo.New()

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	e.Start(":8080")
}
