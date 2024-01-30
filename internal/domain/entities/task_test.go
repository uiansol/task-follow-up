package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	t.Run("should create a task and return it with error nil", func(t *testing.T) {
		task, err := NewTask("techinicianID", "managerID", "test description")

		assert.Nil(t, err)
		assert.NotNil(t, task.ID)
		assert.Equal(t, "techinicianID", task.TechnicianID)
		assert.Equal(t, "managerID", task.ManagerID)
		assert.Equal(t, "test description", task.Description)
		assert.Equal(t, Open, task.Status)
		assert.NotEqual(t, time.Time{}, task.CreatedAt)
		assert.NotEqual(t, time.Time{}, task.ModifiedAt)
		assert.Equal(t, time.Time{}, task.ClosedAt)
	})

	t.Run("should return error if invalid technicianID", func(t *testing.T) {
		_, err := NewTask("", "managerID", "test description")

		assert.NotNil(t, err)
		assert.Equal(t, "technician ID is required", err.Error())
	})

	t.Run("should return error if invalid managerID", func(t *testing.T) {
		_, err := NewTask("technicianID", "", "test description")

		assert.NotNil(t, err)
		assert.Equal(t, "manager ID is required", err.Error())
	})

	t.Run("should return error if invalid description", func(t *testing.T) {
		_, err := NewTask("technicianID", "managerID", "")

		assert.NotNil(t, err)
		assert.Equal(t, "description is required", err.Error())
	})
}
