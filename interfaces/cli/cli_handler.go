// cli/cli_handler.go
package cli

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/todo-app/usecases"
)

type CLIHandler struct {
    useCase usecases.TaskUseCase
}

func NewCLIHandler(u usecases.TaskUseCase) *CLIHandler {
    return &CLIHandler{useCase: u}
}

// Start takes io.Reader and io.Writer to facilitate testing.
func (h *CLIHandler) Start(input io.Reader, output io.Writer) {
    scanner := bufio.NewScanner(input)
    writer := bufio.NewWriter(output)

    for {
        fmt.Fprintln(writer, "\nChoose an action: add, complete, delete, list, or exit")
        fmt.Fprint(writer, "-> ")
        writer.Flush()

        if !scanner.Scan() {
            break
        }

        input := strings.TrimSpace(scanner.Text())
        switch input {
        case "add":
            h.addTask(scanner, writer)
        case "complete":
            h.completeTask(scanner, writer)
        case "delete":
            h.deleteTask(scanner, writer)
        case "list":
            h.listTasks(writer)
        case "exit":
            fmt.Fprintln(writer, "Exiting...")
            writer.Flush()
            return
        default:
            fmt.Fprintln(writer, "Unknown command. Please choose add, complete, delete, list, or exit.")
            writer.Flush()
        }
    }
}

func (h *CLIHandler) addTask(scanner *bufio.Scanner, writer io.Writer) {
    fmt.Fprint(writer, "Enter task title: ")
    writer.(*bufio.Writer).Flush()

    if !scanner.Scan() {
        fmt.Fprintln(writer, "Error reading input.")
        return
    }

    title := strings.TrimSpace(scanner.Text())
    task, err := h.useCase.AddTask(title)
    if err != nil {
        fmt.Fprintln(writer, "Error adding task:", err)
    } else {
        fmt.Fprintf(writer, "Task added with ID %d\n", task.ID)
    }
    writer.(*bufio.Writer).Flush()
}

func (h *CLIHandler) completeTask(scanner *bufio.Scanner, writer io.Writer) {
    fmt.Fprint(writer, "Enter task ID to complete: ")
    writer.(*bufio.Writer).Flush()

    if !scanner.Scan() {
        fmt.Fprintln(writer, "Error reading input.")
        return
    }

    idStr := strings.TrimSpace(scanner.Text())
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Fprintln(writer, "Invalid ID format")
        return
    }
    if err := h.useCase.CompleteTask(id); err != nil {
        fmt.Fprintln(writer, "Error completing task:", err)
    } else {
        fmt.Fprintln(writer, "Task marked as complete")
    }
    writer.(*bufio.Writer).Flush()
}

func (h *CLIHandler) deleteTask(scanner *bufio.Scanner, writer io.Writer) {
    fmt.Fprint(writer, "Enter task ID to delete: ")
    writer.(*bufio.Writer).Flush()

    if !scanner.Scan() {
        fmt.Fprintln(writer, "Error reading input.")
        return
    }

    idStr := strings.TrimSpace(scanner.Text())
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Fprintln(writer, "Invalid ID format")
        return
    }
    if err := h.useCase.DeleteTask(id); err != nil {
        fmt.Fprintln(writer, "Error deleting task:", err)
    } else {
        fmt.Fprintln(writer, "Task deleted")
    }
    writer.(*bufio.Writer).Flush()
}

func (h *CLIHandler) listTasks(writer io.Writer) {
    tasks := h.useCase.ListTasks()
    if len(tasks) == 0 {
        fmt.Fprintln(writer, "No tasks found.")
        writer.(*bufio.Writer).Flush()
        return
    }

    fmt.Fprintln(writer, "Tasks:")
    for _, task := range tasks {
        status := "Pending"
        if task.IsDone {
            status = "Completed"
        }
        fmt.Fprintf(writer, "ID: %d, Title: %s, Status: %s\n", task.ID, task.Title, status)
    }
    writer.(*bufio.Writer).Flush()
}
