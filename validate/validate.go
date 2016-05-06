package validate

import (
	"gopkg.in/bluesuncorp/validator.v8"
	"reflect"
	"strconv"
	"strings"
	"regexp"
)

const (
	anyLanguageCharsRegexString = `^[\p{L}]+$`
	anyLanguageNameRegexString = `^\p{L}[\p{L}\s]*\p{L}$`
)

var (
	V *validator.Validate
	anyLanguageCharsRegex = regexp.MustCompile(anyLanguageCharsRegexString)
	anyLanguageNameRegex = regexp.MustCompile(anyLanguageNameRegexString)
)

var customValidators = map[string]validator.Func{
	"oneof": isIn,
	"anylangchars": anyLanguageChars,
	"anylangname": anyLanguageName,
}



func init() {
	V = validator.New(&validator.Config{
		TagName:      "validate",
		FieldNameTag: "json",
	})

	for key, function := range customValidators {
		if err := V.RegisterValidation(key, function); err != nil {
			panic("Failed to register validation func: " + err.Error())
		}

	}
}

// Checks if value is in list (written as "A#B#C")
// Usage in tag: `validate:"in=first#second#third"`
// NOTE this is not to be used, more of an example.
// Same can be achieved with '|' and the 'eq' validator
func isIn(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	possibleItems := strings.Split(param, "#")
	for _, item := range possibleItems {
		switch fieldKind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if v, err := strconv.ParseInt(item, 10, 64); err == nil {
				return field.Int() == v
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			if v, err := strconv.ParseUint(item, 10, 64); err == nil {
				return field.Uint() == v
			}
		case reflect.Float32, reflect.Float64:
			if v, err := strconv.ParseFloat(item, 64); err == nil {
				return field.Float() == v
			}
		}
		return field.String() == item
	}
	return false
}

func anyLanguageChars(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	return anyLanguageCharsRegex.MatchString(field.String())
}

func anyLanguageName(v *validator.Validate, topStruct reflect.Value, currentStruct reflect.Value, field reflect.Value, fieldtype reflect.Type, fieldKind reflect.Kind, param string) bool {
	return anyLanguageNameRegex.MatchString(field.String())
}

func ValidateStruct(any interface{}) (valErrors validator.ValidationErrors) {

	if errs := V.Struct(any); errs != nil {
		valErrors = errs.(validator.ValidationErrors)
	}
	return
}

//func ValidateStruct(s interface{}) validator.ValidationErrors {
//	if err := V.Struct(s); err != nil {
//		return err.(validator.ValidationErrors)
//	}
//	return map[string]*validator.FieldError{}
//}
