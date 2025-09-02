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
	Update(id int, updatedTask task.Task) error
	GetByID(id int) (*task.Task, error)
	Delete(id int) error
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
		return nil, err
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

func (s *JSONStorage) Update(id int, updatedTask task.Task) error {
	_, err := s.GetByID(id)
	if err != nil {
		return err
	}

	s.Tasks[id] = updatedTask

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
