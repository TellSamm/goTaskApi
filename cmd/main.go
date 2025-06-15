package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"taskServer/internal/db"
	"taskServer/internal/handlers"
	"taskServer/internal/taskService"
	"taskServer/internal/userService"
	"taskServer/internal/web/tasks"
	"taskServer/internal/web/users"
)

func main() {

	database := db.InitDB()
	e := echo.New()

	// Task block
	taskRepo := taskService.NewTaskRepository(database)
	taskSvc := taskService.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskSvc)

	// User block
	userRepo := userService.NewUserRepository(database)
	userSvc := userService.NewUSerService(userRepo)
	userHandler := handlers.NewUserHandler(userSvc)

	// Register handlers
	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	// Run server
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
