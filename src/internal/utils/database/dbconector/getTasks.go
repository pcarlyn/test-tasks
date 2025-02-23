package dbconector

import "test-tasks/internal/models"

func (d *DataBase) GetTasks() ([]models.Task, error) {
	query := "SELECT * FROM tasks;"

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]models.Task, 0)
	for rows.Next() {
		task := models.Task{}
		var statusInt int
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &statusInt, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		task.Status = models.CustomStatus{StatusInt: &statusInt, Original: statusInt}

		tasks = append(tasks, task)
	}
	return tasks, nil
}
