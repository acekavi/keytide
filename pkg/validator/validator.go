package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Initialize sets up the validator
func Initialize() {
    validate = validator.New()
    
    // Register function to get json tag names
    validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return ""
        }
        return name
    })
}

// GetValidator returns the validator instance
func GetValidator() *validator.Validate {
    if validate == nil {
        Initialize()
    }
    return validate
}

// Validate validates a struct and returns validation errors
func Validate(s interface{}) []string {
    err := GetValidator().Struct(s)
    if err == nil {
        return nil
    }

    var errors []string
    for _, err := range err.(validator.ValidationErrors) {
        errors = append(errors, formatError(err))
    }
    
    return errors
}

// formatError formats a validation error into a human-readable message
func formatError(err validator.FieldError) string {
    field := err.Field()
    
    switch err.Tag() {
    case "required":
        return fmt.Sprintf("%s is required", field)
    case "email":
        return fmt.Sprintf("%s must be a valid email", field)
    case "min":
        return fmt.Sprintf("%s must be at least %s characters long", field, err.Param())
    case "max":
        return fmt.Sprintf("%s must be at most %s characters long", field, err.Param())
    default:
        return fmt.Sprintf("%s failed validation: %s", field, err.Tag())
    }
}