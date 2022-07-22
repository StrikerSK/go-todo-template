package Repository

import (
	"github.com/jaswdr/faker"
	"github.com/strikersk/go-todo-template/src/Entity"
	"log"
)

type LocalTodoRepository struct {
	faker faker.Faker
}

func NewLocalTodoRepository() LocalTodoRepository {
	return LocalTodoRepository{
		faker: faker.New(),
	}
}

func (r *LocalTodoRepository) ReadTodo(ID string) (outputResult Entity.TodoEntity, err error) {
	petOne := r.faker.Pet().Name()
	petTwo := r.faker.Pet().Name()

	log.Printf("User provided ID to read: %s\n", ID)
	return Entity.TodoEntity{
		Id: ID,
		TaskCore: Entity.TaskCore{
			Name:        r.faker.Person().FirstName(),
			Description: r.faker.Person().LastName(),
			Done:        r.faker.Bool(),
		},
		SubTasks: []Entity.TaskCore{
			{
				Name:        petOne,
				Description: petOne,
				Done:        r.faker.Bool(),
			},
			{
				Name:        petTwo,
				Description: petTwo,
				Done:        r.faker.Bool(),
			},
		},
	}, nil
}

func (r *LocalTodoRepository) FindAll() []Entity.TodoEntity {
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

func (r *LocalTodoRepository) CreateTodo(inputTask Entity.TodoEntity) (err error) {
	log.Println("User provide new Task input: ", inputTask)
	return
}

func (r *LocalTodoRepository) UpdateTodo(inputTask Entity.TodoEntity) (err error) {
	log.Println("User provide updated Task input for ID [", inputTask.Id, "]: ", inputTask)
	return
}

func (r *LocalTodoRepository) DeleteTodo(ID string) (err error) {
	log.Printf("User provided ID to delete: %s\n", ID)
	return nil
}
