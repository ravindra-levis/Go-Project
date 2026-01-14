package storage

import (
	"example/hello/task-api/models"
	"fmt"
)

type MemoryStore struct {
	tasks  []models.Task
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks:  []models.Task{},
		nextID: 1,
	}
}
func (m *MemoryStore) Create(task models.Task) (models.Task, error) {
	task.ID = m.nextID
	m.nextID++
	m.tasks = append(m.tasks, task)
	return task, nil
}
func (m *MemoryStore) GetAll() ([]models.Task, error) {
	return m.tasks, nil
}

func (m *MemoryStore) Update(id int, done bool) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks[i].Done = done
			return nil
		}
	}
	return fmt.Errorf("Task not found")
}
func (m *MemoryStore) Delete(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task not found")
}
