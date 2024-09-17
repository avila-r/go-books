package validation

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var (
	v = validator.New()
)

// Validate function for any type that can be validated
func Validate[T any](request T) error {
	err := v.Struct(request)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Validation failed for field '%s', condition: %s", err.Field(), err.ActualTag())
		}

		return err
	}

	return nil
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
