package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gryzlegrizz/bukuin_be/repository"
    "gorm.io/gorm"
)

func CheckUsername(c *fiber.Ctx) error {
    username := c.Query("username")
    if username == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username is required"})
    }

    db := c.Locals("db").(*gorm.DB)

    exists, err := repository.UsernameExists(db, username)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
    }

    return c.JSON(fiber.Map{"exists": exists})
}
