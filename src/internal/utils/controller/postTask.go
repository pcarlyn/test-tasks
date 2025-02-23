package controller

import (
	"fmt"
	"net/http"
	"test-tasks/internal/models"
	"test-tasks/internal/utils"
	"test-tasks/internal/utils/database/dbconector"
)

func PostTask(task models.RequestTask) (models.Task, int) {

	config := dbconector.NewConfig()
	d, err := dbconector.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return models.Task{}, http.StatusInternalServerError
	}
	db := dbconector.NewDBConnector(d)
	defer d.Close()

	_, statusCode := GetTaskByTitle(task.Title)
	if statusCode == http.StatusOK {
		return models.Task{}, http.StatusConflict
	}
	statuses, err := db.GetStatuses()
	if err != nil {
		fmt.Printf("Error getting statuses: %v\n", err)
		if err.Error() == "no statuses found in task_statuses table" {
			return models.Task{}, http.StatusNotFound
		}
		return models.Task{}, http.StatusInternalServerError
	}
	statusID, err := utils.GetStatusID(statuses, task.Status)
	if err != nil {
		fmt.Printf("Error getting status ID: %v\n", err)
		return models.Task{}, http.StatusNotFound
	}

	responseTask, err := db.PostTask(task.Title, *task.Description, statusID)
	if err != nil {
		fmt.Printf("Error post task: %v", err)
		return models.Task{}, http.StatusInternalServerError
	}
	return responseTask, http.StatusCreated
}
