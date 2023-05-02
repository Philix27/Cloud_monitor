package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)



func Controller(c *fiber.Ctx) error  {
	fmt.Println("Hello guys")
	fmt.Println(c.BaseURL())
	return c.SendString("Welcome to my go server")
}