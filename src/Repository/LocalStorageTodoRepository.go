package Repository

import (
	"fmt"
	"github.com/strikersk/go-todo-template/src/Entity"
	"log"
)

type LocalStorageTodoRepository struct {
	data []Entity.TodoEntity
}

func NewLocalStorageTodoRepository() LocalStorageTodoRepository {
	return LocalStorageTodoRepository{
		data: []Entity.TodoEntity{},
	}
}

func (r *LocalStorageTodoRepository) ReadTodo(ID string) (outputResult Entity.TodoEntity, err error) {
	log.Printf("User provided ID to read: %s\n", ID)

	for _, item := range r.data {
		if item.Id == ID {
			return item, nil
		}
	}

	return Entity.TodoEntity{}, fmt.Errorf("todo [%s] not found", ID)
}

func (r *LocalStorageTodoRepository) FindAll() []Entity.TodoEntity {
	return r.data
}

func (r *LocalStorageTodoRepository) CreateTodo(inputTask Entity.TodoEntity) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	r.data = append(r.data, inputTask)
	return
}

func (r *LocalStorageTodoRepository) UpdateTodo(inputTask Entity.TodoEntity) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r *LocalStorageTodoRepository) DeleteTodo(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
