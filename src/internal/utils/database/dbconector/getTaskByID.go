package dbconector

import (
	"database/sql"
	"errors"
	"fmt"
	"test-tasks/internal/models"
)

func (d *DataBase) GetTaskByID(id int) (models.Task, error) {
	query := "SELECT * FROM tasks WHERE id = $1"

	row := d.DB.QueryRow(query, id)

	task := models.Task{}

	var statusID int
	err := row.Scan(&task.ID, &task.Title, &task.Description, &statusID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, fmt.Errorf("task with id %q not found", id)
		}
		return models.Task{}, err
	}
	task.Status = models.CustomStatus{StatusInt: &statusID, Original: statusID}
	return task, nil
}
