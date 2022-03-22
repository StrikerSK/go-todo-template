package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-mongo/src/Repository"
	"github.com/strikersk/go-mongo/src/Service"
	"log"
	"net/http"
	"os"
)

func main() {
	Repository.GetDatabaseInstance()

	app := fiber.New()

	todoPath := app.Group("/api/todo")

	todoPath.Get("/:id", func(c *fiber.Ctx) error {
		todo := Service.ReadTodo(c)
		return c.JSON(todo)
	})

	todoPath.Get("", func(c *fiber.Ctx) error {
		tasks := Service.FindTasks()
		return c.JSON(tasks)
	})

	todoPath.Post("", func(c *fiber.Ctx) error {
		tmpOut := Service.CreateTodo(c)
		return c.JSON(map[string]string{"data": tmpOut})
	})

	todoPath.Put("/:id", func(c *fiber.Ctx) error {
		Service.UpdateTodo(c)
		return c.SendStatus(http.StatusOK)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
