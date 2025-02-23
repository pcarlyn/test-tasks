package dbconector

import (
	"fmt"
	"test-tasks/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTasksByTitle(t *testing.T) {
	config := NewConfig()
	d, err := ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	db := NewDBConnector(d)
	assert.NoError(t, err)
	defer db.DB.Close()

	title := "Разработка API"
	task, err := db.GetTaskByTitle(title)
	assert.NoError(t, err)

	var statusValue int
	if task.Status.StatusInt != nil {
		statusValue = *task.Status.StatusInt
	} else {
		t.Fatalf("unexpected status format: %+v", task.Status)
	}

	expected := struct {
		ID          int
		Title       string
		Description *string
		Status      int
	}{
		ID:          1,
		Title:       "Разработка API",
		Description: utils.StrPtr("Создать эндпоинты для управления задачами"),
		Status:      1,
	}

	assert.Equal(t, expected, struct {
		ID          int
		Title       string
		Description *string
		Status      int
	}{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      statusValue,
	})
}
