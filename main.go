package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	fmt.Println(hello())

	app.Listen(":3001")
}

func hello() string {
	return "Hello Golangs"
}
