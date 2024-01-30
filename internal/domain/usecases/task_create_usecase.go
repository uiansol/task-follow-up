package usecases

import (
	"github.com/uiansol/task-follow-up/internal/domain/adapters"
	"github.com/uiansol/task-follow-up/internal/domain/entities"
)

type ITaskCreateUseCase interface {
	Execute(taskCreateInput TaskCreateInput) (TaskCreateOutput, error)
}

type TaskCreateUseCase struct {
	taskRepository adapters.ITaskRepository
}

type TaskCreateInput struct {
	TechinicianID string
	ManagerID     string
	Description   string
}

type TaskCreateOutput struct {
	TaskID string
}

func NewTaskCreateUseCase(taskRepository adapters.ITaskRepository) TaskCreateUseCase {
	return TaskCreateUseCase{
		taskRepository: taskRepository,
	}
}

func (u *TaskCreateUseCase) Execute(userCreateInput TaskCreateInput) (TaskCreateOutput, error) {
	task, err := entities.NewTask(userCreateInput.TechinicianID, userCreateInput.ManagerID, userCreateInput.Description)
	if err != nil {
		return TaskCreateOutput{}, err
	}

	id, err := u.taskRepository.Save(task)
	if err != nil {
		return TaskCreateOutput{}, err
	}

	output := TaskCreateOutput{
		TaskID: id,
	}

	return output, nil
}
