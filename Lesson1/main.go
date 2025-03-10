package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson1/db"
	"github.com/rs-karasal/bitlab_go/lesson1/handlers"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Get("/moods", handlers.GetMoods)
	app.Get("/moods/:id", handlers.GetMoodByID)
	app.Post("/moods", handlers.CreateMood)
	app.Delete("/moods/:id", handlers.DeleteMood)

	app.Listen(":7777")
}
