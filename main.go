package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID          int    `json:"id"`
	Compleleted bool   `json:"completed"`
	Body        string `json:"body"`
}

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	fmt.Println("hello lets go and build go")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"err": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/update/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		//todoId, _ := strconv.Atoi(id)

		for i, value := range todos {
			if fmt.Sprint(value.ID) == id {
				todos[i].Compleleted = !todos[i].Compleleted
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "Todo not found"})
	})

	app.Delete("/api/delete/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, value := range todos {
			if fmt.Sprint(value.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "deleted succesfuly", "todos": todos})
			}
		}

		return c.Status(404).JSON(fiber.Map{"err": "Todo not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
