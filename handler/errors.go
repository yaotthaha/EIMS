package handler

import "github.com/gin-gonic/gin"

type ErrMsg struct {
	code int
	msg  string
}

func (e ErrMsg) ToGinJSON() gin.H {
	return gin.H{"code": e.code, "message": e.msg}
}

var (
	errUnknownMethod      = ErrMsg{code: -1, msg: "Unknown Method"}
	errServerUnknownError = ErrMsg{code: -2, msg: "Server Unknown Error"}
)
