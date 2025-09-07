package validation

import "github.com/go-playground/validator/v10"

var DefaultValidator = validator.New(validator.WithRequiredStructEnabled())
