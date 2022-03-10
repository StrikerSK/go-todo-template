package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-mongo/src/Service"
	"github.com/strikersk/go-mongo/src/repository"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/api/:id", func(c *fiber.Ctx) error {
		tmpPara := c.Params("id")
		return c.JSON(repository.GetTodo(tmpPara))
	})

	app.Post("/api", func(c *fiber.Ctx) error {
		tmpOut := Service.CreateTodo(c)
		return c.Send([]byte(tmpOut))
	})

	app.Put("/api/:id", func(c *fiber.Ctx) error {
		Service.UpdateTodo(c)
		return c.Send([]byte("Updated"))
	})

	log.Fatal(app.Listen(":3000"))
}
