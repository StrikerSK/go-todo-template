package Repository

import (
	"github.com/jaswdr/faker"
	"github.com/strikersk/go-todo-template/src/Entity"
	"log"
	"math/rand"
	"time"
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
	log.Printf("User provided ID to read: %s\n", ID)
	return r.generateTodo(ID), nil
}

func (r *LocalTodoRepository) FindAll() []Entity.TodoEntity {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1).Intn(10)

	var subTasks []Entity.TodoEntity
	for i := 0; i < r1; i++ {
		subTasks = append(subTasks, r.generateTodo(r.faker.UUID().V4()))
	}

	return subTasks
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

func (r *LocalTodoRepository) generateTask() Entity.TaskCore {
	return Entity.TaskCore{
		Name:        r.faker.Person().FirstName(),
		Description: r.faker.Person().LastName(),
		Done:        r.faker.Bool(),
	}
}

func (r *LocalTodoRepository) generateTodo(id string) Entity.TodoEntity {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1).Intn(10)

	var subTasks []Entity.TaskCore
	for i := 0; i < r1; i++ {
		subTasks = append(subTasks, r.generateTask())
	}

	return Entity.TodoEntity{
		Id:       id,
		TaskCore: r.generateTask(),
		SubTasks: subTasks,
	}
}
