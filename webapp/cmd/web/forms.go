package main

import (
	"net/url"
	"strings"
)

type errors map[string][]string

func (e errors) Get(field string) string {
	errorSlice := e[field]
	if len(errorSlice) == 0 {
		return ""
	}

	return errorSlice[0]

}

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

type Form struct {
	Data url.Values
	Errors errors
}

func NewForm(data url.Values) *Form {
	return &Form{
		Data: data,
		Errors: map[string][]string{},
	}
}

func (f *Form) Has(field string) bool {
	x := f.Data.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Check is a generic validation check. We can pass any expression
// that evaluates as a bool as the 1st param
func (f *Form) Check(ok bool, key, message string) {
	if !ok {
		f.Errors.Add(key, message)  
	}
}

// Valid returns true if there are no errors, otherwise false 
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}