package Service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-mongo/src/Entity"
	"github.com/strikersk/go-mongo/src/repository"
	"log"
)

func CreateTodo(c *fiber.Ctx) string {
	var tmpTodo Entity.TodoStructure
	_ = c.BodyParser(&tmpTodo)
	return repository.CreateTodo(tmpTodo)
}

func UpdateTodo(c *fiber.Ctx) {
	var tmpTodo Entity.TodoStructure

	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	id := c.Params("id")
	repository.UpdateTodo(id, tmpTodo)
}
