package routes

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type RouteValidationError struct {
	ValidationError validator.FieldError
}

func (fieldErr RouteValidationError) BuildResponseError() string {
	var stringBuilder strings.Builder

	stringBuilder.WriteString("Validation failed on field '" + fieldErr.ValidationError.Field() + "'")
	stringBuilder.WriteString(", condition: " + fieldErr.ValidationError.ActualTag())

	if fieldErr.ValidationError.Param() != "" {
		stringBuilder.WriteString(" { " + fieldErr.ValidationError.Param() + " }")
	}

	if fieldErr.ValidationError.Value() != nil && fieldErr.ValidationError.Value() != "" {
		stringBuilder.WriteString(fmt.Sprintf(", actual: %v", fieldErr.ValidationError.Value()))
	}

	return stringBuilder.String()
}
