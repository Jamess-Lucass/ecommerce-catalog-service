package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func OptionalJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := getToken(c)

		if token == "" {
			return c.Next()
		}

		claims, err := parseToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": fiber.StatusUnauthorized, "message": "Unauthorized"})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}
