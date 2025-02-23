package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils/database/dbconector"
)

func GetTaskByTitle(title string) (models.Task, int) {
	config := dbconector.NewConfig()
	d, err := dbconector.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return models.Task{}, http.StatusInternalServerError
	}
	db := dbconector.NewDBConnector(d)
	defer d.Close()

	task, err := db.GetTaskByTitle(title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Task{}, http.StatusNotFound
		}
		fmt.Printf("Error getting tasks from database: %v\n", err)
		return models.Task{}, http.StatusInternalServerError
	}
	return task, http.StatusOK
}
