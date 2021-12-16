package req

import "github.com/gin-gonic/gin"

type Interface interface {
	GetCtx() *gin.Context
	SetCtx(ctx *gin.Context)
}

type Binder func(ctx *gin.Context, obj interface{}) error
