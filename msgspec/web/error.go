package web

import (
	"gopkg.in/bluesuncorp/validator.v8"
	"net/http"
	"fmt"
)

// JSON-API error response
//{
//	"jsonapi": { "version": "1.0" },
//	"errors": [
//		{
//			"code":   "123",
//			"source": { "pointer": "/data/attributes/first-name" },
//			"title":  "Value is too short",
//			"detail": "First name must contain at least three characters."
//		},
//		{
//			"code":   "226",
//			"source": { "pointer": "/data/attributes/password" },
//			"title": "Password and password confirmation do not match."
//		}
//	]
//}
type ErrorResponse struct {
	Errors []*ErrorMsg `json:"errors"` // We use an array of *Error to set common errors in variables. More on that later.
}

type Source struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

//Json-API error message.
// {
//	"code":   "123",
//	"source": { "pointer": "/data/attributes/first-name" },
//	"title":  "Value is too short",
//	"detail": "First name must contain at least three characters."
//}
type ErrorMsg struct {
	Code   int     `json:"status"`
	Source *Source `json:"source,omitempty"`
	Title  string  `json:"title,omitempty"`
	Detail string  `json:"detail"`
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func (e *ErrorResponse) AddError(code int, src *Source, title string, detail string) *ErrorResponse {
	e.Errors = append(e.Errors, &ErrorMsg{
		Code:   code,
		Source: src,
		Title:  title,
		Detail: detail,
	})
	return e
}

func (e *ErrorResponse) AddCodeError(code int) *ErrorResponse {
	e.Errors = append(e.Errors, &ErrorMsg{
		Code:   code,
		Source: nil,
		Title:  http.StatusText(code),
		Detail: http.StatusText(code),
	})
	return e
}

func NewValidationErrorResonse(valErrs validator.ValidationErrors) *ErrorResponse {
	errResponse := NewErrorResponse()
	for _, fieldErr := range valErrs {
		errResponse.AddError(400, &Source{Pointer: fieldErr.Name}, fieldErr.Tag, fieldErr.Tag)
	}
	return errResponse
}
