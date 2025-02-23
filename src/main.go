package main

import (
	"fmt"
	"test-tasks/cmd/routes"
	"test-tasks/internal/utils/database/dbconector"

	_ "test-tasks/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func init() {
	err := dbconector.MigrateDB(dbconector.NewConfig())
	if err != nil {
		fmt.Printf("Error migrating database: %v\n", err)
	}
}

// @title TaskAPI
// @version 1.0
// @description API для управления задачами (TODO-лист)
// @host localhost:8080
// @BasePath /api/v1
func main() {
	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	tasks := app.Group("/api/v1/tasks")

	routes.TasksRoutes(tasks)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen(":8080")
}
