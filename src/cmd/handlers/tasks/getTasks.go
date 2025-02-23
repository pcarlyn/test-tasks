package handlers

import (
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils/controller"

	"github.com/gofiber/fiber/v2"
)

// GetTasks пример обработчика
// @Description Возвращает все задачи
// @Tags tasks
// @Accept json
// @Produce json
// @Success 200 {object} []models.Task
// @Failure 404 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /tasks [get]
func GetTasks(c *fiber.Ctx) error {
	tasks, statusCode := controller.GetTasks()

	switch statusCode {
	case http.StatusBadRequest:
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{Code: "404", Error: "User not found"})
	case http.StatusInternalServerError:
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}

	return c.Status(statusCode).JSON(tasks)
}
