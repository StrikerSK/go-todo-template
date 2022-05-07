package Service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/strikersk/go-mongo/src/Entity"
	"github.com/strikersk/go-mongo/src/Repository"
	"log"
)

func ReadTodo(c *fiber.Ctx) Entity.TodoStructure {
	parameter := c.Params("id")

	todo, err := Repository.GetRepository().ReadTodo(parameter)
	if err != nil {
		log.Println(err)
	}

	return todo
}

func CreateTodo(c *fiber.Ctx) string {
	var tmpTodo Entity.TodoStructure
	_ = c.BodyParser(&tmpTodo)

	tmpTodo.Id = uuid.New().String()
	_ = Repository.GetRepository().CreateTodo(tmpTodo)

	return tmpTodo.Id
}

func UpdateTodo(c *fiber.Ctx) {
	var tmpTodo Entity.TodoStructure

	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTodo.Id = c.Params("id")
	_ = Repository.GetRepository().UpdateTodo(tmpTodo)
}

func FindTasks() []Entity.TodoStructure {
	return Repository.GetRepository().FindAll()
}
