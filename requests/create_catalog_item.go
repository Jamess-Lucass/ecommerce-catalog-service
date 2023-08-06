package requests

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/middleware"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/services"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var catalogService *services.CatalogService

type CreateCatalogItemRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=128"`
	Description string  `json:"description" validate:"required,min=2,max=1024"`
	Price       float32 `json:"price" validate:"required"`
}

func (r *CreateCatalogItemRequest) Bind(c *fiber.Ctx, s *services.CatalogService, catalog *models.Catalog, v *validator.Validate) error {
	catalogService = s

	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Struct(r); err != nil {
		return err
	}

	user := c.Locals("claims").(*middleware.Claim)

	var count int64
	if err := catalogService.List(user).Where("name = ?", r.Name).Count(&count).Error; err != nil {
		return utils.Error{Errors: map[string][]string{"Name": {"Unable to verify valid name"}}}
	}

	if count > 0 {
		return utils.Error{Errors: map[string][]string{"Name": {"Name already exists"}}}
	}

	catalog.Name = r.Name
	catalog.Description = r.Description
	catalog.Price = r.Price
	catalog.CreatedBy = uuid.MustParse(user.Subject)

	return nil
}
