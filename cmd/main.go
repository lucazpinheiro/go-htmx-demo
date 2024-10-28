package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lucazpinheiro/go-plus-htmx-demo/internal"

	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	tasksRepo := internal.NewTaskRepository()

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
			"Tasks": tasksRepo.ListTask(),
		})
	})
	app.Put("/task/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).SendString("ID inválido")
		}

		tasksRepo.FlipTaskStatus(id)

		tasks := tasksRepo.ListTask()

		return c.Render("ul", fiber.Map{
			"Tasks": tasks,
		})

	})
	app.Post("/task", func(c *fiber.Ctx) error {
		description := c.FormValue("task-description")

		tasksRepo.CreateTask(description)

		return c.Render("ul", fiber.Map{
			"Tasks": tasksRepo.ListTask(),
		})
	})

	log.Fatal(app.Listen(":3000"))
}
