package handlers

import (
	"os"
	"strings"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/middleware"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.elastic.co/apm/module/apmfiber/v2"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ECSURL struct {
	model.URL
}

func (c *ECSURL) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if c.Path != "" {
		enc.AddString("path", c.Path)
	}

	return nil
}

func (s *Server) Start() error {
	f := fiber.New()
	f.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CORS_ALLOWED_ORIGINS"),
		AllowOriginsFunc: func(origin string) bool {
			return strings.EqualFold(os.Getenv("ENVIRONMENT"), "development")
		},
		AllowCredentials: true,
		MaxAge:           0,
	}))

	f.Use(middleware.SetTraceId(), apmfiber.Middleware())

	f.Use(fiberzap.New(fiberzap.Config{
		Logger: s.logger,
		Fields: []string{"latency", "status", "method"},
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			var fields []zap.Field

			tx := apm.TransactionFromContext(c.Context())
			if tx != nil {
				traceContext := tx.TraceContext()
				fields = append(fields, zap.String("trace.id", traceContext.Trace.String()))
				fields = append(fields, zap.String("transaction.id", traceContext.Span.String()))
				if span := apm.SpanFromContext(c.Context()); span != nil {
					fields = append(fields, zap.String("span.id", span.TraceContext().Span.String()))
				}
			}

			fields = append(fields, zap.Object("url", &ECSURL{model.URL{Path: c.OriginalURL()}}))

			return fields
		},
	}))

	f.Get("/api/healthz", s.Healthz)

	f.Get("/api/v1/catalog", middleware.OptionalJWT(), s.GetAllCatalogItems)
	f.Post("/api/v1/catalog", middleware.JWT(), middleware.Role("Administrator", "Employee"), s.CreateCatalogItem)
	f.Get("/api/v1/catalog/:id", middleware.OptionalJWT(), s.GetCatalogItem)

	f.Post("/api/v1/catalog/:id/like", middleware.JWT(), s.CreateLikeCatalogItem)
	f.Delete("/api/v1/catalog/:id/like", middleware.JWT(), s.DeleteLikeCatalogItem)

	f.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": "No resource found"})
	})

	return f.Listen(":8080")
}
