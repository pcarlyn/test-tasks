package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"test-tasks/internal/models"
	"test-tasks/internal/utils/database/dbconector"
)

func DeleteTask(idstr string) (models.Task, int) {

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

	task, statusCode := GetTaskByID(id)
	if statusCode != http.StatusOK {
		return models.Task{}, http.StatusNotFound
	}

	err = db.DeleteTask(id)
	if err != nil {
		fmt.Printf("Error delete task: %v", err)
		return models.Task{}, http.StatusInternalServerError
	}
	return task, http.StatusOK
}
