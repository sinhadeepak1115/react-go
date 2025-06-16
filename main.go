package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	app := fiber.New()

	app.Get("/api/todo", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	log.Fatal(app.Listen(":3000"))
}

// getTodos handles the GET request to fetch all todos.
func getTodos(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"msg": "Hi Welcome to the Todo API!"})
}

// createTodo handles the POST request to create a new todo.
func createTodo(c *fiber.Ctx) error {
	todo := Todo{}
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if todo.Body == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	return c.Status(http.StatusCreated).JSON(todo)
}

// updateTodo handles the PATCH request to update an existing todo by ID.
func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	var updatedTodo Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Todo updated successfully",
		"id":   id,
		"todo": updatedTodo,
	})
}

// deleteTodo handles the DELETE request to remove a todo by ID.
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete todo"})
	}
}
