package Service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-mongo/src/Entity"
	"github.com/strikersk/go-mongo/src/repository"
)

func CreateTodo(c *fiber.Ctx) string {
	var tmpTodo Entity.TodoStructure
	_ = c.BodyParser(&tmpTodo)
	return repository.CreateTodo(tmpTodo)
}
