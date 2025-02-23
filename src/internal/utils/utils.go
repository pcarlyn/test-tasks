package utils

import (
	"fmt"
	"test-tasks/internal/models"
)

func GetStatusID(statuses map[int]string, status string) (int, error) {
	for id, name := range statuses {
		if name == status {
			return id, nil
		}
	}
	return 0, fmt.Errorf("status %q not found", status)
}

func ResponseBuilder(statuses map[int]string, tasks []models.Task) ([]models.Task, error) {
	responseTasks := make([]models.Task, 0, len(tasks))

	for _, task := range tasks {
		if task.Status.StatusString != nil && task.Status.StatusInt == nil {
			return nil, fmt.Errorf("invalid task status: expected int, got string")
		}

		if task.Status.StatusInt != nil {
			statusStr, ok := statuses[*task.Status.StatusInt]
			if !ok {
				return nil, fmt.Errorf("unknown status id: %d", *task.Status.StatusInt)
			}
			task.Status = models.CustomStatus{StatusString: &statusStr, Original: statusStr}
		}

		responseTasks = append(responseTasks, task)
	}

	return responseTasks, nil
}

func IntPtr(i int) *int {
	return &i
}

func StrPtr(s string) *string {
	return &s
}
