package handler

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func responseOK(ginCtx *gin.Context, data any) {
	h := gin.H{"code": 0, "message": "success"}
	if data != nil {
		h["data"] = data
	}
	ginCtx.JSON(http.StatusOK, h)
}

func parseRequest(s any, fieldName string, fn func(key string, value any, fieldValue string) bool) bool {
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		if !fn(structField.Name, v.Field(i).Interface(), structField.Tag.Get("json")) {
			return false
		}
	}
	return true
}
