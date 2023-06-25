package handlers

import (
	"fmt"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/requests"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/utils"
	"github.com/goatquery/goatquery-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var maxTop = 1_000

// GET: /api/v1/catalog
func (s *Server) GetAllCatalogItems(c *fiber.Ctx) error {
	query := goatquery.Query{
		Top:     c.QueryInt("top", 0),
		Skip:    c.QueryInt("skip", 0),
		Count:   c.QueryBool("count", false),
		OrderBy: c.Query("orderby"),
		Select:  c.Query("select"),
		Search:  c.Query("search"),
		Filter:  c.Query("filter"),
	}

	var items []models.Catalog
	res, count, err := goatquery.Apply(s.catalogService.List(), query, &maxTop, func(db *gorm.DB, term string) *gorm.DB {
		t := fmt.Sprintf("%%%s%%", term)
		return db.Where("name like ? or description like ?", t, t)
	})
	if err != nil {
		return c.Status(400).JSON(goatquery.QueryErrorResponse{Status: 400, Message: err.Error()})
	}

	if err := res.Find(&items).Error; err != nil {
		return c.Status(400).JSON(goatquery.QueryErrorResponse{Status: 400, Message: err.Error()})
	}

	response := goatquery.BuildPagedResponse(items, query, count)

	return c.Status(200).JSON(response)
}

// GET: /api/v1/catalog/1
func (s *Server) GetCatalogItem(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	item, err := s.catalogService.Get(id)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(item)
}

// POST: /api/v1/catalog
func (s *Server) CreateCatalogItem(c *fiber.Ctx) error {
	var catalogItem models.Catalog
	req := &requests.CreateCatalogItemRequest{}
	if err := req.Bind(c, s.catalogService, &catalogItem, s.validator); err != nil {
		return c.Status(400).JSON(utils.NewError(err))
	}

	if err := s.catalogService.Create(&catalogItem); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(catalogItem)
}
