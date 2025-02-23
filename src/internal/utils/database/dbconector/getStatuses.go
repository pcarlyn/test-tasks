package dbconector

import "fmt"

func (d *DataBase) GetStatuses() (map[int]string, error) {
	query := "SELECT * FROM task_statuses;"

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	statuses := make(map[int]string)
	found := false
	for rows.Next() {
		var id int
		var status string
		if err := rows.Scan(&id, &status); err != nil {
			return nil, err
		}
		statuses[id] = status
		found = true
	}

	if !found {
		return nil, fmt.Errorf("no statuses found in task_statuses table")
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return statuses, nil
}
