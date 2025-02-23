package dbconector

import (
	"test-tasks/internal/models"
)

func (d *DataBase) PostTask(title, description string, status_id int) (models.Task, error) {

	query := "INSERT INTO tasks (title, description, status_id) VALUES ($1, $2, $3) RETURNING id, title, description, status_id, created_at, updated_at"

	row := d.DB.QueryRow(query, title, description, status_id)

	var task models.Task
	var statusID int
	err := row.Scan(&task.ID, &task.Title, &task.Description, &statusID, &task.CreatedAt, &task.UpdatedAt)
	task.Status = models.CustomStatus{StatusInt: &statusID, Original: statusID}
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
