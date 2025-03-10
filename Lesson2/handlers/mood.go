package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson2/db"
	"github.com/rs-karasal/bitlab_go/lesson2/models"
	"gorm.io/gorm"
)

func GetMyMoods(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var moods []models.Mood
	db.DB.Where("user_id = ?", userID).Find(&moods)
	return c.JSON(moods)
}

func GetMyMoodByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
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
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Failed to get mood",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	if mood.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	return c.JSON(mood)
}

func CreateMood(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	mood := new(models.Mood)

	if err := c.BodyParser(mood); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	mood.UserID = userID
	db.DB.Create(&mood)
	return c.Status(fiber.StatusCreated).JSON(mood)
}

func DeleteMood(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var mood models.Mood
	findMyMoodResult := db.DB.First(&mood, id)
	if findMyMoodResult.Error != nil || mood.UserID != userID {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mood not found or forbidden"})
	}

	db.DB.Delete(&models.Mood{}, id)
	return c.SendStatus(fiber.StatusNoContent)
}
