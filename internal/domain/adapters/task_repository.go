package adapters

import "github.com/uiansol/task-follow-up/internal/domain/entities"

type ITaskRepository interface {
	Save(task entities.Task) (string, error)
}
