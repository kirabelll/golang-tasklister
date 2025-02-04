// cli/cli_handler_test.go
package cli

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/todo-app/entities"
)

// MockTaskUseCase implements TaskUseCase for testing purposes
type MockTaskUseCase struct {
	tasks  []entities.Task
	nextID int
}

func (m *MockTaskUseCase) AddTask(title string) (entities.Task, error) {
	task := entities.Task{ID: m.nextID, Title: title, IsDone: false}
	m.tasks = append(m.tasks, task)
	m.nextID++
	return task, nil
}

func (m *MockTaskUseCase) CompleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks[i].IsDone = true
			return nil
		}
	}
	return errors.New("task not found")
}

func (m *MockTaskUseCase) DeleteTask(id int) error {
	for i, task := range m.tasks {
		if task.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

func (m *MockTaskUseCase) ListTasks() []entities.Task {
	return m.tasks
}

func TestCLIHandler(t *testing.T) {
	mockUseCase := &MockTaskUseCase{nextID: 1}
	cliHandler := NewCLIHandler(mockUseCase)

	input := "add\nSample Task\nlist\nexit\n"
	output := &bytes.Buffer{}

	cliHandler.Start(strings.NewReader(input), output)

	expectedOutput := `
Choose an action: add, complete, delete, list, or exit
-> Enter task title: Task added with ID 1

Choose an action: add, complete, delete, list, or exit
-> Tasks:
ID: 1, Title: Sample Task, Status: Pending

Choose an action: add, complete, delete, list, or exit
-> Exiting...
`

	if output.String() != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, output.String())
	}
}
