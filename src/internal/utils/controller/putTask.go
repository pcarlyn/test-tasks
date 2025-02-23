package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"test-tasks/internal/models"
	"test-tasks/internal/utils"
	"test-tasks/internal/utils/database/dbconector"
)

func PutTask(idstr string, task models.RequestTask) (models.Task, int) {

	config := dbconector.NewConfig()
	d, err := dbconector.ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return models.Task{}, http.StatusInternalServerError
	}
	db := dbconector.NewDBConnector(d)
	defer d.Close()

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return models.Task{}, http.StatusBadRequest
	}

	_, statusCode := GetTaskByID(id)
	if statusCode != http.StatusOK {
		return models.Task{}, http.StatusNotFound
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

	responseTask, err := db.PutTask(id, task.Title, *task.Description, statusID)
	if err != nil {
		fmt.Printf("Error post task: %v", err)
		return models.Task{}, http.StatusInternalServerError
	}
	return responseTask, http.StatusOK
}
