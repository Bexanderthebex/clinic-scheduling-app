package routes

import (
	ginValidator "github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func NewValidator() *ginValidator.Validate {
	validator := ginValidator.New()
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return validator
}

func CheckForErrors(request interface{}, validator *ginValidator.Validate) *RouteValidationError {
	validationErrors := validator.Struct(request)

	if validationErrors != nil {
		fieldErrors := validationErrors.(ginValidator.ValidationErrors)
		if len(validationErrors.(ginValidator.ValidationErrors)) > 0 {
			return &RouteValidationError{
				ValidationError: fieldErrors[0],
			}
		}
	}

	return nil
}
