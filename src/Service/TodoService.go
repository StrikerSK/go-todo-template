package Service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/strikersk/go-todo-template/src/Entity"
	"github.com/strikersk/go-todo-template/src/ports"
	"log"
	"net/http"
)

type TodoService struct {
	repository ports.ITodoRepository
}

func NewTodoService(repository ports.ITodoRepository) TodoService {
	return TodoService{
		repository: repository,
	}
}

func (s TodoService) CreateTodo(c *fiber.Ctx) error {
	var tmpTodo Entity.TodoEntity
	_ = c.BodyParser(&tmpTodo)

	tmpTodo.Id = uuid.New().String()
	_ = s.repository.CreateTodo(tmpTodo)

	return c.JSON(map[string]string{"data": tmpTodo.Id})
}

func (s TodoService) ReadTodo(c *fiber.Ctx) error {
	parameter := c.Params("id")

	todo, err := s.repository.ReadTodo(parameter)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(todo)
}

func (s TodoService) ReadTodos(c *fiber.Ctx) error {
	return c.JSON(s.repository.FindAll())
}

func (s TodoService) UpdateTodo(c *fiber.Ctx) error {
	var tmpTodo Entity.TodoEntity
	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTodo.Id = c.Params("id")
	_ = s.repository.UpdateTodo(tmpTodo)
	return c.SendStatus(http.StatusOK)
}

func (s TodoService) DeleteTodo(c *fiber.Ctx) error {
	_ = s.repository.DeleteTodo(c.Params("id"))
	return c.SendStatus(http.StatusOK)
}
