package handlers

import (
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils/controller"

	"github.com/gofiber/fiber/v2"
)

// DeleteTask пример обработчика
// @Description Удаляет задачу по id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} []models.Task
// @Failure 404 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/{id} [delete]
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	response, statusCode := controller.DeleteTask(id)
	switch statusCode {
	case http.StatusBadRequest:
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{Code: "404", Error: "Status not found"})
	case http.StatusInternalServerError:
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}

	return c.Status(statusCode).JSON(response)
}
