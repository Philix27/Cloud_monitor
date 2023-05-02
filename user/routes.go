package user

import (
	"github.com/gofiber/fiber/v2"
)


func Routes(app *fiber.App)   {
	app.Get("/user", GetAll)
	app.Get("/user/:id", GetOne)
	app.Post("/user", Create)
	app.Put("/user/:id", Update)
	app.Delete("/user/:id", Delete)
}