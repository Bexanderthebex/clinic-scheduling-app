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
