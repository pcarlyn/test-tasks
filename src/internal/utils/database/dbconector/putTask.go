package dbconector

import "test-tasks/internal/models"

func (d *DataBase) PutTask(id int, title, description string, status_id int) (models.Task, error) {
	query := "UPDATE tasks SET title = $2, description = $3, status_id = $4, updated_at = NOW() WHERE id = $1 RETURNING id, title, description, status_id, created_at, updated_at;"

	row := d.DB.QueryRow(query, id, title, description, status_id)

	var task models.Task
	var statusID int
	err := row.Scan(&task.ID, &task.Title, &task.Description, &statusID, &task.CreatedAt, &task.UpdatedAt)
	task.Status = models.CustomStatus{StatusInt: &statusID, Original: statusID}
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}
