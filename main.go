package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

//type TaskRequest struct {
//	Task string `json:"task"`
//}

func getTask(с echo.Context) error {
	return с.String(http.StatusOK, fmt.Sprintf("Hello, %v!", task))
}

func postTask(c echo.Context) error {
	var requestBody struct {
		Task string `json:"task"`
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Неверный JSON")
	}

	task = requestBody.Task
	return c.String(http.StatusOK, "Задача обновлена!")
}

func main() {
	e := echo.New()
	e.GET("/", getTask)
	e.POST("/", postTask)
	e.Logger.Fatal(e.Start(":8080"))
}
