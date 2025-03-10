package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs-karasal/bitlab_go/lesson2/db"
	"github.com/rs-karasal/bitlab_go/lesson2/models"
	"github.com/rs-karasal/bitlab_go/lesson2/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hash)

	result := db.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var input models.User
	var user models.User

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	result := db.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid password"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}
