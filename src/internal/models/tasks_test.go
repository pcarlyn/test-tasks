package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestCustomStatusUnmarshalJSON(t *testing.T) {
	tests := []struct {
		input       string
		expectedInt *int
		expectedStr *string
		shouldFail  bool
	}{
		{`42`, intPtr(42), nil, false},
		{`"in_progress"`, nil, strPtr("in_progress"), false},
		{`true`, nil, nil, true}, // Ошибочный формат
	}

	for _, test := range tests {
		var status CustomStatus
		err := json.Unmarshal([]byte(test.input), &status)

		if test.shouldFail {
			if err == nil {
				t.Errorf("Expected error for input %s, but got none", test.input)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input %s: %v", test.input, err)
			continue
		}

		if test.expectedInt != nil && (status.StatusInt == nil || *status.StatusInt != *test.expectedInt) {
			t.Errorf("Expected StatusInt %d, got %v", *test.expectedInt, status.StatusInt)
		}

		if test.expectedStr != nil && (status.StatusString == nil || *status.StatusString != *test.expectedStr) {
			t.Errorf("Expected StatusString %s, got %v", *test.expectedStr, status.StatusString)
		}
	}
}

func TestCustomStatusMarshalJSON(t *testing.T) {
	tests := []struct {
		status     CustomStatus
		expected   string
		shouldFail bool
	}{
		{CustomStatus{StatusInt: intPtr(42)}, `42`, false},
		{CustomStatus{StatusString: strPtr("done")}, `"done"`, false},
		{CustomStatus{}, "", true}, // Ошибочный случай
	}

	for _, test := range tests {
		data, err := json.Marshal(test.status)

		if test.shouldFail {
			if err == nil {
				t.Errorf("Expected error for status %+v, but got none", test.status)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for status %+v: %v", test.status, err)
			continue
		}

		if string(data) != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, string(data))
		}
	}
}

func TestTasksUnmarshalJSON(t *testing.T) {
	jsonData := `{
		"id": 1,
		"title": "Fix bug",
		"description": "Fix issue in the payment system",
		"status": "in_progress",
		"created_at": "2024-02-20T12:00:00Z",
		"updated_at": "2024-02-20T13:00:00Z"
	}`

	var task Task
	if err := json.Unmarshal([]byte(jsonData), &task); err != nil {
		t.Fatalf("Error decoding JSON: %v", err)
	}

	expectedStatus := "in_progress"
	if task.Status.StatusString == nil || *task.Status.StatusString != expectedStatus {
		t.Errorf("Expected Status to be '%s', got %+v", expectedStatus, task.Status)
	}
}

func TestTasksMarshalJSON(t *testing.T) {
	createdAt := time.Date(2024, 2, 20, 12, 0, 0, 0, time.UTC)
	updatedAt := time.Date(2024, 2, 20, 13, 0, 0, 0, time.UTC)
	status := "in_progress"

	task := Task{
		ID:          1,
		Title:       "Fix bug",
		Description: nil,
		Status:      CustomStatus{StatusString: &status, Original: status},
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	expectedJSON := `{"id":1,"title":"Fix bug","status":"in_progress","created_at":"2024-02-20T12:00:00Z","updated_at":"2024-02-20T13:00:00Z"}`

	result, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(string(result), expectedJSON) {
		t.Errorf("Expected %s, got %s", expectedJSON, string(result))
	}
}

// Вспомогательные функции для указателей
func intPtr(i int) *int       { return &i }
func strPtr(s string) *string { return &s }
