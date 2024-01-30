package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	Open   Status = "open"
	Closed Status = "closed"
)

type Task struct {
	ID           string
	TechnicianID string
	ManagerID    string
	Description  string
	Status       Status
	CreatedAt    time.Time
	ModifiedAt   time.Time
	ClosedAt     time.Time
}

func NewTask(technicianID, managerID, description string) (Task, error) {

	if technicianID == "" {
		return Task{}, errors.New("technician ID is required")
	}

	if managerID == "" {
		return Task{}, errors.New("manager ID is required")
	}

	if description == "" {
		return Task{}, errors.New("description is required")
	}

	task := Task{
		ID:           uuid.New().String(),
		TechnicianID: technicianID,
		ManagerID:    managerID,
		Description:  description,
		Status:       Open,
		CreatedAt:    time.Now(),
		ModifiedAt:   time.Now(),
		ClosedAt:     time.Time{},
	}

	return task, nil
}
