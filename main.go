package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/template/html/v2"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var tasks = []Task{
	{
		ID:          0,
		Description: "A",
		Done:        false,
	},
	{
		ID:          1,
		Description: "B",
		Done:        false,
	},
	{
		ID:          2,
		Description: "C",
		Done:        false,
	},
}

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/task", func(c *fiber.Ctx) error {
		return c.Render("ul", fiber.Map{
			"Tasks": tasks,
		})
	})
	app.Put("/task/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("ID inválido")
		}
		fmt.Println(id)
		for i, task := range tasks {
			if task.ID == id {
				// Atualiza o estado da tarefa com base no checkbox
				tasks[i].Done = !task.Done
				break
			}
		}

		return c.Render("ul", fiber.Map{
			"Tasks": tasks,
		})

	})
	app.Post("/task", func(c *fiber.Ctx) error {
		description := c.FormValue("task-description")

		newTask := Task{
			Description: description,
		}

		setTaskID(&newTask)
		tasks = append(tasks, newTask)

		return c.Render("ul", fiber.Map{
			"Tasks": tasks,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func setTaskID(task *Task) {
	task.ID = len(tasks)
}
