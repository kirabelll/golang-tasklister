// usecases/task_usecase_impl.go
package usecases

import (
	"errors"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/todo-app/entities"
)

type TaskUseCaseImpl struct {
    tasks []entities.Task
    nextID int
}

func NewTaskUseCase() TaskUseCase {
    return &TaskUseCaseImpl{
        tasks: []entities.Task{},
        nextID: 1,
    }
}

func (u *TaskUseCaseImpl) AddTask(title string) (entities.Task, error) {
    task := entities.Task{ID: u.nextID, Title: title, IsDone: false}
    u.tasks = append(u.tasks, task)
    u.nextID++
    return task, nil
}

func (u *TaskUseCaseImpl) CompleteTask(id int) error {
    for i, task := range u.tasks {
        if task.ID == id {
            u.tasks[i].IsDone = true
            return nil
        }
    }
    return errors.New("task not found")
}

func (u *TaskUseCaseImpl) DeleteTask(id int) error {
    for i, task := range u.tasks {
        if task.ID == id {
            u.tasks = append(u.tasks[:i], u.tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}

func (u *TaskUseCaseImpl) ListTasks() []entities.Task {
    return u.tasks
}
