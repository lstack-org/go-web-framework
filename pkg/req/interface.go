package req

import "github.com/gin-gonic/gin"

type Interface interface {
	GetCtx() *gin.Context
	SetCtx(ctx *gin.Context)
}

type Binder func(ctx *gin.Context, obj interface{}) error

//CustomizeBinder 自定义参数绑定
type CustomizeBinder interface {
	Bind(ctx *gin.Context) error
}
