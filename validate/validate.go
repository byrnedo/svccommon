package validate

import (
	"gopkg.in/bluesuncorp/validator.v8"
)

var (
	V *validator.Validate
)

func init() {
	V = validator.New(&validator.Config{
		FieldNameTag: "validate",
	})
}

func ValidateStruct(s interface{}) validator.ValidationErrors {
	if err := V.Struct(s); err != nil {
		return err.(validator.ValidationErrors)
	}
	return map[string] *validator.FieldError{}
}