package main

import (
	"os"

	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/todo-app/interfaces/cli"
	"github.com/dawitel/the-ultimate-golang-docker-course-for-beginners/module-two/final-projects/todo-app/usecases"
)

func main() {
	taskUseCase := usecases.NewTaskUseCase()
	cliHandler := cli.NewCLIHandler(taskUseCase)
	cliHandler.Start(os.Stdin, os.Stdout)
}
