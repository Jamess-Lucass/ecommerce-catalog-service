package handlers

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/middleware"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/requests"
	"github.com/gofiber/contrib/fiberzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (s *Server) Start() error {
	s.validator.RegisterStructValidation(requests.CreateCatalogItemRequestValidation, requests.CreateCatalogItemRequest{})

	f := fiber.New()
	f.Use(cors.New(cors.Config{AllowOrigins: "*", AllowCredentials: true, MaxAge: 0}))

	f.Use(fiberzap.New(fiberzap.Config{
		Logger: s.logger,
	}))

	f.Get("/api/v1/catalog", s.GetAllCatalogItems)
	f.Post("/api/v1/catalog", middleware.JWT(), middleware.Role("Administrator", "Employee"), s.CreateCatalogItem)
	f.Get("/api/v1/catalog/:id", s.GetCatalogItem)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": "No resource found"})
	})

	return f.Listen(":8080")
}
