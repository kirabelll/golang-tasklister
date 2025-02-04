package usecases

import "github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/todo-app/entities"

type TaskUseCase interface {
	AddTask(title string) (entities.Task, error)
	CompleteTask(id int) error
	DeleteTask(id int) error
	ListTasks() []entities.Task
}
