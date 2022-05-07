package Service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/strikersk/go-todo-template/src/Entity"
	"github.com/strikersk/go-todo-template/src/Repository"
	"log"
	"net/http"
)

func ReadTodo(c *fiber.Ctx) error {
	parameter := c.Params("id")

	todo, err := Repository.GetRepository().ReadTodo(parameter)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	var tmpTodo Entity.TodoStructure
	_ = c.BodyParser(&tmpTodo)

	tmpTodo.Id = uuid.New().String()
	_ = Repository.GetRepository().CreateTodo(tmpTodo)

	return c.JSON(map[string]string{"data": tmpTodo.Id})
}

func UpdateTodo(c *fiber.Ctx) error {
	var tmpTodo Entity.TodoStructure
	if err := c.BodyParser(&tmpTodo); err != nil {
		log.Printf("%v\n", err)
	}

	tmpTodo.Id = c.Params("id")
	_ = Repository.GetRepository().UpdateTodo(tmpTodo)
	return c.SendStatus(http.StatusOK)
}

func DeleteTodo(c *fiber.Ctx) error {
	_ = Repository.GetRepository().DeleteTodo(c.Params("id"))
	return c.SendStatus(http.StatusOK)
}

func FindTasks(c *fiber.Ctx) error {
	return c.JSON(Repository.GetRepository().FindAll())
}
