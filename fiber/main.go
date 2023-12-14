package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// RenderForm renders the HTML form.
func RenderForm(c *fiber.Ctx) error {
	return c.Render("form", fiber.Map{})
}

// ProcessForm processes the form submission.
func ProcessForm(c *fiber.Ctx) error {
	name := c.FormValue("name")
	greeting := fmt.Sprintf("Hello, %s!", name)
	return c.Render("greeting", fiber.Map{"Greeting": greeting})
}

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Serve static files (HTML templates and stylesheets).
	app.Static("/", "./static")

	// Define routes.
	app.Get("/", RenderForm)
	app.Post("/submit", ProcessForm)

	// Start the Fiber app on port 8080.
	app.Listen(":8080")
}
