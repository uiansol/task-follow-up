package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	taskrepository "github.com/uiansol/task-follow-up/internal/task_repository"
)

func TestNewTaskCreateUseCase(t *testing.T) {
	taskRepositoryMock := taskrepository.NewITaskRepository(t)

	t.Run("should return a task create use case", func(t *testing.T) {
		taskCreateUseCase := NewTaskCreateUseCase(taskRepositoryMock)

		assert.NotNil(t, taskCreateUseCase)
		assert.Equal(t, taskRepositoryMock, taskCreateUseCase.taskRepository)
	})
}

func TestExecute(t *testing.T) {
	input := TaskCreateInput{
		TechinicianID: "123",
		ManagerID:     "456",
		Description:   "test description",
	}

	t.Run("should create the task and return error nil", func(t *testing.T) {
		taskRepositoryMock := taskrepository.NewITaskRepository(t)
		taskCreateUseCase := NewTaskCreateUseCase(taskRepositoryMock)

		taskRepositoryMock.On("Save", mock.Anything).Return("abc", nil)

		output, err := taskCreateUseCase.Execute(input)

		assert.Nil(t, err)
		assert.Equal(t, "abc", output.TaskID)
	})

	t.Run("should return error when TechinicianID is empty", func(t *testing.T) {
		taskRepositoryMock := taskrepository.NewITaskRepository(t)
		taskCreateUseCase := NewTaskCreateUseCase(taskRepositoryMock)

		inputTechinicianIDEmpty := TaskCreateInput{
			TechinicianID: "",
			ManagerID:     "456",
			Description:   "test description",
		}

		output, err := taskCreateUseCase.Execute(inputTechinicianIDEmpty)

		assert.Equal(t, "technician ID is required", err.Error())
		assert.Equal(t, TaskCreateOutput{}, output)
	})

	t.Run("should return error when couldn't save", func(t *testing.T) {
		taskRepositoryMock := taskrepository.NewITaskRepository(t)
		taskCreateUseCase := NewTaskCreateUseCase(taskRepositoryMock)

		taskRepositoryMock.On("Save", mock.Anything).Return("", errors.New("test error"))

		output, err := taskCreateUseCase.Execute(input)

		assert.Equal(t, "test error", err.Error())
		assert.Equal(t, TaskCreateOutput{}, output)
	})
}
