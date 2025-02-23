package dbconector

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStatuses(t *testing.T) {
	config := NewConfig()
	d, err := ConnectToDB(config)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	db := NewDBConnector(d)
	assert.NoError(t, err)
	defer db.DB.Close()

	statuses, err := db.GetStatuses()
	assert.NoError(t, err)

	expected := map[int]string{
		1: "new",
		2: "in_progress",
		3: "done",
	}
	assert.Equal(t, expected, statuses)
}
