package routes

import (
	handlers "test-tasks/cmd/handlers/tasks"

	"github.com/gofiber/fiber/v2"
)

func TasksRoutes(group fiber.Router) {
	group.Get("", handlers.GetTasks)
	group.Post("", handlers.PostTask)
	group.Put("/:id", handlers.UpdateTask)
	group.Delete("/:id", handlers.DeleteTask)
}
