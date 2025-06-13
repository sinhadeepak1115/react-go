package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int
	Completed bool
	Body      string
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"msg": "Welcome to the Todo API!"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := Todo{}

		if err := c.BodyParser(&todo); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if todo.Body == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Todo body cannot be empty"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, todo)
		return c.Status(http.StatusCreated).JSON(todo)
	})

	log.Fatal(app.Listen(":3000"))
}
