package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson3/db"
	"github.com/rs-karasal/bitlab_go/lesson3/handlers"
	"github.com/rs-karasal/bitlab_go/lesson3/middleware"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	mood := app.Group("/moods", middleware.Protected())
	mood.Get("/", handlers.GetMyMoods)
	mood.Get("/filter", handlers.GetMoodsFiltered)
	mood.Get("/:id", handlers.GetMyMoodByID)
	mood.Post("/", handlers.CreateMood)
	mood.Delete("/:id", handlers.DeleteMood)

	analytics := app.Group("/analytics", middleware.Protected())
	analytics.Get("/summary", handlers.GetMoodAnalytics)

	app.Listen(":7777")
}
