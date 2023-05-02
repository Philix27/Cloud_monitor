package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)


func GetAll(c *fiber.Ctx) error  {
	fmt.Println("Get all users")
	fmt.Println(c.BaseURL())
	return c.SendString("Get all users service")
}

func GetOne(c *fiber.Ctx) error  {
	fmt.Println("Get all users")
	fmt.Println(c.BaseURL())
	return c.SendString("Get one user service")
}

func Create(c *fiber.Ctx) error  {
	fmt.Println("Create User")
	fmt.Println(c.BaseURL())
	return c.SendString("Create new user service")
}

func Update(c *fiber.Ctx) error  {
	fmt.Println("Update user")
	fmt.Println(c.BaseURL())
	return c.SendString("Update a user service")
}

func Delete(c *fiber.Ctx) error  {
	fmt.Println("Delete user")
	fmt.Println(c.BaseURL())
	return c.SendString("Delete a user")
}