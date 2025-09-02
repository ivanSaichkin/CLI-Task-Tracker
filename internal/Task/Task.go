package task

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusCompleted Status = "completed"
)

type Task struct {
	ID          int       `json:"ID"`
	Description string    `json:"Description"`
	Status      Status    `json:"Status"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
	CompletedAt time.Time `json:"ComplitedAt,omitempty"`
}

func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) UpdateTaskDescription(description string) error {
	if t.Description == description {
		return fmt.Errorf("no difference between old and new")
	}

	t.Description = description
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) SetStatusComplited() error {
	if t.Status == StatusCompleted {
		return fmt.Errorf("task has already been completed")
	}

	t.Status = StatusCompleted
	t.UpdatedAt = time.Now()
	t.CompletedAt = time.Now()
	return nil
}
