package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomStatus struct {
	StatusInt    *int
	StatusString *string
	Original     interface{}
}

func (s *CustomStatus) UnmarshalJSON(data []byte) error {
	var intVal int
	if err := json.Unmarshal(data, &intVal); err == nil {
		*s = CustomStatus{StatusInt: &intVal, Original: intVal}
		return nil
	}

	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		*s = CustomStatus{StatusString: &strVal, Original: strVal}
		return nil
	}

	return fmt.Errorf("invalid status format: %s", string(data))
}

func (s CustomStatus) MarshalJSON() ([]byte, error) {
	if s.StatusInt != nil {
		return json.Marshal(*s.StatusInt)
	}
	if s.StatusString != nil {
		return json.Marshal(*s.StatusString)
	}
	return nil, fmt.Errorf("invalid status fomat")
}

type Task struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description *string      `json:"decription,omitempty"`
	Status      CustomStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type RequestTask struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      string  `json:"status"`
}
