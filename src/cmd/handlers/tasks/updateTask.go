package handlers

import (
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils/controller"

	"github.com/gofiber/fiber/v2"
)

// UpdateTask пример обработчика
// @Description Изменяет задачу по id
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param request body models.RequestTask true "JSON задачи"
// @Success 200 {object} []models.Task
// @Failure 404 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/{id} [put]
func UpdateTask(c *fiber.Ctx) error {
	var request models.RequestTask

	id := c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if request.Title == "" || request.Description == nil || *request.Description == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Title and Description are required",
		})
	}

	response, statusCode := controller.PutTask(id, request)
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
