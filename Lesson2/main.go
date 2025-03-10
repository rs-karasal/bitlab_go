package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson2/db"
	"github.com/rs-karasal/bitlab_go/lesson2/handlers"
	"github.com/rs-karasal/bitlab_go/lesson2/middleware"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	mood := app.Group("/moods", middleware.Protected())
	mood.Get("/", handlers.GetMyMoods)
	mood.Get("/:id", handlers.GetMyMoodByID)
	mood.Post("/", handlers.CreateMood)
	mood.Delete("/:id", handlers.DeleteMood)

	app.Listen(":7777")
}
