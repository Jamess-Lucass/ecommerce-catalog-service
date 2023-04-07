package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func Role(role ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("claims").(*Claim)

		if !lo.Contains(role, user.Role) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"code": fiber.StatusForbidden, "message": "No access to this resource"})
		}

		return c.Next()
	}
}
