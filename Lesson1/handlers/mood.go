package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson1/db"
	"github.com/rs-karasal/bitlab_go/lesson1/models"
)

func GetMoods(c *fiber.Ctx) error {
	var moods []models.Mood
	db.DB.Find(&moods)
	return c.JSON(moods)
}

func GetMoodByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var mood models.Mood
	result := db.DB.First(&mood, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" || result.RowsAffected == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Failed to get mood",
			})
		}
	}

	return c.JSON(mood)
}

func CreateMood(c *fiber.Ctx) error {
	mood := new(models.Mood)
	if err := c.BodyParser(mood); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	db.DB.Create(&mood)
	return c.Status(fiber.StatusCreated).JSON(mood)
}

func DeleteMood(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	result := db.DB.Delete(&models.Mood{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Mood not found",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
