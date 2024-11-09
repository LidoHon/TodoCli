package repo

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/LidoHon/TodoCli/internal/model"
)

type FileStorage struct{
	fileName string
}

// so basically what the bellow constructor function would do is create a new fileStorage object and return it
func NewFile(fileName string) *FileStorage{
	return &FileStorage{fileName: fileName}
}


// save to-dos to the file
func (fs *FileStorage) SaveTodos(todos []model.Todo) error {
	file , err := os.Create(fs.fileName)
	if err != nil{
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	encoder := json.NewEncoder(writer)

	if err :=encoder.Encode(todos); err !=nil{
		return err
	}
	return writer.Flush()
}


// load to-dos

func (fs *FileStorage) LoadTodos()([]model.Todo, error){
	file, err := os.Open(fs.fileName)

	if err !=nil{
		if os.IsNotExist(err){
			return []model.Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var todos []model.Todo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err !=nil{
		return nil, err
	}
	return todos, nil
}

