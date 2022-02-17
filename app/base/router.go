package base

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gopkg.in/guregu/null.v4"
	"reflect"
)

type Router struct {
	*gin.Engine
}

func LoadRouter() Router {
	var r = Router{
		Engine: gin.Default(),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterCustomTypeFunc(nullStringValidator, null.String{})
	}

	r.Static("storage/screenshot/", "./storage/screenshot/")

	return r
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
