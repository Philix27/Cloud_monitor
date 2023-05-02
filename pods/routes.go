package pods

import "github.com/gofiber/fiber/v2"


func Routes(c fiber.Ctx) error  {
	return c.SendString("")
}