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
	return Repository.ReadTodo(parameter)
}

func CreateTodo(c *fiber.Ctx) string {
	var tmpTodo Entity.TodoStructure
	_ = c.BodyParser(&tmpTodo)

	tmpTodo.Id = uuid.New().String()
	_ = Repository.CreateTodo(tmpTodo)

	return tmpTodo.Id
}

func UpdateTodo(c *fiber.Ctx) {
	var tmpTodo Entity.TodoStructure

	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTodo.Id = c.Params("id")
	Repository.UpdateTodo(tmpTodo)
}

func FindTasks() []Entity.TodoStructure {
	return Repository.FindAll()
}
