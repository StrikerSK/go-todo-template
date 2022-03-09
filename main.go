package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-mongo/src/Service"
	"github.com/strikersk/go-mongo/src/repository"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/api/:from", func(c *fiber.Ctx) error {
		tmpPara := c.Params("from")
		return c.JSON(repository.GetTodo(tmpPara))
	})

	app.Post("/api/", func(c *fiber.Ctx) error {
		tmpOut := Service.CreateTodo(c)
		return c.Send([]byte(tmpOut))
	})

	log.Fatal(app.Listen(":3000"))
}
