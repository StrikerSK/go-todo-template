package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-todo-template/src/ports"
)

type TodoHandler struct {
	service ports.ITodoService
}

func NewTodoHandler(service ports.ITodoService) TodoHandler {
	return TodoHandler{
		service: service,
	}
}

func (h TodoHandler) EnrichRouter(app *fiber.App) {
	todoPath := app.Group("/api/todo")
	todoPath.Get("/:id", h.service.ReadTodo)
	todoPath.Get("", h.service.ReadTodos)
	todoPath.Post("", h.service.CreateTodo)
	todoPath.Put("/:id", h.service.UpdateTodo)
	todoPath.Delete("/:id", h.service.DeleteTodo)
}
