package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

func Validation(request interface{}) error {
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				switch e.Tag() {
				case "required":
					message = fmt.Sprintf("Please enter your %s", e.Field())
				case "min":
					if e.Field() == "Password" {
						message = "Passwords must be at least 8 characters long"
					}
				case "email":
					message = "Please enter a valid email address"
				}
			}
			return errors.New(message)
		}
	}
	return nil
}
