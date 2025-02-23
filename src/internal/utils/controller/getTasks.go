package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils"
	"test-tasks/internal/utils/database/dbconector"
)

func GetTasks() ([]models.Task, int) {
	config := dbconector.NewConfig()
	d, err := dbconector.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return nil, http.StatusInternalServerError
	}
	db := dbconector.NewDBConnector(d)
	defer d.Close()

	tasks, err := db.GetTasks()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound
		}
		fmt.Printf("Error getting tasks from database: %v\n", err)
		return nil, http.StatusInternalServerError
	}

	statuses, err := db.GetStatuses()
	if err != nil {
		fmt.Printf("Error getting statuses from database: %v\n", err)
		return nil, http.StatusInternalServerError
	}

	responseTasks, err := utils.ResponseBuilder(statuses, tasks)
	if err != nil {
		fmt.Printf("Error build response: %v\n", err)
		return nil, http.StatusInternalServerError
	}

	return responseTasks, http.StatusOK
}
