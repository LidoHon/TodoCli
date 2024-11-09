package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	usecase "github.com/LidoHon/TodoCli/internal/useCase"
)

// ANSI color codes
const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Underline = "\033[4m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Blue    = "\033[34m"
	Purple  = "\033[35m"
	Cyan    = "\033[36m"


)


func RunCli(todoService *usecase.TodoService){
	scanner :=bufio.NewScanner(os.Stdin)

	for{
		fmt.Println(Purple+ Bold+ "\n choose a command : create | update | delete | listAll | listOne | quit: " + Reset)

		scanner.Scan()
		input := scanner.Text()

		switch input {
			case "create":
				fmt.Print(Blue+ "enter todo title: " +Reset)
				scanner.Scan()
				title := scanner.Text()
				err := todoService.CreateTodo(title)
				if err != nil{
					fmt.Println(err)
				}else{
					fmt.Println( Green + "todo created successfully" + Reset)
				}

			case "update":
				fmt.Print(Blue +"enter todo id to update: " + Reset)
				scanner.Scan()
				id,_ := strconv.Atoi(scanner.Text())
				if err := todoService.UpdateTodo(id); err != nil{
					fmt.Println(err)
				}
				fmt.Println(Green + "todo marked completed successfully" + Reset)
				

			case "delete":
				fmt.Print(Blue + "enter todo id to delete: " + Reset)
				scanner.Scan()
				id,_ := strconv.Atoi(scanner.Text())
				if err := todoService.DeleteTodos(id); err != nil{
					fmt.Println(err)
				}
				fmt.Println(Green + "todo deleted successfully" + Reset)
				

			case "listAll":
				todos := todoService.GetTodo()
				fmt.Println("todos: ")
				for _, todo := range todos{
					status := "not completed"
					if todo.Completed{
						status = "completed"
					}
					fmt.Printf(Cyan + "ID: %d, Title: %s, Status: %s"+Reset+"\n", todo.ID, todo.Title, status + Reset)
				}

			case "listOne":
				fmt.Print(Blue + "enter todo id to get: " + Reset)
				scanner.Scan()
				id,_ := strconv.Atoi(scanner.Text())
				todo, err := todoService.GetTodoById(id)
				if err != nil{
					fmt.Println(err)
				}
					fmt.Printf(Cyan+"ID: %d, Title: %s, Completed: %t"+Reset+"\n", todo.ID, todo.Title, todo.Completed)

				

			case "quit":
				return

			default:
				fmt.Println(Red + "invalid command" + Reset)
				
		}
	}
}