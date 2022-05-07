package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-todo-template/src/Repository"
	"github.com/strikersk/go-todo-template/src/Service"
	"log"
)

func main() {
	Repository.SetLocalRepository()

	app := fiber.New()

	todoPath := app.Group("/api/todo")
	todoPath.Get("/:id", Service.ReadTodo)
	todoPath.Get("", Service.FindTasks)
	todoPath.Post("", Service.CreateTodo)
	todoPath.Put("/:id", Service.UpdateTodo)
	todoPath.Delete("/:id", Service.DeleteTodo)

	log.Fatal(app.Listen(fmt.Sprintf(":8080")))
}
