package handler

import "github.com/gin-gonic/gin"

func (r *RunnginCtx) Init(ginEngine *gin.Engine) error {
	ginEngine.Any("/api/employee", r.handlerEmployee)
	return nil
}

func (r *RunnginCtx) Close() error {
	return nil
}
