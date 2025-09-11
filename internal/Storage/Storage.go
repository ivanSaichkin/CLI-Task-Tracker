package storage

import (
	"encoding/json"
	"fmt"
	task "ivan/CLI-Task-Tracker/internal/Task"
	"os"
)

type Storage interface {
	Add(task task.Task) error
	List() ([]task.Task, error)
	Update(id int, newDescription string) error
	GetByID(id int) (*task.Task, error)
	Delete(id int) error
	GetComplitedTasks() ([]task.Task, error)
	GetPendingTasks() ([]task.Task, error)
	SetTaskStatusComplited(id int) error
}

type JSONStorage struct {
	filePath string
	Tasks    map[int]task.Task `json:"Tasks"`
	NextID   int               `json:"NextID"`
}

func NewJSONStorage(filePath string) (*JSONStorage, error) {
	storage := &JSONStorage{
		filePath: filePath,
		Tasks:    make(map[int]task.Task),
		NextID:   1,
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return storage, nil
		}
		return nil, err
	}

	if len(file) == 0 {
		return storage, nil
	}

	err = json.Unmarshal(file, storage)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *JSONStorage) Save() error {
	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}

	if err = os.Truncate(s.filePath, 0); err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *JSONStorage) Add(task task.Task) error {
	task.ID = s.NextID
	s.NextID++
	s.Tasks[task.ID] = task

	return s.Save()
}

func (s *JSONStorage) List() ([]task.Task, error) {
	tasks := make([]task.Task, 0, len(s.Tasks))
	for _, t := range s.Tasks {
		tasks = append(tasks, t)
	}

	if len(tasks) == 0 {
		return []task.Task{}, fmt.Errorf("no tasks here")
	}

	return tasks, nil
}

func (s *JSONStorage) GetByID(id int) (*task.Task, error) {
	task, ok := s.Tasks[id]
	if !ok {
		return nil, fmt.Errorf("no task with id %v", id)
	}
	return &task, nil
}

func (s *JSONStorage) Update(id int, newDescription string) error {
	task, err := s.GetByID(id)
	if err != nil {
		return err
	}

	err = task.UpdateTaskDescription(newDescription)
	if err != nil {
		return err
	}

	s.Tasks[id] = *task

	return s.Save()
}

func (s *JSONStorage) Delete(id int) error {
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}

	delete(s.Tasks, id)

	return s.Save()
}

func (s *JSONStorage) GetComplitedTasks() ([]task.Task, error) {
	complitedTasks := make([]task.Task, 0)

	for _, task := range s.Tasks {
		if task.Status == "completed" {
			complitedTasks = append(complitedTasks, task)
		}
	}

	if len(complitedTasks) == 0 {
		return []task.Task{}, fmt.Errorf("no complited tasks")
	}

	return complitedTasks, nil
}

func (s *JSONStorage) GetPendingTasks() ([]task.Task, error) {
	pendingTasks := make([]task.Task, 0)

	for _, task := range s.Tasks {
		if task.Status == "pending" {
			pendingTasks = append(pendingTasks, task)
		}
	}

	if len(pendingTasks) == 0 {
		return []task.Task{}, fmt.Errorf("no pending tasks")
	}

	return pendingTasks, nil
}

func (s *JSONStorage) SetTaskStatusComplited(id int) error {
	task, err := s.GetByID(id)
	if err != nil {
		return err
	}

	err = task.SetStatusComplited()
	if err != nil {
		return err
	}

	s.Tasks[id] = *task

	return s.Save()
}
