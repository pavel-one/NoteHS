package validations

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"
	"reflect"
)

func SetNullValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(nullStringValidator, null.String{})
	}
}

func nullStringValidator(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(null.String); ok {
		if !valuer.Valid {
			return nil
		}
		return valuer.String
	}

	return nil
}
