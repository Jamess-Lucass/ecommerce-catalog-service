package handlers

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/middleware"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// POST: /api/v1/catalog/123/like
func (s *Server) CreateLikeCatalogItem(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user := c.Locals("claims").(*middleware.Claim)

	catalogUserLike := &models.CatalogUserLike{
		CatalogID: id,
		UserID:    uuid.MustParse(user.Subject),
	}

	if err := s.catalogUserLikesService.Create(catalogUserLike); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// DELETE: /api/v1/catalog/123/like
func (s *Server) DeleteLikeCatalogItem(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user := c.Locals("claims").(*middleware.Claim)

	catalogUserLike := &models.CatalogUserLike{
		CatalogID: id,
		UserID:    uuid.MustParse(user.Subject),
	}

	if err := s.catalogUserLikesService.Delete(catalogUserLike); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
