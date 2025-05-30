package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

var task string

func getTask(с echo.Context) error {
	return с.String(http.StatusOK, fmt.Sprintf("Hello, %v!", task))
}

func postTask(c echo.Context) error {
	var requestBody struct {
		Task string `json:"task"`
	}
	decoder := json.NewDecoder(c.Request().Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Неверный JSON: "+err.Error())
	}

	task = requestBody.Task
	return c.String(http.StatusOK, "Задача обновлена!")
}

func main() {
	e := echo.New()
	e.GET("/task", getTask)
	e.POST("/task", postTask)
	e.Start(":8080")
}
