package main

import (
	"fmt"

	"github.com/LidoHon/TodoCli/internal/interfaces"
	"github.com/LidoHon/TodoCli/internal/model"
	"github.com/LidoHon/TodoCli/internal/repo"
	usecase "github.com/LidoHon/TodoCli/internal/useCase"
)

func main(){
	// load to-dos from file
	fileStorage := repo.NewFile("todos.txt")
	todos, err := fileStorage.LoadTodos()
	if err != nil{
		fmt.Println("error loading todos from file:", err)
		todos = []model.Todo{}
		
	}
	// initialize todo service with the loaded todos above
	todoService := usecase.NewTodoService(todos, fileStorage)

	// start cli
	interfaces.RunCli(todoService)
}