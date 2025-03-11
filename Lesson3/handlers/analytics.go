package handlers

import (
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson3/db"
	"github.com/rs-karasal/bitlab_go/lesson3/models"
)

func GetMoodsFiltered(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	fromStr := c.Query("from")
	toStr := c.Query("to")

	var moods []models.Mood
	query := db.DB.Where("user_id = ?", userID)

	if fromStr != "" && toStr != "" {
		from, err1 := time.Parse("2006-01-02", fromStr)
		to, err2 := time.Parse("2006-01-02", toStr)
		if err1 != nil || err2 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD"})
		}
		to = to.Add(24 * time.Hour)
		query = query.Where("created_at BETWEEN ? AND ?", from, to)
	}

	query.Order("created_at desc").Find(&moods)
	return c.JSON(moods)
}

func GetMoodAnalytics(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var moods []models.Mood
	db.DB.Where("user_id = ?", userID).Find(&moods)

	if len(moods) == 0 {
		return c.JSON(fiber.Map{"message": "No data yet"})
	}

	var totalValue int
	var bestDay, worstDay models.Mood
	bestDay.Value = 0
	worstDay.Value = 11

	for _, mood := range moods {
		totalValue += int(mood.Value)

		if mood.Value > bestDay.Value {
			bestDay = mood
		}
		if mood.Value < worstDay.Value {
			worstDay = mood
		}
	}

	average := float64(totalValue) / float64(len(moods))

	return c.JSON(fiber.Map{
		"total":     len(moods),
		"average":   math.Round(average*100) / 100,
		"best_day":  bestDay.CreatedAt.Format("2006-01-02"),
		"worst_day": worstDay.CreatedAt.Format("2006-01-02"),
	})
}
