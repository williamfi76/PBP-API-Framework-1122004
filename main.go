package main

import (
	c "fiber-api/controller"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	r := fiber.New()

	r.Get("/users", c.GetAllUsers)
	r.Get("/users/:id", c.GetUserByID)
	r.Post("/users", c.InsertNewUser)
	r.Put("/users/:id", c.UpdateUser)
	r.Delete("/users/:id", c.DeleteUser)

	log.Fatal(r.Listen(":2008"))
	// app.Listen(":2008")
}
