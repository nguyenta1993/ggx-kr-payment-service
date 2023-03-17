package validation

import (
	"github.com/gogovan/ggx-kr-service-utils/validation"
)

type ValidationRulesEngine struct {
}

func NewValidationRulesEngine() *ValidationRulesEngine {
	return &ValidationRulesEngine{}
}

func (v *ValidationRulesEngine) GetValidations() []validation.CustomValidation {
	return []validation.CustomValidation{}
}
