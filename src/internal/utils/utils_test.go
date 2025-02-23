package utils

import (
	"test-tasks/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseBuilder_Success(t *testing.T) {
	statuses := map[int]string{
		1: "new",
		2: "in_progress",
		3: "done",
	}

	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Status: models.CustomStatus{StatusInt: IntPtr(1)}},
		{ID: 2, Title: "Task 2", Status: models.CustomStatus{StatusInt: IntPtr(2)}},
		{ID: 3, Title: "Task 3", Status: models.CustomStatus{StatusInt: IntPtr(3)}},
	}

	expected := []models.Task{
		{ID: 1, Title: "Task 1", Status: models.CustomStatus{StatusString: StrPtr("new"), Original: "new"}},
		{ID: 2, Title: "Task 2", Status: models.CustomStatus{StatusString: StrPtr("in_progress"), Original: "in_progress"}},
		{ID: 3, Title: "Task 3", Status: models.CustomStatus{StatusString: StrPtr("done"), Original: "done"}},
	}

	result, err := ResponseBuilder(statuses, tasks)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestResponseBuilder_UnknownStatus(t *testing.T) {
	statuses := map[int]string{
		1: "new",
		2: "in_progress",
	}

	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Status: models.CustomStatus{StatusInt: IntPtr(1)}},
		{ID: 4, Title: "Task 4", Status: models.CustomStatus{StatusInt: IntPtr(99)}},
	}

	result, err := ResponseBuilder(statuses, tasks)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown status id: 99")
}

func TestResponseBuilder_InvalidStatusString(t *testing.T) {
	statuses := map[int]string{
		1: "new",
		2: "in_progress",
	}

	tasks := []models.Task{
		{ID: 1, Title: "Task 1", Status: models.CustomStatus{StatusInt: IntPtr(1)}},
		{ID: 2, Title: "Task 2", Status: models.CustomStatus{StatusString: StrPtr("already string")}},
	}

	result, err := ResponseBuilder(statuses, tasks)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid task status: expected int, got string")
}

func TestResponseBuilder_EmptyTasks(t *testing.T) {
	statuses := map[int]string{
		1: "new",
		2: "in_progress",
	}

	var tasks []models.Task

	result, err := ResponseBuilder(statuses, tasks)

	assert.NoError(t, err)
	assert.Empty(t, result)
}
