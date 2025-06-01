package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var tasks = []Task{}

func getTask(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный JSON"})
	}
	newTask := Task{
		ID:     uuid.NewString(),
		Title:  req.Title,
		Status: req.Status,
	}
	tasks = append(tasks, newTask)
	return c.JSON(http.StatusCreated, newTask)
}

func patchTask(c echo.Context) error {
	id := c.Param("id")
	var req struct {
		Title  *string `json:"title"`
		Status *string `json:"status"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный JSON"})
	}
	for i, t := range tasks {
		if t.ID == id {
			if req.Title != nil {
				tasks[i].Title = *req.Title
			}
			if req.Status != nil {
				tasks[i].Status = *req.Status
			}
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.String(http.StatusOK, "Задача удалена!")
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
}

func main() {
	e := echo.New()
	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.PATCH("/task/:id", patchTask)
	e.DELETE("/task/:id", deleteTask)
	e.Start(":8080")
}
