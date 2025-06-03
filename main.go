package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Task struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Status    string         `json:"status"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

var db *gorm.DB // глобальная переменная для доступа к БД

func getTask(c echo.Context) error {
	var taskList []Task
	if err := db.Find(&taskList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка получения задач из БД"})
	}
	return c.JSON(http.StatusOK, taskList)
}

func postTask(c echo.Context) error {
	var req Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный JSON"})
	}
	req.ID = uuid.New().String()

	if err := db.Create(&req).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось сохранить задачу в базу данных"})
	}
	return c.JSON(http.StatusCreated, req)
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
	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Status != nil {
		task.Status = *req.Status
	}

	if err := db.Save(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось обновить задачу"})
	}

	return c.JSON(http.StatusOK, task)
}

func deleteTask(c echo.Context) error {
	id := c.Param("id")

	var task Task
	if err := db.First(&task, "id = ?", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
	}

	if err := db.Delete(&task).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Ошибка при удалении задачи"})
	}
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	dsn := "host=localhost user=postgres password=postgres dbname=taskdb port=5435 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	db.AutoMigrate(&Task{})

	e.GET("/tasks", getTask)
	e.POST("/tasks", postTask)
	e.PATCH("/tasks/:id", patchTask)
	e.DELETE("/tasks/:id", deleteTask)
	e.Start(":8080")
}
