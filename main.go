package main

import (
	"fmt"

	"github.com/Philix27/cloud_monitor/user"
	"github.com/gofiber/fiber/v2"
)


func main() {
	fmt.Println("Hi guys")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error { return c.SendString("Hello Peeps")})
	user.Routes(app)
	app.Listen(":3000")
}
