package dbconector

import (
	"fmt"
	"test-tasks/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	config := NewConfig()
	d, err := ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	db := NewDBConnector(d)
	assert.NoError(t, err)
	defer db.DB.Close()

	tasks, err := db.GetTasks()
	assert.NoError(t, err)

	// Приводим tasks к срезу анонимных структур без дат
	actual := make([]struct {
		ID          int
		Title       string
		Description *string
		Status      int
	}, len(tasks))

	for i, task := range tasks {
		// Достаём статус из CustomStatus (он у тебя в базе как int)
		var statusValue int
		if task.Status.StatusInt != nil {
			statusValue = *task.Status.StatusInt
		} else {
			t.Fatalf("unexpected status format: %+v", task.Status)
		}

		actual[i] = struct {
			ID          int
			Title       string
			Description *string
			Status      int
		}{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      statusValue,
		}
	}

	// Ожидаемые данные без дат
	expected := []struct {
		ID          int
		Title       string
		Description *string
		Status      int
	}{
		{1, "Разработка API", utils.StrPtr("Создать эндпоинты для управления задачами"), 1},
		{2, "Написание документации", utils.StrPtr("Описать все методы API"), 2},
		{3, "Тестирование функционала", utils.StrPtr("Проверить работу всех эндпоинтов"), 3},
		{4, "Оптимизация базы данных", utils.StrPtr("Настроить индексы и оптимизировать запросы"), 1},
	}

	assert.Equal(t, expected, actual)
}
