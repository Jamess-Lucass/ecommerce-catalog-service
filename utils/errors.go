package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var caser = cases.Title(language.BritishEnglish)

type Error struct {
	Errors map[string][]string `json:"errors"`
}

func (m Error) Error() string {
	return "boom"
}

type ErrorResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func NewError(err error) ErrorResponse {
	e := ErrorResponse{Code: fiber.StatusBadRequest, Message: err.Error(), Errors: map[string][]string{}}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrorResponse{Code: fiber.StatusNotFound, Message: "Resource not found"}
	}

	if errors, ok := err.(Error); ok {
		e.Errors = errors.Errors
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, err := range validationErrors {
			property := toCamelCase(err.Field())
			e.Errors[property] = []string{getErrorMessage(err)}
		}

		e.Message = "One or more validation errors occurred."
	}

	return e
}

func toCamelCase(str string) string {
	words := strings.Split(str, " ")
	key := strings.ToLower(words[0])

	for _, word := range words[1:] {
		key += caser.String(word)
	}

	return key
}

func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "min":
		if err.Kind() == reflect.Slice {
			return fmt.Sprintf("This field must contain greater than %s element(s)", err.Param())
		}
		return fmt.Sprintf("This field must be greater than %s character(s)", err.Param())
	case "max":
		if err.Kind() == reflect.Slice {
			return fmt.Sprintf("This field must contain less than %s element(s)", err.Param())
		}
		return fmt.Sprintf("This field must be less than than %s character(s)", err.Param())
	default:
		return err.Param()
	}
}
