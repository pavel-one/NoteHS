package Exceptions

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type HttpException struct {
	err validator.ValidationErrors
}

func New(err validator.ValidationErrors) *HttpException {
	return &HttpException{err}
}

func (e HttpException) FormatToFront() map[string]string {
	//var out []string
	out := make(map[string]string)

	for _, err := range e.err {
		out[err.Field()] = strings.ToLower(err.ActualTag())
	}

	return out
}
