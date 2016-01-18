package web
import (
	"testing"
	"github.com/byrnedo/svccommon/validate"
)

func TestIsJsonPathCorrect(t *testing.T) {
	type Item struct {
		SubItem string `validate:"len=10" json:"customsubitem"`
	}
	type TestStruct struct{
		SubStruct Item `json:"customitem"`
	}
	testStruct := TestStruct{Item{"dsfa"}}
	valErrs := validate.ValidateStruct(&testStruct)
	if valErrs == nil {
		t.Error("Should not be nil")
	}
	if valErrs["TestStruct.SubStruct.SubItem"].NameNamespace != "TestStruct.customitem.customsubitem" {
		t.Error("Not expected namespaced naming", valErrs["TestStruct.SubStruct.SubItem"] )
	}

	errResp := NewValidationErrorResonse(valErrs)
	t.Log(errResp)
	if len(errResp.Errors) == 0 {
		t.Error("No errors created")
	}

	if errResp.Errors[0].Source.Pointer != "customitem.customsubitem" {
		t.Error("Incorrect json pointer")
	}
}

