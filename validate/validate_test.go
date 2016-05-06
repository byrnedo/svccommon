package validate

import (
	"testing"
)

func TestValidateStructAcceptsStructs(t *testing.T) {
	var x = struct {
		A string
	}{"Hello"}

	defer func() {
		if r := recover(); r != nil {
			t.Error("Should not have panicked")
		}
	}()

	ValidateStruct(x)
}

func TestValidateAnyLangChars(t *testing.T) {

	var x = struct {
		A string `validate:"required,anylangchars"`
	}{""}

	for _, val := range []string{"hello ", "$$", "/?", "12" } {
		x.A = val
		errs := ValidateStruct(&x)
		if errs == nil {
			t.Error("Should have failed")
		}
	}

	for _, val := range []string{"hello", "hellå", "你好" } {
		x.A = val

		errs := ValidateStruct(&x)
		if errs != nil {
			t.Error("Should not have failed")
		}
	}
}

func TestValidateAnyLangName(t *testing.T) {

	var x = struct {
		A string `validate:"required,anylangname"`
	}{""}

	for _, val := range []string{ "$$", "/?", "12", " donal", "donal " } {
		x.A = val
		errs := ValidateStruct(&x)
		if errs == nil {
			t.Error("Should have failed")
		}
	}

	for _, val := range []string{"donal", "dånal b", "你好 tada" } {
		x.A = val

		errs := ValidateStruct(&x)
		if errs != nil {
			t.Error("Should not have failed")
		}
	}
}

func TestValidateStructDoesNotAcceptOtherTypes(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Log("yay, panicked")
		}
	}()

	ValidateStruct("lkjkj")

	t.Error("Should have panicked.")
}

