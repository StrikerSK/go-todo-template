package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/strikersk/go-todo-template/src/Repository"
	"github.com/strikersk/go-todo-template/src/Service"
	Handler "github.com/strikersk/go-todo-template/src/handler"
	"log"
)

func main() {
	app := fiber.New()

	repository := Repository.NewLocalTodoRepository()
	service := Service.NewTodoService(&repository)
	handler := Handler.NewTodoHandler(service)

	handler.EnrichRouter(app)

	err := app.Listen(fmt.Sprintf(":8080"))
	if err != nil {
		log.Fatalln(err)
	}
}
