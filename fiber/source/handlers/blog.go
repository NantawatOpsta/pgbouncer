package handlers

import (
	"myapp/entities"
	"myapp/settings"

	"github.com/gofiber/fiber/v2"
)

func GetBlogs(c *fiber.Ctx) error {
	var blogs []entities.Blog

	settings.Database.Find(&blogs)
	return c.Status(200).JSON(blogs)
}

