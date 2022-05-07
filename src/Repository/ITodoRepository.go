package Repository

import (
	"github.com/strikersk/go-mongo/src/Entity"
)

var mainRepository ITodoRepository

func SetMainRepository(input ITodoRepository) {
	mainRepository = input
}

func GetRepository() ITodoRepository {
	return mainRepository
}

type ITodoRepository interface {
	FindAll() []Entity.TodoStructure
	CreateTodo(Entity.TodoStructure) error
	ReadTodo(string) (Entity.TodoStructure, error)
	UpdateTodo(Entity.TodoStructure) error
	DeleteTodo(string) error
}
