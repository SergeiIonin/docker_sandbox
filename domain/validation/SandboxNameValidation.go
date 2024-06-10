package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
	"regexp"
)

type SandboxNameValidation struct {
}

func NewSandboxNameValidation() *SandboxNameValidation {
	return &SandboxNameValidation{}
}

// validates the sandbox name. It should start with a letter, contain only letters, numbers, and underscores
func (snv *SandboxNameValidation) ValidateSandboxName(input string) (err error, value string) {
	validate := validator.New()
	validate.RegisterValidation("valid_sandbox_name", func(fl validator.FieldLevel) bool {
		name := fl.Field().String()
		match, err := regexp.MatchString("^[a-zA-Z][\\w]*[a-zA-Z0-9_]$", name)
		if err != nil {
			return false
		}
		return match
	})
	if err := validate.Var(input, "valid_sandbox_name"); err != nil {
		log.Printf("error validating sandbox name: %v", err)
		return errors.New("Invalid sandbox name"), ""
	}
	return nil, input
}
