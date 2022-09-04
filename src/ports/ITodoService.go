package ports

import (
	"github.com/gofiber/fiber/v2"
)

type ITodoService interface {
	CreateTodo(c *fiber.Ctx) error
	ReadTodo(c *fiber.Ctx) error
	ReadTodos(c *fiber.Ctx) error
	UpdateTodo(c *fiber.Ctx) error
	DeleteTodo(c *fiber.Ctx) error
}
