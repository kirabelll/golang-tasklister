package usecases

import (
    "testing"
)

func TestAddTask(t *testing.T) {
    useCase := NewTaskUseCase()

    task, err := useCase.AddTask("Test task")
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }

    if task.ID != 1 {
        t.Errorf("expected task ID 1, got %d", task.ID)
    }

    if task.Title != "Test task" {
        t.Errorf("expected title 'Test task', got %s", task.Title)
    }

    if task.IsDone {
        t.Errorf("expected task to be not done, got done")
    }
}

func TestCompleteTask(t *testing.T) {
    useCase := NewTaskUseCase()
    task, _ := useCase.AddTask("Complete this task")

    err := useCase.CompleteTask(task.ID)
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }

    if !useCase.ListTasks()[0].IsDone {
        t.Errorf("expected task to be completed, but it was not")
    }
}

func TestCompleteTaskNotFound(t *testing.T) {
    useCase := NewTaskUseCase()

    err := useCase.CompleteTask(999) // ID that does not exist
    if err == nil {
        t.Error("expected error for non-existent task, got none")
    }
}

func TestDeleteTask(t *testing.T) {
    useCase := NewTaskUseCase()
    task, _ := useCase.AddTask("Delete this task")

    err := useCase.DeleteTask(task.ID)
    if err != nil {
        t.Errorf("expected no error, got %v", err)
    }

    if len(useCase.ListTasks()) != 0 {
        t.Errorf("expected task list to be empty, but it was not")
    }
}

func TestDeleteTaskNotFound(t *testing.T) {
    useCase := NewTaskUseCase()

    err := useCase.DeleteTask(999) // ID that does not exist
    if err == nil {
        t.Error("expected error for non-existent task, got none")
    }
}

func TestListTasks(t *testing.T) {
    useCase := NewTaskUseCase()
    useCase.AddTask("Task 1")
    useCase.AddTask("Task 2")

    tasks := useCase.ListTasks()
    if len(tasks) != 2 {
        t.Errorf("expected 2 tasks, got %d", len(tasks))
    }

    if tasks[0].Title != "Task 1" || tasks[1].Title != "Task 2" {
        t.Error("expected tasks to match added titles")
    }
}
