package handlers

import (
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils/controller"

	"github.com/gofiber/fiber/v2"
)

// PostTask пример обработчика
// @Description Добавляет задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body models.RequestTask true "JSON задачи"
// @Success 200 {object} []models.Task
// @Failure 409 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/tasks/ [post]
func PostTask(c *fiber.Ctx) error {

	var request models.RequestTask

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
	response, statusCode := controller.PostTask(request)
	switch statusCode {
	case http.StatusConflict:
		return c.Status(http.StatusConflict).JSON(models.ErrorResponse{Code: "409", Error: "Conflict"})
	case http.StatusBadRequest:
		return c.Status(http.StatusBadRequest).JSON(models.ErrorResponse{Code: "400", Error: "Bad Request"})
	case http.StatusNotFound:
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{Code: "404", Error: "Status not found"})
	case http.StatusInternalServerError:
		return c.Status(http.StatusInternalServerError).JSON(models.ErrorResponse{Code: "500", Error: "Internal Server Error"})
	}

	return c.Status(statusCode).JSON(response)
}
