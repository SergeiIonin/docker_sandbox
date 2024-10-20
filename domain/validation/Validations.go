package validation

type Validations struct {
	SandboxNameValidation *SandboxNameValidation
}

func NewValidations() *Validations {
	return &Validations{
		SandboxNameValidation: NewSandboxNameValidation(),
	}
}

func (v *Validations) ValidateSandboxName(input string) (value string, err error) {
	return v.SandboxNameValidation.ValidateSandboxName(input)
}
