package usecase

import (
	"errors"

	"github.com/LidoHon/TodoCli/internal/model"
	"github.com/LidoHon/TodoCli/internal/repo"
)

type TodoService struct{
	todos []model.Todo
	fileStorage *repo.FileStorage
}


func NewTodoService(todos []model.Todo, fileStorage *repo.FileStorage) *TodoService{
	return &TodoService{todos: todos, fileStorage: fileStorage}

}


// create a todo

func (s *TodoService) CreateTodo(title string)error {
	todo := model.Todo{
		ID: len(s.todos) + 1,
		Title: title,
		Completed: false,
	}
	s.todos = append(s.todos, todo)
	return s.fileStorage.SaveTodos(s.todos)
}


// update to-do status
func (s *TodoService) UpdateTodo(id int) error {
	for i, todo := range s.todos{
		if todo.ID == id{
			s.todos[i].Completed = !s.todos[i].Completed
			return s.fileStorage.SaveTodos(s.todos)
		}
	}
	return errors.New("todo not found")
}


// get to-dos
func (s *TodoService) GetTodo() []model.Todo{
	return s.todos
}

// get to-do by id
func (s *TodoService) GetTodoById(id int) (model.Todo, error){
	for _, todo := range s.todos{
		if todo.ID == id{
			return todo, nil
		}
	}
	return model.Todo{}, errors.New("todo not found")
}

// delete Todos

func (s *TodoService) DeleteTodo(id int) error {
	for i, todo := range s.todos{
		if todo.ID == id{
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return s.fileStorage.SaveTodos(s.todos)
		}
	}
	return errors.New("todo not found")
}


