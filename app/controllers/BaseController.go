package controllers

import (
	"app/base"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	DB *base.DB
}

func (c Controller) Error(errors map[string]interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"errors":  errors,
	})
}

func (c Controller) Success(msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"msg":     msg,
	})
}
