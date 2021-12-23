package ValidationExeption

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Exception struct {
	err validator.ValidationErrors
}

func New(err error) *Exception {
	e := err.(validator.ValidationErrors)
	return &Exception{e}
}

func (e Exception) FormatToFront() map[string]string {
	//var out []string
	out := make(map[string]string)

	for _, err := range e.err {
		out[err.Field()] = strings.ToLower(err.ActualTag())
	}

	return out
}
