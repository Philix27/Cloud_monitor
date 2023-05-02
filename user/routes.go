package user

import (
	"github.com/gofiber/fiber/v2"
)


func Routes(app *fiber.App)   {
	app.Get("/user", GetOne)
	app.Post("/user", Create)
	app.Put("/user", Update)
	app.Delete("/user", Delete)
}