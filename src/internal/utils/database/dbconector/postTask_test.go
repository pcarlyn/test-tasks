package dbconector

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostTask(t *testing.T) {
	config := NewConfig()
	d, err := ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	db := NewDBConnector(d)
	assert.NoError(t, err)
	defer db.DB.Close()

	// Входные данные
	title := "Тестовая задача"
	description := "Описание тестовой задачи"
	status := 1

	// Вызываем PostTask
	task, err := db.PostTask(title, description, status)
	assert.NoError(t, err)

	// Проверяем, что статус — это int
	var statusValue int
	if task.Status.StatusInt != nil {
		statusValue = *task.Status.StatusInt
	} else {
		t.Fatalf("unexpected status format: %+v", task.Status)
	}

	// Ожидаемая структура
	expected := struct {
		Title       string
		Description string
		Status      int
	}{
		Title:       title,
		Description: description,
		Status:      status,
	}

	// Проверяем, что данные совпадают
	assert.Equal(t, expected, struct {
		Title       string
		Description string
		Status      int
	}{
		Title:       task.Title,
		Description: *task.Description,
		Status:      statusValue,
	})

	// Проверяем, что ID, created_at и updated_at заполнены
	assert.NotZero(t, task.ID, "ID задачи не должен быть нулевым")
	assert.NotZero(t, task.CreatedAt, "created_at не должен быть пустым")
	assert.NotZero(t, task.UpdatedAt, "updated_at не должен быть пустым")
}
