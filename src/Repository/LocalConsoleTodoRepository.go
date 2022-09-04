package Repository

import (
	"github.com/strikersk/go-todo-template/src/Entity"
	"log"
)

type LocalConsoleTodoRepository struct{}

func NewLocalConsoleTodoRepository() LocalConsoleTodoRepository {
	return LocalConsoleTodoRepository{}
}

func (r *LocalConsoleTodoRepository) ReadTodo(ID string) (outputResult Entity.TodoEntity, err error) {
	log.Printf("User provided ID to read: %s\n", ID)
	return Entity.TodoEntity{
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

func (r *LocalConsoleTodoRepository) FindAll() []Entity.TodoEntity {
	return []Entity.TodoEntity{
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

func (r *LocalConsoleTodoRepository) CreateTodo(inputTask Entity.TodoEntity) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	return
}

func (r *LocalConsoleTodoRepository) UpdateTodo(inputTask Entity.TodoEntity) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r *LocalConsoleTodoRepository) DeleteTodo(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
