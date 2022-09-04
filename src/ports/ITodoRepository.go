package ports

import "github.com/strikersk/go-todo-template/src/Entity"

type ITodoRepository interface {
	FindAll() []Entity.TodoEntity
	CreateTodo(Entity.TodoEntity) error
	ReadTodo(string) (Entity.TodoEntity, error)
	UpdateTodo(Entity.TodoEntity) error
	DeleteTodo(string) error
}
