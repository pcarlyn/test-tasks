package dbconector

import "fmt"

func (d *DataBase) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	result, err := d.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}

	return nil
}
