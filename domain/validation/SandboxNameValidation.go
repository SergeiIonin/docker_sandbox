package validation

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type SandboxNameValidation struct {
}

func NewSandboxNameValidation() *SandboxNameValidation {
	return &SandboxNameValidation{}
}

// validates the sandbox name. It should start with a letter, contain only letters, numbers, and underscores
func (snv *SandboxNameValidation) ValidateSandboxName(input string) (value string, err error) {
	validate := validator.New()
	regErr := validate.RegisterValidation("valid_sandbox_name", func(fl validator.FieldLevel) bool {
		name := fl.Field().String()
		match, err := regexp.MatchString("^[a-zA-Z][\\w]*[a-zA-Z0-9_]$", name)
		if err != nil {
			return false
		}
		return match
	})
	if regErr != nil {
		log.Printf("error registering validation: %v", regErr)
		return "", fmt.Errorf("error registering validation. %w", regErr)
	}
	if err := validate.Var(input, "valid_sandbox_name"); err != nil {
		log.Printf("error validating sandbox name: %v", err)
		return "", errors.New("invalid sandbox name")
	}

	return input, nil
}
