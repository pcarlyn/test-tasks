package dbconector

import (
	"database/sql"
	"errors"
	"fmt"
	"test-tasks/internal/models"
)

func (d *DataBase) GetTaskByTitle(title string) (models.Task, error) {
	query := "SELECT * FROM tasks WHERE title = $1"

	row := d.DB.QueryRow(query, title)

	task := models.Task{}

	var statusID int
	err := row.Scan(&task.ID, &task.Title, &task.Description, &statusID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, fmt.Errorf("task with title %q not found", title)
		}
		return models.Task{}, err
	}
	task.Status = models.CustomStatus{StatusInt: &statusID, Original: statusID}
	return task, nil
}
