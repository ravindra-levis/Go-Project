package storage

import "example/hello/task-api/models"

type TaskStore interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	Update(id int, done bool) error
	Delete(id int) error
}
