package datastore

import (
	"errors"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	DueDate     time.Time `json:"due_date"` // Format: YYYY-MM-DD
}

type TaskStore struct {
	tasks  []Task
	nextId int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  []Task{},
		nextId: 1,
	}
}

func (store TaskStore) GetAllTaskes() []Task {
	return store.tasks
}

func (store *TaskStore) AddTask(task Task) Task {
	task.ID = store.nextId
	store.nextId++

	store.tasks = append(store.tasks, task)
	return task
}

func (store *TaskStore) DeleteTask(id int) error {
	for i, task := range store.tasks {
		if id == task.ID {
			store.tasks = append(store.tasks[:i], store.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not Found")
}

func (store *TaskStore) UpdateTask(id int, updatedTask Task) (*Task, error) {
	for i, task := range store.tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				store.tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				store.tasks[i].Description = updatedTask.Description
			}
			if updatedTask.Completed {
				store.tasks[i].Completed = updatedTask.Completed
			}
			if !updatedTask.DueDate.IsZero() {
				store.tasks[i].DueDate = updatedTask.DueDate
			}
			return &store.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
