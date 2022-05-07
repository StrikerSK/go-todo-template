package Repository

import (
	"github.com/strikersk/go-todo-template/src/Entity"
	"log"
)

type LocalTodoRepository struct{}

func SetLocalRepository() {
	SetMainRepository(&LocalTodoRepository{})
}

func (r *LocalTodoRepository) ReadTodo(ID string) (outputResult Entity.TodoStructure, err error) {
	log.Printf("User provided ID to read: %s\n", ID)
	return Entity.TodoStructure{
		Id: "123",
		TaskCore: Entity.TaskCore{
			Name:        "MainTask",
			Description: "This represents main task",
			Done:        false,
		},
		SubTasks: []Entity.TaskCore{
			{
				Name:        "SubTask1",
				Description: "This should be sub task 1",
				Done:        false,
			},
			{
				Name:        "SubTask2",
				Description: "This should be sub task 2",
				Done:        false,
			},
		},
	}, nil
}

func (r *LocalTodoRepository) FindAll() []Entity.TodoStructure {
	return []Entity.TodoStructure{
		{
			Id: "123",
			TaskCore: Entity.TaskCore{
				Name:        "MainTask",
				Description: "This represents main task",
				Done:        false,
			},
			SubTasks: []Entity.TaskCore{
				{
					Name:        "SubTask1",
					Description: "This should be sub task 1",
					Done:        false,
				},
				{
					Name:        "SubTask2",
					Description: "This should be sub task 2",
					Done:        false,
				},
			},
		},
	}
}

func (r *LocalTodoRepository) CreateTodo(inputTask Entity.TodoStructure) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	return
}

func (r *LocalTodoRepository) UpdateTodo(inputTask Entity.TodoStructure) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r *LocalTodoRepository) DeleteTodo(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
