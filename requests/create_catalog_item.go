package requests

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/middleware"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/services"
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

	catalog.Name = r.Name
	catalog.Description = r.Description
	catalog.Price = r.Price
	catalog.CreatedBy = uuid.MustParse(user.Subject)

	return nil
}

func CreateCatalogItemRequestValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(CreateCatalogItemRequest)

	var count int64
	if err := catalogService.List().Where("name = ?", req.Name).Count(&count).Error; err != nil {
		sl.ReportError(req.Name, "Name", "Name", "unknown", "Unable to verify valid name")
	}

	if count > 0 {
		sl.ReportError(req.Name, "Name", "Name", "notFound", "Name already exists")
	}
}
